package main

import (
	httpserver "github.com/thumperq/golib/servers/http"
	"github.com/thumperq/messaging/mailbox/api"
	"github.com/thumperq/messaging/mailbox/internal/application"
	"github.com/thumperq/messaging/mailbox/internal/infrastructure/db"
)

func bootstrap(apiSrv *httpserver.ApiServer) {
	dbFactory := db.NewDbFactory(apiSrv.ConfigManager)
	appFactory := application.NewApplicationFactory(dbFactory)

	api.SetupUserApi(appFactory, apiSrv.Engine)
	api.SetupMailboxApi(appFactory, apiSrv.Engine)
}
