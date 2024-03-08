package util

type Message struct {
	Message string
}

func (e Message) Error() string {
	return e.Message
}

func Error(msg string) Message {
	return Message{Message: msg}
}
