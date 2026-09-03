package redis

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func SetLang(lang string) language.Tag {
	switch strings.ToUpper(lang) {
	case "EN":
		return language.English
	case "ES":
		return language.Spanish
	case "FR":
		return language.French
	default:
		return language.English
	}
}

func GetAndSet(client *redis.Client, key , lang string) error {
	ctx := context.Background()

	quest, err := client.Get(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("cannot get quest: %w", err)
	}

	quest = cases.Title(SetLang(lang)).String(quest)

	err = client.Set(ctx, key, quest, 0).Err()
	if err != nil {
		return fmt.Errorf("cannot update quest: %w", err)
	}

	fmt.Printf("Quest is now: %s\n", client.Get(ctx, key).Val())
	return nil
}
