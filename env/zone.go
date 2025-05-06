// Copyright 2024-2025 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Работа с зоной приложения, в которой оно запускается

package env

import "fmt"

type Zone string

const (
	ZoneLocal      Zone = "LOCAL"       // локальная разработка на машине разработчика
	ZoneDev        Zone = "DEV"         // тестирование приложения
	ZoneTest       Zone = "TEST"        // интеграционные тесты приложения
	ZoneStage      Zone = "STAGE"       // предпродуктивная среда
	ZoneProd       Zone = "PROD"        // продуктивная среда
	ZoneNotDefined Zone = "NOT_DEFINED" // не определена
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
