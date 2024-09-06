package main

import (
	"context"
	"os"

	"github.com/thumperq/golib/application"
	"github.com/thumperq/golib/database"
	"github.com/thumperq/golib/environment"
	"github.com/thumperq/golib/logging"
	"github.com/thumperq/golib/messaging"
	httpserver "github.com/thumperq/golib/servers/http"
	"github.com/thumperq/wms/mailbox/api"
	"github.com/thumperq/wms/mailbox/internal/app"
	"github.com/thumperq/wms/mailbox/internal/consumers"
	"github.com/thumperq/wms/mailbox/internal/domain"
	"github.com/thumperq/wms/mailbox/internal/infrastructure/db"
)

func main() {
	env, err := environment.NewEnv()
	if err != nil {
		logging.TraceLogger(context.Background()).
			Err(err).
			Msg("error creating new environment")
		os.Exit(1)
	}
	err = env.Bootstrap(bootstrap)
	if err != nil {
		logging.TraceLogger(context.Background()).
			Err(err).
			Msg("error bootstrapping environment")
	}
}
func bootstrap(env *environment.Env, apiSrv *httpserver.ApiServer) error {
	env.DbFactory.Register(func(dbf database.DbFactory) any {
		return db.NewUserDb(dbf.PgDb())
	}).Register(func(dbf database.DbFactory) any {
		return db.NewMailboxDb(dbf.PgDb())
	})

	env.AppFactory.Register(func(appFactory application.AppFactory) any {
		return app.NewMailboxApp(env.Broker, environment.GetRepo[db.MailboxDB]())
	}).Register(func(appFactory application.AppFactory) any {
		return app.NewUserApp(env.Broker, environment.GetRepo[db.UserDB]())
	})

	err := runSubscribers(context.Background(), env)
	if err != nil {
		return err
	}
	api.SetupUserApi(environment.GetApp[app.UserApp](), apiSrv.Engine)
	api.SetupMailboxApi(environment.GetApp[app.MailboxApp](), apiSrv.Engine)
	return nil
}

func runSubscribers(ctx context.Context, env *environment.Env) error {
	mailboxConsumers := consumers.NewMailboxConsumer(environment.GetRepo[db.MailboxDB]())
	err := messaging.NewSubscriber(env.Broker).
		Subscribe(ctx, "wms", "mailbox", domain.MailboxTopic, mailboxConsumers.ConsumeMailboxCreatedEvent)
	if err != nil {
		return err
	}
	return nil
}
