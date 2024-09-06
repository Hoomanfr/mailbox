package main

import (
	"context"

	"github.com/thumperq/golib/application"
	"github.com/thumperq/golib/database"
	"github.com/thumperq/golib/environment"
	"github.com/thumperq/golib/logging"
	"github.com/thumperq/wms/mailbox/api"
	"github.com/thumperq/wms/mailbox/internal/app"
	"github.com/thumperq/wms/mailbox/internal/consumers"
	"github.com/thumperq/wms/mailbox/internal/domain"
	"github.com/thumperq/wms/mailbox/internal/infrastructure/db"
)

func main() {
	err := environment.NewEnv().
		WithBroker().
		WithDbFactory().
		WithAppFactory().
		WithWorker().
		Bootstrap(bootstrap)

	if err != nil {
		logging.TraceLogger(context.Background()).
			Err(err).
			Msg("error bootstrapping environment")
	}
}

func bootstrap(env *environment.Env) error {
	setupDbs(env)
	setupApps(env)
	setupApis(env)
	err := runSubscribers(env)
	if err != nil {
		return err
	}
	return nil
}

func setupDbs(env *environment.Env) {
	env.DbFactory.Register(func(dbf database.DbFactory) any {
		return db.NewUserDb(dbf.PgDb())
	}).Register(func(dbf database.DbFactory) any {
		return db.NewMailboxDb(dbf.PgDb())
	})
}

func setupApps(env *environment.Env) {
	env.AppFactory.Register(func(appFactory application.AppFactory) any {
		return app.NewMailboxApp(env.Broker, environment.GetRepo[db.MailboxDB]())
	}).Register(func(appFactory application.AppFactory) any {
		return app.NewUserApp(env.Broker, environment.GetRepo[db.UserDB]())
	})
}

func setupApis(env *environment.Env) {
	api.SetupUserApi(environment.GetApp[app.UserApp](), env.ApiServer.Engine)
	api.SetupMailboxApi(environment.GetApp[app.MailboxApp](), env.ApiServer.Engine)
}

func runSubscribers(env *environment.Env) error {
	ctx := context.Background()
	mailboxConsumers := consumers.NewMailboxConsumer(environment.GetRepo[db.MailboxDB]())
	err := env.Worker.Run(mailboxConsumers)(ctx, "wms", "mailbox", domain.MailboxTopic)
	if err != nil {
		return err
	}
	return nil
}
