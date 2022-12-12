package utils

import (
	"context"
	"fmt"
	"time"
)

const (
	PublishKey = "websocket"
)

// Publish 发布消息到redis
func Publish(cxt context.Context, channel, msg string) error {
	fmt.Println("Publish msg=", msg)
	return RedisClient.Publish(cxt, channel, msg).Err()
}

// Subscribe 订阅redis消息
func Subscribe(cxt context.Context, channel string) (string, error) {
	subscribe := RedisClient.PSubscribe(cxt, channel)
	message, err := subscribe.ReceiveMessage(cxt)
	if err != nil {
		fmt.Println("Subscribe err", err)
		return "", err
	}
	fmt.Println("Subscribe=", message.Payload)
	return message.Payload, err
}

func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	cxt := context.Background()
	RedisClient.Set(cxt, key, val, timeTTL)
}
