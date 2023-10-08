package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/JakubC-projects/work-telegram-bot/src/config"
	"github.com/JakubC-projects/work-telegram-bot/src/models"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Url,
		Password: config.C.Redis.Password,
		DB:       0,
	})
}

func SaveRegistration(ctx context.Context, chatID int64, messageID int, reg models.RegistrationState) error {
	regBody, err := json.Marshal(reg)
	if err != nil {
		return fmt.Errorf("cannot marshal registration: %w", err)
	}

	return client.Set(ctx, regKey(chatID, messageID), regBody, time.Hour).Err()
}

func GetRegistration(ctx context.Context, chatID int64, messageID int) (models.RegistrationState, error) {
	var registration models.RegistrationState

	res, err := client.Get(ctx, regKey(chatID, messageID)).Result()
	if err == redis.Nil {
		return registration, ErrNotFound
	}
	if err != nil {
		return registration, err
	}

	if err := json.Unmarshal([]byte(res), &registration); err != nil {
		return registration, fmt.Errorf("cannot unmarshal registration: %w", err)
	}

	return registration, nil
}

func regKey(chatID int64, messageID int) string {
	return fmt.Sprintf("%d:%d", chatID, messageID)
}

var ErrNotFound = errors.New("not found")
