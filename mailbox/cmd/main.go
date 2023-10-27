package main

import (
	"os"

	httpserver "github.com/thumperq/golib/servers/http"
)

func main() {
	exitCode := <-httpserver.ListenAndServe(bootstrap)
	os.Exit(exitCode)
}
