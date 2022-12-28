package command

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-redis/redis"
)

func Execute(client *redis.Client, ctx context.Context, cmds ...string) {
	if len(cmds) > 0 {
		if cmds[0] == "exit" {
			Exit(client)
			return
		}

		if strings.ToLower(cmds[0]) == "ttl" {
			if len(cmds) > 1 {
				TTL(client, cmds[1])
			}
			return
		}

		if strings.ToLower(cmds[0]) == "get" {
			if len(cmds) > 1 {
				Get(client, cmds[1])
			}
			return
		}

		args := []interface{}{}
		for _, cmd := range cmds {
			args = append(args, cmd)
		}

		result, err := client.Do(args...).Result()
		if err != nil {
			if err == redis.Nil {
				fmt.Println("ERROR: key does not exists")
				return
			}

			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		if reflect.TypeOf(result).Kind() == reflect.Slice {
			for _, v := range result.([]interface{}) {
				fmt.Printf("%v\n", v)
			}
		} else {
			fmt.Printf("%v\n", result)
		}
	}
}
