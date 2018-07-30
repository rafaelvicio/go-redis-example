package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Key value: ", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key 2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2 value: ", val2)
	}
}
