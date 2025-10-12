package main

import (
	"github.com/urfave/cli/v3"
)

var cmd = &cli.Command{
	Name:           "potato-squad.de",
	Usage:          "haters gonna hate, potatoes gonna potate",
	DefaultCommand: "run",
}
