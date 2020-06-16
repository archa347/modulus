package memory

type Queue struct {
	shards map[int64]Shard
}

func New() Queue {

}
