package main

import (
	"github.com/lemmack/multihasher/internal/server"
)

// Consider adding a version variable which is assigned at runtime by a makefile
// var version string

// Starts the server
func main() {
	server.Start()
}
