// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Описание работы с метриками SQL-запросов.
// Реализуется в библиотеке работы с SQL-базой.

package sqldb

import (
	"time"
)

// Metrics - интерфейс описания объекта для записи метрик SQL-запросов
type Metrics interface {
	SetSqlDbQueryMetrics(m *Metric)
}

// Metric - данные метрики запроса
type Metric struct {
	Name    string
	ConnKey string
	StartDt time.Time
	Err     error
}
