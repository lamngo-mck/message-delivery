package service

import "message-delivery/application/port/out"

type MessageDeliveryService struct {
	messaging out.MessageBroker
}
