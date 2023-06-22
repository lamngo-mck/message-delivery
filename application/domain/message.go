package domain

import "time"

type Message struct {
	Content   Content
	CreatedAt time.Time
}

type Content interface {
	GetDetail() interface{}
}

type WebPushMessage struct {
}

func (w WebPushMessage) GetDetail() interface{} {
	return nil
}

type EmailMessage struct {
}

func (e EmailMessage) GetDetail() interface{} {
	return nil
}

type OnsiteMessage struct {
}

func (o OnsiteMessage) GetDetail() interface{} {
	return nil
}
