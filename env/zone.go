// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Работа с зоной приложения, в которой оно запускается

package env

import "fmt"

type Zone string

const (
	ZoneLocal      Zone = "LOCAL"
	ZoneDev        Zone = "DEV"
	ZoneTest       Zone = "TEST"
	ZoneStage      Zone = "STAGE"
	ZoneProd       Zone = "PROD"
	ZoneNotDefined Zone = "NOT_DEFINED"
)

// NewZone - парсинг зоны из строки
func NewZone(v string) Zone {
	return Zone(v)
}

// IsValid - проверка валидности зоны
func (z Zone) IsValid() error {
	switch z {
	case ZoneLocal, ZoneDev, ZoneTest, ZoneStage, ZoneProd:
		return nil
	}
	return fmt.Errorf("[environment] неизвестная среда запуска приложения: %s", z)
}

// String - строковое представление зоны
func (z Zone) String() string {
	return string(z)
}
