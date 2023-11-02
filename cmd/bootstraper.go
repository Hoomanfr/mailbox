package main

import (
	httpserver "github.com/thumperq/golib/servers/http"
	"github.com/thumperq/wms/mailbox/api"
	"github.com/thumperq/wms/mailbox/internal/application"
	"github.com/thumperq/wms/mailbox/internal/infrastructure/db"
)

func bootstrap(apiSrv *httpserver.ApiServer) {
	dbFactory := db.NewDbFactory(apiSrv.ConfigManager)
	appFactory := application.NewApplicationFactory(dbFactory)

	api.SetupUserApi(appFactory, apiSrv.Engine)
	api.SetupMailboxApi(appFactory, apiSrv.Engine)
}
