package out

import "message-delivery/application/domain"

type MessageRepository struct {
}

func (m MessageRepository) Create(message domain.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) GetByID(id string) (domain.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
