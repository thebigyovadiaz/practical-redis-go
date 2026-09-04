package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	pubSubChan = "challenge"
)

type Team struct {
	name    string
	client  *redis.Client
	channel *redis.PubSub
}

func getTeams(client *redis.Client) []Team {
	teams := make([]Team, 3)
	allTeams := []string{"team:Grumblebum", "team:Knucklewimp", "team:Snarkdumbthimble"}

	for i, name := range allTeams {
		teams[i].name = name
		teams[i].client = client
	}

	return teams
}

func (team *Team) subscribe() error {
	ctx := context.Background()
	pubSub := team.client.Subscribe(ctx, pubSubChan)

	reply, err := pubSub.Receive(ctx)
	if err != nil {
		return fmt.Errorf("subscribing to channel '%s' failed: %w", pubSubChan, err)
	}

	switch reply.(type) {
	case *redis.Subscription:
		fmt.Println("Successfully type subscription")
	case *redis.Message:
		fmt.Println("The channel is already active and contains messages")
	case *redis.Pong:
		fmt.Println("Let's call it a success")
	default:
		return fmt.Errorf("subscribing to a channel '%s' failed: received a reply of type %T, expected: *redis.Subscription", pubSubChan, reply)
	}

	team.channel = pubSub

	fmt.Printf("%s subscribed to channel '%s'\n", team.name, pubSubChan)
	return nil
}

type Res struct {
	result string
	err    error
}

func (team *Team) receive(ctx context.Context, resChan chan<- Res) {
	defer close(resChan)

	ch := team.channel.Channel()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				// The pub-sub channel has been closed
				return
			}

			resChan <- Res{fmt.Sprintf("%s received challenge '%s'", team.name, msg.Payload), nil}
		case <-ctx.Done():
			resChan <- Res{"", ctx.Err()}
			return
		}
	}
}

func publish(client *redis.Client, challenge string) error {
	ctx := context.Background()
	return client.Publish(ctx, pubSubChan, challenge).Err()
}

func PubSub(client *redis.Client) error {
	ctx := context.Background()

	// Step 1: subscribe each team
	teams := getTeams(client)
	for i := range 3 {
		err := teams[i].subscribe()
		if err != nil {
			return fmt.Errorf("subscribing failed: %w", err)
		}
	}

	// Step 2:
	for i := range int64(5) {
		challenge := client.ZRange(ctx, "challenges", i, i).Val()[0]
		err := publish(client, challenge)
		if err != nil {
			return fmt.Errorf("cannot publish challenge %s: %w", challenge, err)
		}
	}

	// Step 3: receive published messages
	rch := make(chan Res)
	for i := range 3 {
		go teams[i].receive(ctx, rch)
	}

	for msg := range rch {
		if msg.err != nil {
			return fmt.Errorf("cannot receive challenge: %w", msg.err)
		}
		fmt.Println(msg.result)
	}

	defer close(rch)
	return nil
}
