// Copyright 2024-2025 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Описание трассировщика SQL-запросов.
// Реализуется в библиотеке работы с SQL-базой.

package sqldb

import (
	"time"
)

// Tracer - интерфейс описания трассировщика SQL-запросов
type Tracer interface {
	AddSqlDbQuerySpan(s *Span)
}

// Span - данные трассировки SQL-запросов
type Span struct {
	Name    string
	TraceId string
	ConnKey string
	StartDt time.Time
	Err     error
}
