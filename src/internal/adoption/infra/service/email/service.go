package email

import (
	"context"
	"fmt"
	"time"

	"github.com/college-redberry/open-adoption/internal/adoption/infra/constants"
	"github.com/mailgun/mailgun-go/v4"
)

type MailgunService struct {
	client *mailgun.MailgunImpl
	sender string
}

func New() *MailgunService {
	mg := mailgun.NewMailgun(constants.MAILGUN_DOMAIN, constants.MAILGUN_API_KEY)
	return &MailgunService{
		client: mg,
		sender: constants.MAILGUN_SENDER,
	}
}

func (m *MailgunService) Send(to, subject, body string) error {
	message := m.client.NewMessage(
		m.sender,
		subject,
		body,
		to,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, _, err := m.client.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
