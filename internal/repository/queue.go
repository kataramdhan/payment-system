package repository

import "github.com/hibiken/asynq"

func NewRedisClient() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr: "localhost:6379",
	})
}
