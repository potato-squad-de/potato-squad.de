package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/potato-squad-org/site/web"
	"github.com/urfave/cli/v2"
)

func init() {
	app.Commands = append(app.Commands, runCmd)
}

var httpServer *http.Server

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Run the site",
	Aliases: []string{
		"start",
		"serve",
	},
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:       "port",
			Usage:      "Port to run the site on",
			Aliases:    []string{"p"},
			EnvVars:    []string{"PORT"},
			Required:   true,
			HasBeenSet: true,
			Value:      defaultHTTPPort,
			Action: func(ctx *cli.Context, i int) error {
				httpServer = web.HTTPServer(i)
				return nil
			},
		},
	},
	Action: func(c *cli.Context) error {
		if httpServer == nil {
			return errors.New("http server not initialized")
		}
		go func() {
			if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				slog.Error("Error running http server", slog.String("msg", err.Error()))
			}
		}()

		slog.Info("server running", slog.String("addr", httpServer.Addr))

		slog.Info("press ctrl+c to exit")

		cancelSig := make(chan os.Signal, 1)
		signal.Notify(cancelSig, os.Interrupt)
		slog.Info("cancellation signal received", slog.Any("signal", <-cancelSig))

		ctx, cancel := context.WithTimeout(c.Context, 15*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			slog.Error("Error shutting down http server", slog.String("msg", err.Error()))
		}
		slog.Info("server shut down")

		return nil
	},
}
