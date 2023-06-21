package delivery

type MessageSender interface {
	send(message interface{}) (response interface{}, err error)
}
