package db

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// RedisCon struct for redis
type RedisCon struct {
	Connection *redis.Pool
}

var connection RedisCon

// RedisConnect connection
func RedisConnect() {
	c := redis.Pool{
		MaxActive: 4000,
		MaxIdle:   300, // adjust to your needs
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				fmt.Println(err.Error())
				errInfo := err
				fmt.Println(errInfo)
				return nil, err
			}
			// if _, err := c.Do("AUTH", cachekey); err != nil {
			// 	fmt.Println(err.Error())
			// 	c.Close()
			// 	return nil, err
			// }
			fmt.Println("redis connection success")
			return c, err

		},
	}

	connection.Connection = &c
}

//GetUserID for get userid
func GetUserID(token string) string {
	y := connection.Connection.Get()
	value, error := redis.String(y.Do("GET", token))
	fmt.Println("value", value)
	y.Close()
	if error != nil {
		fmt.Println(error, "que?")
		return ""
	}

	return value
}

//SetUserID for get userid
func SetUserID(token string, userID string) {
	y := connection.Connection.Get()
	y.Do("SET", token, userID)
	y.Close()
}

//DeleteUserID for get userid
func DeleteUserID(token string) {
	y := connection.Connection.Get()
	y.Do("DEL", token)
	y.Close()
}
