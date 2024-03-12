package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"sso/internal/lib/logger/handlers/slogpretty"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO инициализируем конфиг
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO логгер
	log := setupLogger(cfg.Env)
	log.Info(
		"starting application",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.GRPC.Port),
		slog.Any("configuration", cfg),
	)

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	go application.GRPCSer.MustRun()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	// ждет сигнал от операционной системы (тут SIGTERM или SIGINT) и передает в канал
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// фукнция висит и ждет значения из канала
	sign := <-stop
	log.Info("application stoppping", slog.String("signal", sign.String()))
	// как только получили, то гасим сервер
	application.GRPCSer.Stop()
	log.Info("application stop")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	// вид лога зависит от окружения
	switch env {
	case envLocal:
		//// текст удобен при просмотре
		//log = slog.New(
		//    slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		//)
		log = setupPrettySlog()
	case envDev:
		// json удобен для метрик
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
