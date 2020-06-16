package queue

type Queue interface {
	push(Message)
	listShards()
}

type ShardMetadata struct {
	size int64
	 int64
}
