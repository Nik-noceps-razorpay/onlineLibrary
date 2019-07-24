package DB_connection

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/olivere/elastic.v5"
)

var DbRedis *redis.Client
var DbElastic *elastic.Client

func InitRedisDB() {
	DbRedis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := DbRedis.Ping().Result()

	if err != nil {
		fmt.Println("error in InitDB() function")
		fmt.Println(pong, "the error is :", err)
	} else {
		fmt.Println("Redis database initialized without any errors")
	}


}

func InitElasticDB() {
	ctx := context.Background()

	var err error

	DbElastic, err = elastic.NewClient()

	fmt.Println("dbelastic 1", DbElastic)

	if err != nil {
		fmt.Println("\n error in creating NewClient for elastic")
		panic(err)
	}

	info, code, err1 := DbElastic.Ping("http://127.0.0.1:9200").Do(ctx)

	if err1 != nil {
		fmt.Println("error in  client ping")
		panic(err1)
	} else {
		fmt.Println("Elastic database initialized without any errors")
	}

	fmt.Println("Elasticsearch returned with code", code, " and version", info.Version.Number)


}

