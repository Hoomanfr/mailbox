package application

import (
	"github.com/thumperq/messaging/mailbox/internal/infrastructure/db"
)

type AppFactory struct {
	MailboxApp MailboxApp
	UserApp    UserApp
}

func NewApplicationFactory(DbFactory db.DbFactory) AppFactory {
	return AppFactory{
		MailboxApp: NewMailboxApp(DbFactory),
		UserApp:    NewUserApp(DbFactory),
	}
}
