package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	root "github.com/potato-squad-org/site"
	"github.com/urfave/cli/v2"
)

var (
	app = &cli.App{
		Name:     root.Meta().Name(),
		HelpName: "site",
		Usage:    "potato-squad.org Site",
	}
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.RunContext(ctx, os.Args); err != nil {
		slog.Error("Error running app", slog.String("msg", err.Error()))
		os.Exit(1)
	}
}
