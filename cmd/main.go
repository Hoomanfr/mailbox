package main

import (
	"context"
	"os"

	"github.com/thumperq/golib/logging"
	httpserver "github.com/thumperq/golib/servers/http"
)

func main() {
	env := NewEnv()
	exitCode := <-httpserver.ListenAndServe(env.Bootstrap)
	err := env.Broker.Disconnect()
	if err != nil {
		logging.TraceLogger(context.Background()).
			Err(err).
			Msg("error disconnecting from broker")
	}
	os.Exit(exitCode)
}
