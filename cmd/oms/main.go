package main

import (
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/temporalio/reference-app-orders-go/app/billing"
	"github.com/temporalio/reference-app-orders-go/app/fraud"
	"github.com/temporalio/reference-app-orders-go/app/order"
	"github.com/temporalio/reference-app-orders-go/app/shipment"
	billingv1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
	orderv1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	"github.com/temporalio/reference-app-orders-go/internal/service"
	"github.com/temporalio/reference-app-orders-go/internal/temporalutil"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "oms",
		Usage: "Order Management System",
		Flags: service.Flags,
		Before: func(cmd *cli.Context) error {
			level := slog.LevelInfo
			if cmd.Bool("debug") {
				level = slog.LevelDebug
			}
			logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: level,
			}))
			slog.SetDefault(logger)
			return nil
		},
		Commands: []*cli.Command{
			service.NewServiceCommands(
				service.Must(billingv1.NewWorkerCliCommand(billingv1.NewWorkerCliOptions().WithClient(service.NewClient))),
				"billing", "billing service commands", billing.RunServer, billing.RunWorker,
			),

			{
				Name:  "codec-server",
				Usage: "Codec Server decrypts payloads for display by Temporal CLI and Web UI",
				Flags: append([]cli.Flag{
					&cli.Int64Flag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "Port number on which the Codec Server will listen for requests",
						Value:   8089,
					},
					&cli.StringFlag{
						Name:    "url",
						Aliases: []string{"u"},
						Usage:   "Temporal Web UI base URL (allow CORS for that origin)",
					},
				}, service.Flags...),
				Action: func(cmd *cli.Context) error {
					codecCorsURL := cmd.String("url")
					if codecCorsURL != "" {
						log.Printf("Codec Server will allow requests from Temporal Web UI at: %s\n", codecCorsURL)

						if strings.HasSuffix(codecCorsURL, "/") {
							log.Println("Warning: Temporal Web UI base URL ends with '/'")
						}
					}
					codecPort := cmd.Int("port")

					log.Printf("Starting Codec Server on port %d\n", codecPort)
					err := temporalutil.RunCodecServer(codecPort, codecCorsURL)

					return err
				},
			},

			service.NewServiceCommands(
				service.Must(orderv1.NewWorkerCliCommand(orderv1.NewWorkerCliOptions().WithClient(service.NewClient))),
				"order", "order service commands", order.RunServer, order.RunWorker,
			),

			service.NewServiceCommands(
				service.Must(shipmentv1.NewWorkerCliCommand(shipmentv1.NewWorkerCliOptions().WithClient(service.NewClient))),
				"shipment", "shipment service commands", shipment.RunServer, shipment.RunWorker,
			),

			service.NewServiceCommands(&cli.Command{
				Name:  "fraud",
				Usage: "run fraud service",
			}, "fraud", "fake external service", fraud.RunServer),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
