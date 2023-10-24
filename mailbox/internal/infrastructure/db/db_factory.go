package db

import (
	"github.com/Hoomanfr/golib/config"
	"github.com/Hoomanfr/golib/database"
	"github.com/Masterminds/squirrel"
)

var sb = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type DbFactory struct {
	UserDb    UserDB
	MailboxDb MailboxDB
}

func NewDbFactory(cfg *config.ConfigManager) DbFactory {
	pgDb := database.NewPostgresConnection(cfg)
	return DbFactory{
		UserDb:    NewUserDb(pgDb),
		MailboxDb: NewMailboxDb(pgDb),
	}
}
