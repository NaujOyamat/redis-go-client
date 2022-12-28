package command

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

func Get(ctx context.Context, client *redis.Client, key string) {
	result, err := client.WithContext(ctx).Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("ERROR: key does not exists")
			return
		}

		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	pretty, err := prettyJsonString(result)
	if err != nil {
		fmt.Println(result)
	} else {
		fmt.Println(pretty)
	}
}

func prettyJsonString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
