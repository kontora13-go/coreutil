// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Пакет организации трассировки запросов в приложении

package trace

import "github.com/google/uuid"

// Id - идентификатор трассировки запроса
type Id = string

// NewId - генерация нового идентификатора трассировки
func NewId() Id {
	return uuid.New().String()
}
