//go:build !libznn

package main

import (
	"github.com/zenon-network/go-zenon/app"
)

// hqzd is the hyperqube command-line client
func main() {
	app.Run()
}
