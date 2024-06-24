package database

import (
	"context"
	"fmt"
	"log"

	"github.com/MachadoMichael/notifications/infra"
	"github.com/go-redis/redis/v8"
)

var EnterpriseRepo *Repo
var client *redis.Client

func Init() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     infra.Config.DbAddress,
		Password: infra.Config.DbPassword,
		DB:       infra.Config.DbName,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(pong)
	client = rdb
	EnterpriseRepo = NewRepo(ctx, rdb)
}

func CloseDb() {
	client.Close()
}
