package command

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path"

	"github.com/go-redis/redis"
)

func Keys(ctx context.Context, client *redis.Client, pattern string) {
	fmt.Println("Getting keys...")
	var total uint64

	iter := client.WithContext(ctx).Scan(0, pattern, 0).Iterator()
	for iter.Next() {
		total++
		fmt.Println(iter.Val())
	}

	if err := iter.Err(); err != nil {
		if err == redis.Nil {
			fmt.Println("ERROR: pattern does not exists")
			return
		}

		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	fmt.Printf("Total: %d\n", total)
}

func FKeys(ctx context.Context, client *redis.Client, pattern string) {
	pf := "/tmp/keys.rd"
	if c, e := os.Getwd(); e == nil {
		pf = path.Join(c, "keys.rd")
	}

	fmt.Printf("Creating file (%s)...\n", pf)
	f, err := os.Create(pf)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	defer f.Close()

	wf := bufio.NewWriter(f)
	var total uint64

	fmt.Println("Getting keys...")
	iter := client.WithContext(ctx).Scan(0, pattern, 0).Iterator()
	for iter.Next() {
		total++
		fmt.Println(iter.Val())
		_, err := wf.WriteString(fmt.Sprintf("%s\n", iter.Val()))
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}
	}

	if err := iter.Err(); err != nil {
		if err == redis.Nil {
			fmt.Println("ERROR: pattern does not exists")
			return
		}

		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	fmt.Printf("Total: %d\n", total)
}
