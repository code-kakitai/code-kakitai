package notification

import (
	"context"
)

type Notificationer interface {
	Dispatch(ctx context.Context, data MetaData, content Content) error
}

type MetaData interface {
	GetNotificationTo() []string
}

type Content interface {
	GetContent() string
}
