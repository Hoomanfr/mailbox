package consumers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/thumperq/golib/messaging"
	"github.com/thumperq/wms/mailbox/internal/domain"
	"github.com/thumperq/wms/mailbox/internal/infrastructure/db"
)

type MailboxConsumer struct {
	dbFactory db.DbFactory
}

func NewMailboxConsumer(dbFactory db.DbFactory) MailboxConsumer {
	return MailboxConsumer{
		dbFactory: dbFactory,
	}
}

func (c MailboxConsumer) ConsumeMailboxTopics(ctx context.Context, event domain.MailboxCreated) error {
	time.Sleep(5 * time.Second)
	err := c.dbFactory.MailboxDb.ActivateMailbox(ctx, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c MailboxConsumer) ConsumeMailboxCreatedEvent(ctx context.Context, msg messaging.Message) error {
	time.Sleep(5 * time.Second)
	var event domain.MailboxCreated
	err := json.Unmarshal(msg.Data, &event)
	if err != nil {
		return err
	}
	err = c.dbFactory.MailboxDb.ActivateMailbox(ctx, event.ID)
	if err != nil {
		return err
	}
	return nil
}
