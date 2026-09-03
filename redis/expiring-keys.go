package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func ExpiringKeys(client *redis.Client) error {
	ctx := context.Background()

	// Add a temporary player
	err := client.HSet(ctx, "player:10", "name", "Crymyios", "score", 0, "team", "Knucklewimp", "challenges_completed", 0).Err()
	if err != nil {
		return fmt.Errorf("cannot set player:10: %w", err)
	}

	// Set an expirationtime for player:10
	playerExpired := client.Expire(ctx, "player:10", time.Second).Val()
	if !playerExpired {
		return fmt.Errorf("cannot set expiration time for player:10")
	}

	// Get player:10
	for range 3 {
		val, err := client.HGet(ctx, "player:10", "name").Result()
		if err != nil {
			fmt.Printf("player:10 has expired: %v\n", err)
			return nil
		}

		fmt.Printf("player:10's name: %s\n", val)
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}
