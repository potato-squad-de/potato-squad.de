package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/potato-squad-de/potato-squad.de/web"
	"github.com/urfave/cli/v3"
)

var httpServer *http.Server

func init() {
	cmd.Commands = append(cmd.Commands, &cli.Command{
		Name:    "serve",
		Aliases: []string{"run"},
		Flags: []cli.Flag{
			&cli.Uint16Flag{
				Name:     "http-port",
				Sources:  cli.EnvVars("HTTP_PORT"),
				Required: true,
				Value:    uint16(8080),
			},
		},
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			httpServer = &http.Server{
				Addr:    fmt.Sprintf(":%d", c.Uint16("http-port")),
				Handler: web.Handler(),
			}

			return ctx, nil
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			go func() {
				if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
					slog.Error("failed to run web server", slog.String("err", err.Error()))
				}
			}()

			slog.Info("web server running", slog.String("addr", httpServer.Addr))
			slog.Info("press ctrl+c to stop")

			cancelSig := make(chan os.Signal, 1)
			signal.Notify(cancelSig, os.Interrupt)
			slog.Info("cancellation signal received", slog.Any("sig", <-cancelSig))

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			if err := httpServer.Shutdown(ctx); err != nil {
				return fmt.Errorf("failed to shut down the http server: %s", err.Error())
			}
			slog.Info("server shut down")

			return nil
		},
	})
}
