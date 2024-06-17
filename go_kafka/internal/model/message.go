package model

type Status int

const (
	Failed = iota
	Complete
	InProgress
)

type Message struct {
	Status Status `json:"status"`
	Value  string `json:"value"`
}
