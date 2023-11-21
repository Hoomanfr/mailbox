package consumers

import (
	"context"

	"github.com/thumperq/golib/logging"
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
	logging.TraceLogger(ctx).
		Info().
		Str("event", event.Event).
		Str("id", event.ID).
		Str("user_id", event.UserID).
		Str("email", event.Email).
		Time("created_at", event.CreatedAt).
		Msgf("%s event consumed", event.Event)
	return nil
}
