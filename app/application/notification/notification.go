package notification

import "context"

type MailNotifier interface {
	Send(ctx context.Context, content []MailContent) error
}

type MailContent struct {
	To      string
	Subject string
	Body    string
}

// メールの一斉送信数
const emailBatchSize = 1000
