package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/thumperq/messaging/mailbox/internal/common"
)

type Mailbox struct {
	ID        string
	UserID    string
	Email     string
	CreatedAt time.Time
}

func NewMailbox(userId string, email string) (*Mailbox, error) {
	if strings.TrimSpace(userId) == "" {
		return nil, errors.New(common.ErrInvalidUserId)
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New(common.ErrInvalidEmail)
	}
	return &Mailbox{
		ID:        uuid.New().String(),
		UserID:    userId,
		Email:     email,
		CreatedAt: time.Now(),
	}, nil
}
