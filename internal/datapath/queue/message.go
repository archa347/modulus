package queue

type Message struct {
	shardKey int64
	body     string
}

func NewMessage(shardKey int64, body string) Message {
	return Message{shardKey, body}
}
