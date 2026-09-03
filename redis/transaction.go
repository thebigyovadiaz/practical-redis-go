package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Transaction(client *redis.Client) error {
	ctx := context.Background()

	_, err := client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		// Move Sykios to team Grumblebum
		err := pipe.HSet(ctx, "player:1", "team", "Grumblebum").Err()
		if err != nil {
			return err
		}

		// Move Nidios to team Grumblebum
		err = pipe.HSet(ctx, "player:2", "team", "Grumblebum").Err()
		if err != nil {
			return err
		}

		// Move Belaeos to team Grumblebum
		err = pipe.HSet(ctx, "player:4", "team", "Grumblebum").Err()
		if err != nil {
			return err
		}

		// Move Tiaitia to team Knucklewimp
		err = pipe.HSet(ctx, "player:3", "team", "Knucklewimp").Err()
		if err != nil {
			return err
		}

		// Team update: remove Belaeos from team Knucklewimp
		err = pipe.SRem(ctx, "team:Knucklewimp", "Belaeos").Err()
		if err != nil {
			return err
		}

		// Team update: add Tiaitia to team Knucklewimp
		err = pipe.SAdd(ctx, "team:Knucklewimp", "Tiaitia").Err()
		if err != nil {
			return err
		}

		// Add team Grumblebum
		err = pipe.SAdd(ctx, "team:Grumblebum", "Sykios", "Nidios", "Belaeos").Err()
		if err != nil {
			return err
		}

		// Remove team Dorkfoot. A set is removed by removing all elements.
		err = pipe.SRem(ctx, "team:Dorkfoot", "Sykios", "Nidios", "Tiaitia").Err()
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Printf("TxPipelined failed: %v\n", err)
	}

	fmt.Printf("Sykios's new team: %s\n", client.HGet(ctx, "player:1", "team").Val())
	fmt.Printf("Belaeos's new team: %s\n", client.HGet(ctx, "player:4", "team").Val())
	fmt.Printf("Tiaitia's new team: %s\n", client.HGet(ctx, "player:3", "team").Val())
	fmt.Printf("Team Grumblebum: %s\n", client.SMembers(ctx, "team:Grumblebum").Val())
	fmt.Printf("Team Knucklewimp: %s\n", client.SMembers(ctx, "team:Knucklewimp").Val())

	return nil
}
