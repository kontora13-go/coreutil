// Copyright 2024-2025 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Интерфейс, описывающий взаимодействие с Key-Value хранилищем

package keyvaluedb

import (
	"context"
)

// Connection - интерфейс подключения к Key-Value
type Connection interface {
	Write(ctx context.Context, key string, value any) error
	Read(ctx context.Context, key string, value any) error
	Close() error
}
