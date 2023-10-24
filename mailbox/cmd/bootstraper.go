package main

import (
	httpserver "github.com/Hoomanfr/golib/servers/http"
	"github.com/Hoomanfr/messaging/mailbox/api"
	"github.com/Hoomanfr/messaging/mailbox/internal/application"
	"github.com/Hoomanfr/messaging/mailbox/internal/infrastructure/db"
)

func bootstrap(apiSrv *httpserver.ApiServer) {
	dbFactory := db.NewDbFactory(apiSrv.ConfigManager)
	appFactory := application.NewApplicationFactory(dbFactory)

	api.SetupUserApi(appFactory, apiSrv.Engine)
	api.SetupMailboxApi(appFactory, apiSrv.Engine)
}
