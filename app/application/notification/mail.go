package notification

import (
	ownerDomain "github/code-kakitai/code-kakitai/domain/owner"
	userDomain "github/code-kakitai/code-kakitai/domain/user"
)

type MailContent struct {
	Title   string
	Message string
}

func (c *MailContent) GetContent() string {
	return c.Message
}

type MailMetaData struct {
	From string
	To   []string
}

func (m *MailMetaData) GetNotificationTo() []string {
	return m.To
}

func NewMetaDataFromUser(us []*userDomain.User) MetaData {
	var emails []string
	for _, u := range us {
		emails = append(emails, u.Email())
	}
	return &MailMetaData{
		From: defaultFrom,
		To:   emails,
	}
}

const defaultFrom = "noreply@code-kakitai.com"

func NewMetaDataFromOwner(o []*ownerDomain.Owner) MetaData {
	var emails []string
	for _, u := range o {
		emails = append(emails, u.Email())
	}
	return &MailMetaData{
		From: defaultFrom,
		To:   emails,
	}
}
