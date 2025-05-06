// Описание дополнительных данных сообщения лога
// для message broker

package logger

// MessageBrokerData - данные для логгирования сообщений message broker
type MessageBrokerData struct {
	Action      string `json:"action"`
	TraceId     string `json:"trace_id"`
	MessageId   string `json:"message_id"`
	MessageName string `json:"message_name"`
	Reason      string `json:"reason"`
	Status      string `json:"status"`
	Latency     int64  `json:"latency"`
	Data        string `json:"data"`
}

// NewMessageBrokerData - создание нового объекта для логгирования сообщений message broker
func NewMessageBrokerData() *MessageBrokerData {
	return &MessageBrokerData{}
}
