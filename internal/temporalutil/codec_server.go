package temporalutil

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/cludden/protoc-gen-go-temporal/pkg/codec"
	"github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
	billingv1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
	orderv1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	"go.temporal.io/sdk/converter"
)

func newCORSHTTPHandler(origin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type,X-Namespace")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

// RunCodecServer launches the Codec Server on the specified port, enabling
// CORS for the Temporal Web UI at the specified URL
func RunCodecServer(port int, url string) error {
	// The EncryptionKeyID attribute is omitted when creating the Codec
	// instance below because the Codec Server only decrypts. It locates
	// the encryption key ID from the payload's metadata.
	handler := converter.NewPayloadCodecHTTPHandler(
		&Codec{},
		codec.NewProtoJSONCodec(scheme.New(
			billingv1.WithWorkerSchemeTypes(),
			orderv1.WithWorkerSchemeTypes(),
			shipmentv1.WithWorkerSchemeTypes(),
		)),
	)

	if url != "" {
		handler = newCORSHTTPHandler(url, handler)
	}

	srv := &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(port),
		Handler: handler,
	}

	errCh := make(chan error, 1)
	go func() { errCh <- srv.ListenAndServe() }()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	select {
	case <-sigCh:
		_ = srv.Close()
	case err := <-errCh:
		return err
	}

	return nil
}
