// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Пакет работы с SQL базами данных (MS SQL, PostgreSQL, SQLite и т.д.)
// Описывает интерфейс взаимодействия с подключением
// и содержит структуры для описания запросов к БД

package sqldb

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Connection - интерфейс описания подключения к БД
type Connection interface {
	Ping(ctx context.Context) error
	GetConnection() *sqlx.DB
	GetDriverName() string
	GetConnectionKey() string
	GetDatabaseName() string
	Close() error

	Exec(ctx context.Context, query *Query) (result sql.Result, err error)
	Select(ctx context.Context, dest any, query *Query) (err error)
	Get(ctx context.Context, dest any, query *Query) (err error)
	SelectMaps(ctx context.Context, query *Query) (result []map[string]any, err error)
	GetMap(ctx context.Context, query *Query) (result map[string]any, err error)
}

// ConnectionPool - интерфейс описаниия пула подключений к БД
type ConnectionPool interface {
	Add(name string, conn Connection) error
	Get(name string) (Connection, error)
	Close() error
}
