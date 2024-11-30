package service

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/oklog/run"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/temporalio/reference-app-orders-go/internal/config"
	"github.com/temporalio/reference-app-orders-go/internal/interceptors"
	"github.com/temporalio/reference-app-orders-go/internal/temporalutil"
	"github.com/uber-go/tally/v4"
	"github.com/uber-go/tally/v4/prometheus"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	sdktally "go.temporal.io/sdk/contrib/tally"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"gorm.io/gorm"
)

type (
	RunParams struct {
		Config   config.AppConfig
		DB       *gorm.DB
		Logger   *slog.Logger
		Temporal client.Client
	}

	Runnable func(context.Context, *RunParams) error
)

// global Flags
var Flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "debug",
		Usage:   "enable debug logging",
		EnvVars: []string{"DEBUG"},
	},
	// The encryption key ID is a string that can be used to look up an encryption
	// key (e.g., from a key management system). If this option is specified, then
	// inputs to Workflows and Activities, as well as the outputs returned by the
	// Workflows and Activities, will be encrypted with that key before being sent
	// by the Client in this application. This Client will likewise decrypt them
	// upon receipt. The Temporal CLI and Web UI will be unable to view the original
	// (unencrypted) data unless you run a Codec server and configure them to use it.
	&cli.StringFlag{
		Name:    "encryption-key",
		Aliases: []string{"k"},
		Usage:   "id of key used to encrypt temporal payloads",
		EnvVars: []string{"OMS_ENCRYPTION_KEY"},
	},
}

// CreateClientOptionsFromEnv creates a client.Options instance, configures
// it based on environment variables, and returns that instance. It
// supports the following environment variables:
//
//	TEMPORAL_ADDRESS: Host and port (formatted as host:port) of the Temporal Frontend Service
//	TEMPORAL_NAMESPACE: Namespace to be used by the Client
//	TEMPORAL_TLS_CERT: Path to the x509 certificate
//	TEMPORAL_TLS_KEY: Path to the private certificate key
//
// If these environment variables are not set, the client.Options
// instance returned will be based on the SDK's default configuration.
func CreateClientOptionsFromEnv() (client.Options, error) {
	hostPort := os.Getenv("TEMPORAL_ADDRESS")
	namespaceName := os.Getenv("TEMPORAL_NAMESPACE")

	// Must explicitly set the Namepace for non-cloud use.
	if strings.Contains(hostPort, ".tmprl.cloud:") && namespaceName == "" {
		return client.Options{}, errors.New("namespace name unspecified; required for Temporal Cloud")
	}

	if namespaceName == "" {
		namespaceName = "default"
	}

	scope, err := newPrometheusScope(prometheus.Configuration{
		ListenAddress: "0.0.0.0:9090",
		TimerType:     "histogram",
	})
	if err != nil {
		return client.Options{}, fmt.Errorf("error initializing metrics scope: %w", err)
	}

	clientOpts := client.Options{
		HostPort:       hostPort,
		Interceptors:   []interceptor.ClientInterceptor{interceptors.NewValidation()},
		Namespace:      namespaceName,
		Logger:         log.NewStructuredLogger(slog.Default()),
		MetricsHandler: sdktally.NewMetricsHandler(scope),
	}

	if certPath := os.Getenv("TEMPORAL_TLS_CERT"); certPath != "" {
		cert, err := tls.LoadX509KeyPair(certPath, os.Getenv("TEMPORAL_TLS_KEY"))
		if err != nil {
			return clientOpts, fmt.Errorf("failed loading key pair: %w", err)
		}

		clientOpts.ConnectionOptions.TLS = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}

	return clientOpts, nil
}

// RunAPIServer runs a API HTTP server for the given service.
func RunAPIServer(ctx context.Context, hostPort string, router http.Handler, logger *slog.Logger) error {
	srv := &http.Server{
		Addr:    hostPort,
		Handler: loggingMiddleware(logger, router),
	}

	logger.Info("Listening", "endpoint", "http://"+hostPort)

	errCh := make(chan error, 1)
	go func() { errCh <- srv.ListenAndServe() }()

	select {
	case <-ctx.Done():
		srv.Close()
	case err := <-errCh:
		return err
	}

	return nil
}

// RunConnectServer runs a API HTTP server for the given service.
func RunConnectServer(ctx context.Context, hostPort string, router http.Handler, logger *slog.Logger) error {
	srv := &http.Server{
		Addr:    hostPort,
		Handler: h2c.NewHandler(router, &http2.Server{}),
	}

	logger.Info("Listening", "endpoint", "http://"+hostPort)

	errCh := make(chan error, 1)
	go func() { errCh <- srv.ListenAndServe() }()

	select {
	case <-ctx.Done():
		srv.Close()
	case err := <-errCh:
		return err
	}

	return nil
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// NewClient initializes a temporal client using configuration derived from the
// command context
func NewClient(cmd *cli.Context) (client.Client, error) {
	clientOptions, err := CreateClientOptionsFromEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to create client options: %w", err)
	}
	if encryptionKeyID := cmd.String("encryption-key"); encryptionKeyID != "" {
		if cmd.Bool("debug") {
			slog.Default().Debug("Enabling encrypting Data Converter", slog.String("key_id", encryptionKeyID))
		}
		ddc := converter.NewCompositeDataConverter(
			converter.NewNilPayloadConverter(),
			converter.NewByteSlicePayloadConverter(),
			converter.NewProtoPayloadConverter(),
			converter.NewProtoJSONPayloadConverter(),
			converter.NewJSONPayloadConverter(),
		)
		clientOptions.DataConverter = temporalutil.NewEncryptionDataConverter(ddc, encryptionKeyID)
	}

	c, err := client.Dial(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("client error: %w", err)
	}
	return c, nil
}

// NewServiceCommands returns a modified version of the provided command with an
// explicit name, usage, and service subcommand that can be used to run 1 or
// more runnables
func NewServiceCommands(cmd *cli.Command, name, usage string, runnables ...Runnable) *cli.Command {
	cmd.Name = name
	cmd.Usage = usage
	injectGlobalFlags(cmd)
	cmd.Subcommands = append(cmd.Subcommands, &cli.Command{
		Name:  "service",
		Usage: fmt.Sprintf("run %s service", name),
		Flags: Flags,
		Action: func(cmd *cli.Context) error {
			logger := slog.Default().With(slog.String("service", name))
			cfg, err := config.AppConfigFromEnv()
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			c, err := NewClient(cmd)
			if err != nil {
				return fmt.Errorf("failed to initialize temporal client: %w", err)
			}
			defer c.Close()

			dbPath := path.Join(cfg.DataDir, "api-store.db")
			db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
			if err != nil {
				return fmt.Errorf("failed to open database: %w", err)
			}

			var g run.Group
			for _, run := range runnables {
				sctx, scancel := context.WithCancel(cmd.Context)
				g.Add(
					func() error {
						return run(sctx, &RunParams{Config: cfg, DB: db, Logger: logger, Temporal: c})
					},
					func(err error) {
						scancel()
					},
				)
			}
			return g.Run()
		},
	})
	return cmd
}

func injectGlobalFlags(cmd *cli.Command) {
	cmd.Flags = append(cmd.Flags, Flags...)
	for _, scmd := range cmd.Subcommands {
		injectGlobalFlags(scmd)
	}
}

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		iw := instrumentedResponseWriter{w, http.StatusOK}

		next.ServeHTTP(&iw, r)

		level := slog.LevelDebug
		if iw.status >= 500 {
			level = slog.LevelError
		}

		logger.Log(
			context.Background(), level,
			fmt.Sprintf("%d %s %s", iw.status, r.Method, r.URL.Path),
			"method", r.Method, "status", iw.status, "path", r.URL.Path,
		)
	})
}

func newPrometheusScope(c prometheus.Configuration) (tally.Scope, error) {
	reporter, err := c.NewReporter(
		prometheus.ConfigurationOptions{
			Registry: prom.NewRegistry(),
			OnError: func(err error) {
				slog.Default().Error("error in prometheus reporter", slog.AnyValue(err))
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error creating prometheus reporter", err)
	}
	scopeOpts := tally.ScopeOptions{
		CachedReporter:  reporter,
		Separator:       prometheus.DefaultSeparator,
		SanitizeOptions: &sdktally.PrometheusSanitizeOptions,
	}
	scope, _ := tally.NewRootScope(scopeOpts, time.Second)
	scope = sdktally.NewPrometheusNamingScope(scope)

	return scope, nil
}

type instrumentedResponseWriter struct {
	http.ResponseWriter
	status int
}

func (r *instrumentedResponseWriter) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}
