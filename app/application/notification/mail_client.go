package notification

import "context"

type MailClient interface {
	Send(ctx context.Context, content []MailContent) error
}

type MailContent struct {
	To      string
	Subject string
	Body    string
}
