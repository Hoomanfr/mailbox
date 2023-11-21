package application

import (
	"github.com/thumperq/golib/messaging"
	"github.com/thumperq/wms/mailbox/internal/infrastructure/db"
)

type AppFactory struct {
	MailboxApp MailboxApp
	UserApp    UserApp
}

func NewApplicationFactory(DbFactory db.DbFactory, broker *messaging.Broker) AppFactory {
	return AppFactory{
		MailboxApp: NewMailboxApp(DbFactory, broker),
		UserApp:    NewUserApp(DbFactory, broker),
	}
}
