package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-redis/redis"
	"github.com/naujoyamat/redis-go-client/command"
)

var (
	client *redis.Client

	host = flag.String("h", "localhost", "h=localhost")
	port = flag.Int("p", 6379, "p=6379")
	pwd  = flag.String("pwd", "", "pwd=password")
	db   = flag.Int("db", 0, "db=0")
)

func main() {
	flag.Parse()

	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", *host, *port),
		Password:     *pwd,
		DB:           *db,
		ReadTimeout:  -1,
		WriteTimeout: 1 * time.Minute,
	})

	if client == nil {
		panic(fmt.Errorf("redis client nil..."))
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	loop(ctx, client)
}

func prompt() {
	fmt.Printf("%s:%d[%d]> ", *host, *port, *db)
}

func loop(ctx context.Context, rdb *redis.Client) {
	scanner := bufio.NewScanner(os.Stdin)

	prompt()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if scanner.Scan() {
				line := strings.Trim(strings.TrimRight(scanner.Text(), "\n"), " ")
				if line != "" {
					cmds := strings.Split(line, " ")
					command.Execute(client, ctx, cmds...)
				}
			}
			prompt()
		}
	}
}
