package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// @TODO: Endpoint for redis-monitor
// 1. list all key
// 2. add

// Response ...
type Response struct {
	RedisSize int64       `json:"redis_size"`
	Data      interface{} `json:"data"`
}

var client *redis.Client

func connectRedis() (err error) {
	// connect redis
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// test redis connection
	pong, err := client.Ping().Result()
	if err != nil {
		log.Println(err.Error())
		return
	}

	if pong != "" {
		log.Println("Success Connect redis ")
	}

	return
}

func main() {
	err := connectRedis()
	if err != nil {
		log.Println(err)
	}
	// mux
	route := mux.NewRouter()
	route.HandleFunc("/", RedisHandler)

	// Bind to a port and pass our router in
	log.Println("server run on port :8000")
	log.Fatal(http.ListenAndServe(":8000", route))
}

// RedisHandler ...
func RedisHandler(w http.ResponseWriter, r *http.Request) {
	res := &Response{}

	type foo struct {
		Key string `json:"key"`
	}

	// determine the redis size
	redisSize, _ := client.DBSize().Result()
	// fmt.Println(redisSize)
	// set redis size to struct
	res.RedisSize = redisSize

	// get all redis key-value
	redisKey, _ := client.Keys("*").Result()

	m := map[string]interface{}{}
	for _, key := range redisKey {
		val, _ := client.Get(key).Result()
		m[key] = val
	}
	// assign map data
	res.Data = m

	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(res)
	w.Write(b)
}

// set redis
// err = client.Set("world", "jurrasic", 0).Err()
// fmt.Println(err)

// err = client.Set("chiki", "choko", 0).Err()
// fmt.Println(err)

// err = client.Set("chiki", "chuku", 0).Err()
// fmt.Println(err)

// err = client.Set("chuku", "choco", 0).Err()
// fmt.Println(err)
