package service

import "message-delivery/application/port/out"

type MessageDeliveryService struct {
	repo out.MessageRepository
}
