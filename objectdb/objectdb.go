// Copyright 2024-2025 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Описание интерфейса работы с объектным хранилищем

package objectdb

type Connection interface {
	// TODO: запись, чтение объекта
	Close() error
}
