// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Описание типов полей для сообщения лога

package logger

type Field = map[string]any

func Any(key string, val any) Field {
	return Field{
		key: val,
	}
}

func Error(err error) Field {
	return Field{
		"error": err.Error(),
	}
}

func String(key string, val string) Field {
	return Field{
		key: val,
	}
}

func Int(key string, val int) Field {
	return Field{
		key: val,
	}
}

func Float(key string, val float64) Field {
	return Field{
		key: val,
	}
}
