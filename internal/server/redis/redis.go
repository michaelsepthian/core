package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"time"

	"github.com/spf13/viper"
	"gitlab.com/systeric/internal/chat/backend/core/internal/server/config"
)

var client *redis.Client

func Setup() *redis.Client {
	ctx := context.Background()
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     viper.GetString(config.RedisAddr),
			Password: viper.GetString(config.RedisPassword),
			DB:       viper.GetInt(config.RedisDatabase),
			//TLSConfig: &tls.Config{},
		})
		_, err := client.Ping(ctx).Result()
		if err != nil {
			log.Fatal("Redis Error ", err)
			panic(err)
		}
		log.Print("Redis Started")
	}
	return client
}

func SetObject(rdb *redis.Client, key string, data interface{}, ttl time.Duration) error {
	ctx := context.Background()

	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	set := rdb.Set(ctx, key, string(byteData), ttl)

	if err = set.Err(); err != nil {
		return err
	}
	log.Printf("Set Redis Success, data: %v", string(byteData))
	return nil
}

func GetObject(rdb *redis.Client, key string, objPtr interface{}) error {
	ctx := context.Background()
	get := rdb.Get(ctx, key)

	if err := get.Err(); err != nil {
		return err
	}
	result, err := get.Result()
	data := []byte(result)

	if err != nil {
		return err
	}
	err = json.Unmarshal(data, objPtr)
	if err != nil {
		return err
	}
	log.Printf("Get Redis Success, Result: %v", data)
	return nil
}
