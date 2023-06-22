package out

import "message-delivery/application/domain"

type MessageRepository interface {
	Create(message domain.Message) error
	GetByID(id string) (domain.Message, error)
	Delete(id string) error
}
