package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisCon struct {
	Connection *redis.Client
}

var connection RedisCon

var ctx context.Context

func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	connection.Connection = client

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	return connection.Connection
}

//GetUserID for get userid
func GetUserID(token string) string {
	// value, error := connection.Connection.Get(ctx, token).Result()

	// if error != nil {
	// 	fmt.Println(error)
	// 	return ""
	// }

	return token
}
