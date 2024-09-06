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
	mailboxDb db.MailboxDB
}

func NewMailboxConsumer(mailboxDb db.MailboxDB) MailboxConsumer {
	return MailboxConsumer{
		mailboxDb: mailboxDb,
	}
}

func (c MailboxConsumer) Handle(ctx context.Context, msg messaging.Message) error {
	time.Sleep(5 * time.Second)
	var event domain.MailboxCreated
	err := json.Unmarshal(msg.Data, &event)
	if err != nil {
		return err
	}
	err = c.mailboxDb.ActivateMailbox(ctx, event.ID)
	if err != nil {
		return err
	}
	return nil
}
