// Copyright 2024-2025 Kontora13. All rights reserved.
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

// Close - функция закрытия логгера
func (l *DefaultLogger) Close() (err error) {
	return nil
}

// Debug - функция логгирования отладочных сообщений
func (l *DefaultLogger) Debug(msg string, fields ...map[string]any) {
}

// Info - функция логгирования информационных сообщений
func (l *DefaultLogger) Info(msg string, fields ...map[string]any) {
}

// Warning - функция логгирования предупреждающих сообщений
func (l *DefaultLogger) Warning(msg string, fields ...map[string]any) {
}

// Error - функция логгирования ошибок
func (l *DefaultLogger) Error(err error, fields ...map[string]any) {
}

// Fatal - функция логгирования фатальных сообщений
func (l *DefaultLogger) Fatal(msg string, fields ...map[string]any) {
}

// HttpClientRequest - функция логгирования событий http-клиента
func (l *DefaultLogger) HttpClientRequest(msg string, data *HttpClientData, fields ...map[string]any) {
}

// HttpServerRequest - функция логгирования событий http-сервера
func (l *DefaultLogger) HttpServerRequest(msg string, data *HttpServerData, fields ...map[string]any) {
}

// MessageBrokerAction - функция логгирования событий брокера сообщений	sa
func (l *DefaultLogger) MessageBrokerAction(msg string, data *MessageBrokerData, fields ...map[string]any) {
}
