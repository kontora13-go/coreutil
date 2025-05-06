// Copyright 2024-2025 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Работа с версией приложения

package env

import (
	"fmt"

	"github.com/hashicorp/go-version"
)

// EmptyVersion - "пустая" версия
var EmptyVersion = Version{version.Version{}}

// Version - объект для хранения версии приложения
type Version struct {
	version.Version
}

// String - строковое представление версии
func (v *Version) String() string {
	return v.Version.String()
}

// NewVersion - парсинг версии из строки
func NewVersion(s string) (Version, error) {
	v, err := version.NewVersion(s)

	if err != nil {
		return EmptyVersion, fmt.Errorf("[environment] некорректная версия приложения: %s", s)
	}
	return Version{
		Version: *v,
	}, nil
}
