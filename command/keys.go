package command

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Keys(client *redis.Client, pattern string) {
	result, err := client.Keys(pattern).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("ERROR: pattern does not exists")
			return
		}

		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	for _, s := range result {
		fmt.Println(s)
	}
}
