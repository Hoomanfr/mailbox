package application

import (
	"context"

	"github.com/thumperq/golib/messaging"
	"github.com/thumperq/wms/mailbox/internal/domain"
	"github.com/thumperq/wms/mailbox/internal/infrastructure/db"
)

type MailboxRequest struct {
	UserId string `json:"user_id" binding:"required"`
	Email  string `json:"email" binding:"required"`
}

type MailboxResponse struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Email  string `json:"email"`
}

type MailboxApp struct {
	dbFactory db.DbFactory
	broker    *messaging.Broker
}

func NewMailboxApp(DbFactory db.DbFactory, broker *messaging.Broker) MailboxApp {
	app := MailboxApp{
		dbFactory: DbFactory,
		broker:    broker,
	}
	return app
}

func (app MailboxApp) CreateMailbox(ctx context.Context, request MailboxRequest) (string, error) {
	mailbox, err := domain.NewMailbox(request.UserId, request.Email)
	if err != nil {
		return "", err
	}
	if err := app.dbFactory.MailboxDb.Create(ctx, *mailbox); err != nil {
		return "", err
	}
	event := domain.NewMailboxCreated(mailbox)
	err = app.broker.Publish(event.Event, event)
	if err != nil {
		return "", err
	}
	return mailbox.ID, nil
}

func (app MailboxApp) UserMailboxes(ctx context.Context, userId string) ([]MailboxResponse, error) {
	mailboxes, err := app.dbFactory.MailboxDb.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	if len(mailboxes) == 0 {
		return nil, nil
	}
	result := []MailboxResponse{}
	for _, mb := range mailboxes {
		result = append(result, MailboxResponse{
			Id:     mb.ID,
			UserId: mb.UserID,
			Email:  mb.Email,
		})
	}
	return result, nil
}
