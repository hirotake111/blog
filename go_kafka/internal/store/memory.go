package store

import (
	"context"
	"fmt"
	"log"

	"github.com/hirotake111/blog/go_kafka/internal/model"
)

type InMemoryStore struct {
	store []*model.Message
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		store: make([]*model.Message, 0),
	}
}

func (s *InMemoryStore) Add(ctx context.Context, msg *model.Message) error {
	s.store = append(s.store, msg)
	return nil
}

func (s *InMemoryStore) Report() error {
	var count, complete, failed, inProgress int
	for _, m := range s.store {
		count++
		switch m.Status {
		case model.Complete:
			complete++
		case model.Failed:
			failed++
		case model.InProgress:
			inProgress++
		}
	}
	if count >= 10 {
		log.Print("==== Summery of message streaming ====")
		fmt.Printf(
			"delivered %d%%, complete %d%%, failed %d%%, in progress %d%%\n",
			count/10,
			complete*100/count,
			failed*100/count,
			inProgress*100/count,
		)
	}
	return nil
}
