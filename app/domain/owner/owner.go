package owner

import (
	"context"

	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/ulid"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Owner struct {
	id    string
	name  string
	email string
}

func NewOwner(name string, email string) (*Owner, error) {
	// 名前のバリデーション
	if err := validation.Validate(name,
		validation.Required, validation.Length(nameLengthMin, nameLengthMax)); err != nil {
		return nil, errors.NewError("名前の値が不正です。")
	}

	// メールアドレスのバリデーション
	if err := validation.Validate(email,
		validation.Required, is.Email,
	); err != nil {
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

type OwnerRepository interface {
	Save(ctx context.Context) error
	FindById(ctx context.Context, id string) (*Owner, error)
}
