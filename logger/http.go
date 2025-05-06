// Copyright 2024-2025 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Описание дополнительных данных сообщения лога
// для http- сервера и клиента

package logger

// HttpClientData - данные для логгирования http-клиента
type HttpClientData struct {
	StatusCode int              `json:"status_code"`
	Path       string           `json:"path"`
	Host       string           `json:"host"`
	Url        string           `json:"url"`
	Method     string           `json:"method"`
	RemoteAddr string           `json:"remote_addr"`
	Latency    int64            `json:"latency"`
	TraceId    string           `json:"trace_id"`
	Request    HttpRequestData  `json:"request"`
	Response   HttpResponseData `json:"response"`
}

func NewHttpClientData() *HttpClientData {
	return &HttpClientData{
		Request:  HttpRequestData{},
		Response: HttpResponseData{},
	}
}

// HttpServerData - данные для логгирования http-сервера
type HttpServerData struct {
	StatusCode int              `json:"status_code"`
	Path       string           `json:"path"`
	Host       string           `json:"host"`
	Url        string           `json:"url"`
	Method     string           `json:"method"`
	RemoteAddr string           `json:"remote_addr"`
	Latency    int64            `json:"latency"`
	TraceId    string           `json:"trace_id"`
	Request    HttpRequestData  `json:"request"`
	Response   HttpResponseData `json:"response"`
}

func NewHttpServerData() *HttpServerData {
	return &HttpServerData{
		Request:  HttpRequestData{},
		Response: HttpResponseData{},
	}
}

// HttpRequestData - данные http-запроса
type HttpRequestData struct {
	Header map[string]string `json:"header"`
	Query  map[string]string `json:"query"`
	Body   string            `json:"body"`
}

// AddHeader - добавление информации о http-заголовке
func (r *HttpRequestData) AddHeader(key string, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}

	r.Header[key] = value
}

// AddQuery - добавление информации о параметре http-запроса
func (r *HttpRequestData) AddQuery(key string, value string) {
	if r.Query == nil {
		r.Query = make(map[string]string)
	}

	r.Query[key] = value
}

// HttpResponseData - данные ответа на http-запрос
type HttpResponseData struct {
	Header map[string]string `json:"header"`
	Body   string            `json:"body"`
}

// AddHeader - добавление информации о http-заголовке
func (r *HttpResponseData) AddHeader(key string, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}

	r.Header[key] = value
}
