package main

import (
	"github.com/lemmack/multihasher/internal/server"
)

// Consider adding a version variable which is assigned at runtime by a makefile
// var version string

const portString string = ":8000"                  // Port the server will run on
const localClient string = "http://127.0.0.1:5500" // Address of a local client if one is used, set to empty string to disable

func main() {
	server.Start(portString, localClient)
}
