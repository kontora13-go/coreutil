// Copyright 2024-2025 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

package trace

import "time"

// Span - данные трассировки
type Span struct {
	startDt time.Time
	name    string
}

func NewSpan(name string, startDt time.Time) Span {
	return Span{
		name:    name,
		startDt: startDt,
	}
}
