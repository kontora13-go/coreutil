// Copyright 2024-2025 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Описание лвсех возможных логгеров, используемых в приложении

package logger

// Доступные типы логгеров. От типа зависит набор данных и функций записи логов
type Type string

const (
	Runtime       Type = "runtime"        // запись событий в runtime
	HttpClient    Type = "http_client"    // запись событий http-клиента
	HttpServer    Type = "http_server"    // запись событий http-сервера
	MessageBroker Type = "message_broker" // запись событий messege broker
	EventManager  Type = "event_manager"  // запись событий event manager
)

func (l Type) String() string {
	return string(l)
}

// RuntimeLogger - логгер runtime-событий
type RuntimeLogger interface {
	// Debug - runtime debug log write
	Debug(msg string, fields ...map[string]any)
	// Info - runtime info log write
	Info(msg string, fields ...map[string]any)
	// Warning - runtime warning log write
	Warning(msg string, fields ...map[string]any)
	// Error = runtime error log write
	Error(err error, fields ...map[string]any)
	// Fatal - runtime fatal log write
	Fatal(msg string, fields ...map[string]any)

	Close() (err error)
}

// HttpClientLogger - логгер событий http-клиента
type HttpClientLogger interface {
	HttpClientRequest(msg string, data *HttpClientData, fields ...map[string]any)

	Close() (err error)
}

// HttpServerLogger - логгер событий http-сервера
type HttpServerLogger interface {
	HttpServerRequest(msg string, data *HttpServerData, fields ...map[string]any)

	Close() (err error)
}

// MessageBrokerLogger - логгер событий Message Broker
type MessageBrokerLogger interface {
	MessageBrokerAction(msg string, data *MessageBrokerData, fields ...map[string]any)

	Close() (err error)
}

// EventLogger - логгер событий Message Broker
type EventLogger interface {
	IncomingEvent(msg string, data any, fields ...map[string]any)
	OutgoingEvent(msg string, data any, fields ...map[string]any)

	Close() (err error)
}
