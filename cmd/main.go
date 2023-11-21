package main

import (
	"os"

	httpserver "github.com/thumperq/golib/servers/http"
	"github.com/thumperq/wms/mailbox/api"
)

func main() {
	env := &api.Env{}
	exitCode := <-httpserver.ListenAndServe(env.Bootstrap)
	os.Exit(exitCode)
}
