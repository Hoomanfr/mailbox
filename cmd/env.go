package main

import (
	"context"

	"github.com/thumperq/golib/config"
	"github.com/thumperq/golib/logging"
	"github.com/thumperq/golib/messaging"
	httpserver "github.com/thumperq/golib/servers/http"
	"github.com/thumperq/wms/mailbox/api"
	"github.com/thumperq/wms/mailbox/internal/application"
	"github.com/thumperq/wms/mailbox/internal/consumers"
	"github.com/thumperq/wms/mailbox/internal/domain"
	"github.com/thumperq/wms/mailbox/internal/infrastructure/db"
)

type Env struct {
	Cfg    config.CfgManager
	Broker *messaging.Broker
}

func NewEnv() *Env {
	logging.SetupLogging()
	return &Env{}
}

func (env *Env) Bootstrap(apiSrv *httpserver.ApiServer) {
	env.Cfg = config.NewConfigManager()
	dbFactory := db.NewDbFactory(env.Cfg)
	env.Broker = env.runSubscribers(context.Background(), dbFactory)
	appFactory := application.NewApplicationFactory(dbFactory, env.Broker)

	api.SetupUserApi(appFactory, apiSrv.Engine)
	api.SetupMailboxApi(appFactory, apiSrv.Engine)
}

func (env *Env) runSubscribers(ctx context.Context, dbFactory db.DbFactory) *messaging.Broker {
	panicIfError := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	broker, err := messaging.NewBroker(env.Cfg, "wms", "mailbox")
	panicIfError(err)
	err = broker.Connect()
	panicIfError(err)
	mailboxConsumers := consumers.NewMailboxConsumer(dbFactory)
	err = messaging.NewSubscriber[domain.MailboxCreated](broker).
		Subscribe(ctx, "wms", "mailbox", domain.MailboxTopic, mailboxConsumers.ConsumeMailboxCreatedEvent)
	panicIfError(err)
	return broker
}
