package suite

import (
	"context"
	"net"
	"sso/internal/config"
	"strconv"
	"testing"

	ssov1 "github.com/vrnvgasu/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcHost = "localhost"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient // клиент для работы с grpc сервером
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper() //при ошибке в стеке вызовов функция теста не указывается, как финальная
	t.Parallel()

	cfg := config.MustLoadByPath("../config/local.yaml")

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	// grpc клиент
	cc, err := grpc.DialContext(
		context.Background(),
		grpcAdders(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()), // говорим, что будем использовать небезопасное соединение
	)
	if err != nil {
		t.Fatal("grpc ser return nil, nilver connection failed: %w", err)
	}

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: ssov1.NewAuthClient(cc),
	}
}

func grpcAdders(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
