package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error("application error", slog.String("err", err.Error()))
		os.Exit(1)
	}
}
