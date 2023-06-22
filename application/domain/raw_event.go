package domain

import "time"

type Event interface {
	GetContent() string
}

type RawEvent struct {
	Source    Source
	Content   string
	CreatedAt time.Time
}

func (r RawEvent) GetContent() string {
	return r.Content
}

type Source string

var (
	SourceSDK      Source = "SDK"
	SourceFile     Source = "File"
	SourceInternal Source = "Internal"
)
