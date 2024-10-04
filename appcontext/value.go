// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

package appcontext

// AppCtxValue - структура для хранения произвольных значений
type AppCtxValue struct {
	m map[string]any
}

// NewAppCtxValue - создание новой value-структуры
// m - исходная map
func NewAppCtxValue(m map[string]any) AppCtxValue {
	return AppCtxValue{
		m: m,
	}
}

// Add - добавление значения
func (v AppCtxValue) Add(key string, val any) {
	v.m[key] = val
}

// Value - извлечение значения по ключу
func (v AppCtxValue) Value(key string) any {
	return v.m[key]
}

// Exist - проверка существования значения по ключу
func (v AppCtxValue) Exist(key string) bool {
	_, exist := v.m[key]
	return exist
}
