package consumers

import (
	"context"
	"time"

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

func (c MailboxConsumer) ConsumeMailboxCreatedEvent(ctx context.Context, event domain.MailboxCreated) error {
	time.Sleep(5 * time.Second)
	err := c.dbFactory.MailboxDb.ActivateMailbox(ctx, event.ID)
	if err != nil {
		return err
	}
	return nil
}
