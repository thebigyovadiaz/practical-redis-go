package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Pipeline(client *redis.Client) error {
	ctx := context.Background()

	_, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		err := pipe.HSet(ctx, "player:7", "score", 15, "challenges_completed", 1).Err()
		if err != nil {
			return err
		}

		err = pipe.HSet(ctx, "player:8", "score", 18, "challenges_completed", 2).Err()
		if err != nil {
			return err
		}

		err = pipe.HSet(ctx, "player:9", "score", 12, "challenges_completed", 1).Err()
		return err
	})

	if err != nil {
		return fmt.Errorf("pipeline failed: %v", err)
	}

	fmt.Printf("Player 7's score: %s, challenges completed: %s\n",
		client.HGet(ctx, "player:7", "score").Val(),
		client.HGet(ctx, "player:7", "challenges_completed").Val())

	fmt.Printf("Player 8's score: %s, challenges completed: %s\n",
		client.HGet(ctx, "player:8", "score").Val(),
		client.HGet(ctx, "player:8", "challenges_completed").Val())

	fmt.Printf("Player 9's score: %s, challenges completed: %s\n",
		client.HGet(ctx, "player:9", "score").Val(),
		client.HGet(ctx, "player:9", "challenges_completed").Val())

	return nil
}
