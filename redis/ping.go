package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func PingRequest(client *redis.Client) error {
	// Only need a background context
	ctx := context.Background()
	// Ping Redis Server
	fmt.Println(client.Ping(ctx))

	// Get client info
	info, err := client.ClientInfo(ctx).Result()
	if err != nil {
		return fmt.Errorf("method ClientInfo failed: %v", err)
	}

	fmt.Printf("%#v\n", info)
	return nil
}
