package app

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
	Status string `json:"status"`
}

type MailboxApp struct {
	Broker    *messaging.Broker
	MailboxDb db.MailboxDB
}

func NewMailboxApp(broker *messaging.Broker, mailboxDb db.MailboxDB) MailboxApp {
	return MailboxApp{
		Broker:    broker,
		MailboxDb: mailboxDb,
	}
}

func (app MailboxApp) CreateMailbox(ctx context.Context, request MailboxRequest) (string, error) {
	mailbox, err := domain.NewMailbox(request.UserId, request.Email)
	if err != nil {
		return "", err
	}
	if err := app.MailboxDb.Create(ctx, *mailbox); err != nil {
		return "", err
	}
	event := domain.NewMailboxCreated(mailbox)
	err = app.Broker.Publish(event.Event, event)
	if err != nil {
		return "", err
	}
	return mailbox.ID, nil
}

func (app MailboxApp) UserMailboxes(ctx context.Context, userId string) ([]MailboxResponse, error) {
	mailboxes, err := app.MailboxDb.FindByUserId(ctx, userId)
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
			Status: mb.Status,
		})
	}
	return result, nil
}
