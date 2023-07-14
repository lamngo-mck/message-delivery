package delivery

import (
	"context"
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"message-delivery/common"
)

type Email struct {
	Client *mailjet.Client
}

func NewEmail(config common.EmailConfig) Email {
	return Email{
		Client: mailjet.NewMailjetClient(config.PublicKey, config.PrivateKey),
	}
}

func (e *Email) Send(ctx context.Context, content string, endpoints []string) error {
	data := mailjet.MessagesV31{
		Info: []mailjet.InfoMessagesV31{
			{
				From: &mailjet.RecipientV31{
					Email: "",
					Name:  "",
				},
				To:       e.transformRecipients(endpoints),
				Subject:  "TEST",
				TextPart: content,
			},
		},
		SandBoxMode: false,
	}
	response, err := e.Client.SendMailV31(&data)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func (e *Email) transformRecipients(endpoints []string) *mailjet.RecipientsV31 {
	var recipients mailjet.RecipientsV31
	for _, endpoint := range endpoints {
		recipients = append(recipients, mailjet.RecipientV31{
			Email: endpoint,
			Name:  endpoint,
		})
	}
	return &recipients
}
