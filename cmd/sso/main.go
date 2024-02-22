package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	// TODO инициализируем конфиг
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO логгер

	// TODO инициализация приложения (app)

	// TODO запустить gRPC-сервер приложения

}
