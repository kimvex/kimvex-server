package db

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisCon struct {
	Connection redis.Conn
}

var connection RedisCon

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err, "Error conection redis")
	}

	connection.Connection = c

	return connection.Connection
}

//GetUserID for get userid
func GetUserID(token string) string {
	value, error := redis.String(connection.Connection.Do("get", "foo"))
	fmt.Println("value", value)

	if error != nil {
		fmt.Println(error)
		return ""
	}

	return value
}
