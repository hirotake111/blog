package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/hirotake111/blog/go_kafka/internal/model"
	"github.com/hirotake111/blog/go_kafka/internal/store"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal("error creating new consmer: ", err)
	}
	if err := c.Subscribe(topicName, nil); err != nil {
		log.Fatal("error initilating a subscription: ", err)
	}
	loop(context.Background(), c, store.NewInMemoryStore())

}

const (
	topicName = "mytopic"
)

type Store interface {
	Add(ctx context.Context, msg *model.Message) error
	Report() error
}

func loop(ctx context.Context, c *kafka.Consumer, store Store) {
	log.Println("==== consumer started ====")
	for {
		km, err := c.ReadMessage(5 * time.Second)
		if err == nil {
			msg := new(model.Message)
			if err := json.Unmarshal(km.Value, msg); err != nil {
				log.Printf("error storing message %s\n", err.Error())
				continue
			}
			log.Printf("Message on %s: %+v\n", km.TopicPartition, msg)
			if err := store.Add(ctx, msg); err != nil {
				log.Printf("error storing message %s\n", err.Error())
				continue
			}
		} else if err.(kafka.Error).IsTimeout() {
			store.Report()
		} else {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			log.Printf("Consumer error: %v (%v)\n", err, km)
		}
	}
}
