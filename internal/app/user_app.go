package app

import (
	"context"

	"github.com/thumperq/golib/messaging"
	"github.com/thumperq/wms/mailbox/internal/domain"
	"github.com/thumperq/wms/mailbox/internal/infrastructure/db"
)

type UserRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"user_name"`
}

type UserApp struct {
	Broker *messaging.Broker
	UserDb db.UserDB
}

func NewUserApp(broker *messaging.Broker, userDb db.UserDB) UserApp {
	return UserApp{
		Broker: broker,
		UserDb: userDb,
	}
}

func (app UserApp) CreateUser(ctx context.Context, request UserRequest) (string, error) {
	user, err := domain.NewUser(request.UserName, request.Password)
	if err != nil {
		return "", err
	}
	if err = app.UserDb.Create(ctx, *user); err != nil {
		return "", err
	}
	return user.ID, nil
}

func (app UserApp) FindUserById(ctx context.Context, id string) (*UserResponse, error) {
	user, err := app.UserDb.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return &UserResponse{
		Id:       user.ID,
		Username: user.UserName,
	}, nil
}
