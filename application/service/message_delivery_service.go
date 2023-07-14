package service

import (
	"context"
	"fmt"
	"message-delivery/application/domain"
	"message-delivery/application/port/out"
	"strings"
)

type MessageDeliveryService struct {
	messageRepo  out.MessageRepository
	templateRepo out.TemplateStorage
	sender       out.Sender
}

func (m MessageDeliveryService) Process(ctx context.Context, event domain.Event) error {
	message, err := m.messageRepo.GetByID(ctx, event.ID)
	if err != nil {
		return err
	}
	template, err := m.templateRepo.GetByID(ctx, message.TemplateId)
	if err != nil {
		return err
	}
	content := m.personalizeMessage(template, message.Attributes)
	return m.sender.Send(ctx, content, message.Endpoints)
}

func (m MessageDeliveryService) personalizeMessage(template string, attributes map[string]interface{}) string {
	for k, v := range attributes {
		template = strings.ReplaceAll(template, k, fmt.Sprintf("%v", v))
	}
	return template
}
