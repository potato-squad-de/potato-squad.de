//go:build dev
// +build dev

package main

import (
	_ "github.com/joho/godotenv/autoload"
)

var (
	defaultHTTPPort = 8080
)
