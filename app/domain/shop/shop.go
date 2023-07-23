package shop

import (
	"unicode/utf8"

	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/ulid"
)

type Shop struct {
	id          string
	ownerID     string
	name        string
	description string
}

func ReconstructShop(
	id string,
	ownerID string,
	name string,
	description string,
) (*Shop, error) {
	return newShop(
		id,
		ownerID,
		name,
		description,
	)
}

func NewShop(
	ownerID string,
	name string,
	description string,
) (*Shop, error) {
	return newShop(
		"",
		ownerID,
		name,
		description,
	)
}

func newShop(
	id string,
	ownerID string,
	name string,
	description string,
) (*Shop, error) {
	// idが空文字の時は新規作成
	if id == "" {
		id = ulid.NewULID()
	}
	// ownerIDのバリデーション
	if !ulid.IsValid(ownerID) {
		return nil, errors.NewError("オーナーIDの値が不正です。")
	}
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errors.NewError("名前の値が不正です。")
	}
	// 説明のバリデーション
	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errors.NewError("説明の値が不正です。")
	}

	return &Shop{
		id:          id,
		ownerID:     ownerID,
		name:        name,
		description: description,
	}, nil
}

const (
	// 名前の最大値/最小値
	nameLengthMin = 1
	nameLengthMax = 100

	// 説明の最大値/最小値
	descriptionLengthMin = 1
	descriptionLengthMax = 1000
)
