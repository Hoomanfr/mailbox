package db

import (
	"github.com/Masterminds/squirrel"
	"github.com/thumperq/golib/config"
	"github.com/thumperq/golib/database"
)

var sb = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type DbFactory struct {
	UserDb    UserDB
	MailboxDb MailboxDB
}

func NewDbFactory(cfg config.CfgManager) DbFactory {
	pgDb := database.NewPostgresConnection(cfg)
	return DbFactory{
		UserDb:    NewUserDb(pgDb),
		MailboxDb: NewMailboxDb(pgDb),
	}
}
