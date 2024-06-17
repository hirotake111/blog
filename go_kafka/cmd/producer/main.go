package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/hirotake111/blog/go_kafka/internal/model"
)

var topicName string = "mytopic"

func main() {
	start := time.Now()
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		log.Fatal("error createting a new producer: ", err)
	}
	defer func() {
		log.Printf("time elapsed: %s\n", time.Since(start))
		p.Close()
	}()

	wg := sync.WaitGroup{}
	go reportResults(p, &wg)
	topicPartition := kafka.TopicPartition{
		Topic:     &topicName,
		Partition: kafka.PartitionAny,
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			m := model.Message{
				Status: model.Status(rand.Intn(3)),
				Value:  fmt.Sprintf("message #%d", i),
			}
			b, err := json.Marshal(m)
			if err != nil {
				log.Print("error marshaling message: ", err.Error())
			} else {
				p.Produce(&kafka.Message{
					TopicPartition: topicPartition,
					Value:          b,
				}, nil)
			}
		}(i)
	}
	// wait for messages before shutting down
	p.Flush(15 * 1000)
	wg.Wait()
}

func reportResults(p *kafka.Producer, wg *sync.WaitGroup) {
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Printf("delivery failed: %+v\n", ev.TopicPartition)
			} else {
				// log.Printf("delivered message to: %+v\n", ev.TopicPartition)
			}

		}
		wg.Done()
	}
}
