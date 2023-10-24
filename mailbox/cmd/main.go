package main

import (
	"os"

	httpserver "github.com/Hoomanfr/golib/servers/http"
)

func main() {
	exitCode := <-httpserver.ListenAndServe(bootstrap)
	os.Exit(exitCode)
}
