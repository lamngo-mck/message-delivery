package service

type EventCollector interface {
	collect(event interface{}) error
}
