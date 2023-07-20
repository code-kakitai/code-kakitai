package owner

import (
	"net/mail"
	"unicode/utf8"

	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/ulid"
)

type Owner struct {
	id    string
	name  string
	email string
}

func NewOwner(name string, email string) (*Owner, error) {
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errors.NewError("名前の値が不正です。")
	}

	// メールアドレスのバリデーション
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errors.NewError("メールアドレスの値が不正です。")
	}

	return &Owner{
		id:    ulid.NewULID(),
		name:  name,
		email: email,
	}, nil
}

func (o *Owner) Name() string {
	return o.name
}

func (o *Owner) Email() string {
	return o.email
}

const (
	nameLengthMax = 255
	nameLengthMin = 1
)
