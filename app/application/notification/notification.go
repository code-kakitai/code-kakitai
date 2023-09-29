package notification

import "context"

type Notifier interface {
	SendMail(ctx context.Context, content []MailContent) error
}

type MailContent struct {
	To      string
	Subject string
	Body    string
}

// メールの一斉送信数
const emailBatchSize = 1000
