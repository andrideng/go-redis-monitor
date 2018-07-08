package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func redisFoo() {
	// connect to local redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// test redis connection
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// res := client.String("*").Result()
	// determine redis size
	res := client.DBSize()
	fmt.Println(res)

	// get all redis based on key
	a, _ := client.Keys("*").Result()
	for k, v := range a {
		// fmt.Println(k, v)
		fmt.Println(client.Type(v))
		fmt.Println(k, v)
		fmt.Println(client.Get(v))

		fmt.Println("xxxxxxxxxxx")
	}
	// not value: hash, zset, set
	fmt.Println(a)
}
