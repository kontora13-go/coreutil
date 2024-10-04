// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Описание логгера-заглушки

package logger

// Default - создание пустого логгера
func Default() *DefaultLogger {
	return &DefaultLogger{}
}

// DefaultLogger - пустой логгер по-умолчанию,
// реализующий интерфейсы логгеров всех типов
type DefaultLogger struct {
}

func (l *DefaultLogger) Close() (err error) {
	return nil
}

func (l *DefaultLogger) Debug(msg string, fields ...map[string]any) {
}

func (l *DefaultLogger) Info(msg string, fields ...map[string]any) {
}

func (l *DefaultLogger) Warning(msg string, fields ...map[string]any) {
}

func (l *DefaultLogger) Error(err error, fields ...map[string]any) {
}

func (l *DefaultLogger) Fatal(msg string, fields ...map[string]any) {
}

func (l *DefaultLogger) HttpClientRequest(msg string, data *HttpClientData, fields ...map[string]any) {
}

func (l *DefaultLogger) HttpServerRequest(msg string, data *HttpServerData, fields ...map[string]any) {
}

func (l *DefaultLogger) IncomingEvent(msg string, data any, fields ...map[string]any) {
}

func (l *DefaultLogger) OutgoingEvent(msg string, data any, fields ...map[string]any) {
}
