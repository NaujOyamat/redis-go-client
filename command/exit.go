package command

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func Exit(client *redis.Client) {
	fmt.Println("Closing connection...")
	if err := client.Close(); err != nil {
		panic(err)
	}

	time.Sleep(1200 * time.Millisecond)

	fmt.Println("Bye!")
	os.Exit(0)
}
