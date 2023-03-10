package command

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

func TTL(ctx context.Context, client *redis.Client, key string) {
	result, err := client.WithContext(ctx).TTL(key).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("ERROR: key does not exists")
			return
		}

		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	val := result.String()

	if val == "-2s" {
		val = "key don't exists"
	}
	if val == "-1s" {
		val = "key don't have ttl"
	}

	fmt.Printf("TTL: %s\n", val)
}
