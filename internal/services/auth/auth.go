package auth

import (
	"context"
	"log/slog"
	"sso/internal/domain/models"
	"time"
)

type Auth struct {
	log          *slog.Logger
	userServer   UserServer
	userProvider UserProvider
	appProvider  AppProvider
	tokenTTL     time.Duration
}

type UserServer interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// New returns a new instance of Auth service
func New(
	log *slog.Logger,
	userServer UserServer,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:          log,
		userServer:   userServer,
		userProvider: userProvider,
		appProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string, appID int) (string, error) {
	panic("implement me")
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, password string) (int64, error) {
	panic("implement me")
}

func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	panic("implement me")
}
