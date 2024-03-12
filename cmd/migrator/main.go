package main

import (
	"errors"
	"flag"
	"fmt"

	// библиотека для миграций
	"github.com/golang-migrate/migrate/v4"
	// драйвер для sqlite3
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	// драйвер для получения миграций из файлов
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// migrationsTable имя таблицы для списка миграций (еще будет отдельная таблица для функциональных тестов)
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied")
}
