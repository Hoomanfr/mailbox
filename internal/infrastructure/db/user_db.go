package db

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thumperq/golib/database"
	"github.com/thumperq/wms/mailbox/internal/domain"
)

type UserDB interface {
	Create(ctx context.Context, user domain.User) error
	FindById(ctx context.Context, id string) (*domain.User, error)
}
type UserDb struct {
	pgDb database.PgDB
}

func NewUserDb(pgDb database.PgDB) UserDB {
	return UserDb{
		pgDb: pgDb,
	}
}

func (db UserDb) Create(ctx context.Context, user domain.User) error {
	err := db.pgDb.WithTransaction(ctx, func(tx pgx.Tx) error {
		sql, args, err := sb.Insert("users").
			Columns("id", "user_name", "password", "created_at").
			Values(user.ID, user.UserName, user.Password, time.Now().UTC()).
			ToSql()
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (db UserDb) FindById(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	err := db.pgDb.WithConnection(ctx, func(c *pgxpool.Conn) error {
		sql, args, err := sb.Select("id", "user_name", "password", "created_at").
			From("users").
			Where(squirrel.Eq{"id": id}).
			ToSql()
		if err != nil {
			return err
		}
		row := c.QueryRow(ctx, sql, args...)
		err = row.Scan(&user.ID, &user.UserName, &user.Password, &user.CreatedAt)
		if err != nil {
			return err
		}
		return nil
	})
	return &user, err
}
