package main

import (
	"github.com/obscuronet/obscuro-playground/go/host/hostrunner"
)

// Runs an Obscuro host as a standalone process.
func main() {
	config := hostrunner.ParseConfig()
	// We set the logs outside of `RunHost` so we can override the logging in tests.
	hostrunner.SetLogs(config.LogPath)
	hostrunner.RunHost(config)
}