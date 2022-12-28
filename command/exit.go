package command

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func Exit(ctx context.Context, client *redis.Client) {
	fmt.Println("Closing connection...")
	if err := client.WithContext(ctx).Close(); err != nil {
		panic(err)
	}

	time.Sleep(1200 * time.Millisecond)

	fmt.Println("Bye!")
	os.Exit(0)
}
