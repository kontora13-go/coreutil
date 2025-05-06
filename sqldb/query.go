// Copyright 2024-2025 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

package sqldb

import "time"

// Query - структура для описания SQL-запроса
type Query struct {
	Label   string
	Text    string
	Arg     any
	Trace   Tracer
	Timeout time.Duration
}
