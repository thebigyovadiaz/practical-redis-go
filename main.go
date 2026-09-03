package main

import (
	"fmt"

	"github.com/thebigyovadiaz/practical-redis-go/redis"
)

func main() {
	// Create new redis client
	client := redis.NewRedisClient("localhost:6379", 0)
	defer client.Close()

	// Ping request to redis client
	/*err := redis.PingRequest(client)
	if err != nil {
		fmt.Printf("Ping failed: %v\n", err)
	}*/

	fmt.Println()
	fmt.Println()

	// Get and Set a register in db
	err := redis.GetAndSet(client, "description", "es")
	if err != nil {
		fmt.Printf("GetAndSet failed: %v\n", err)
	}

	fmt.Println()
	fmt.Println()

	err = redis.ExpiringKeys(client)
	if err != nil {
		fmt.Printf("ExpiringKeys failed: %v\n", err)
	}

	fmt.Println()
	fmt.Println()

	err = redis.Pipeline(client)
	if err != nil {
		fmt.Printf("Pipeline failed: %v\n", err)
	}

	fmt.Println()
	fmt.Println()

	err = redis.Transaction(client)
	if err != nil {
		fmt.Printf("Transaction failed: %v\n", err)
	}
}
