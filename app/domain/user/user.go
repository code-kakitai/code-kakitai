package user

import (
	"unicode/utf8"

	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/strings"
	"github.com/code-kakitai/go-pkg/ulid"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	id          string
	lastName    string
	firstName   string
	email       string
	phoneNumber string
	address     address
}

func NewUser(
	lastName string,
	firstName string,
	email string,
	phoneNumber string,
	prefecture string,
	city string,
	addressExtra string,
) (*User, error) {
	// 名前のバリデーション
	if err := validation.Validate([]string{lastName, firstName},
		validation.Each(validation.Required, validation.Length(nameLengthMin, nameLengthMax))); err != nil {
		return nil, errors.NewError("名前の値が不正です。")
	}

	// メールアドレスのバリデーション
	if err := validation.Validate(email,
		validation.Required, is.Email,
	); err != nil {
		return nil, errors.NewError("メールアドレスの値が不正です。")
	}

	// phoneNumberからハイフンを除く
	phoneNumber = strings.RemoveHyphen(phoneNumber)
	// 電話番号のバリデーション
	if _, ok := phoneNumberDigitMap[utf8.RuneCountInString(phoneNumber)]; !ok {
		return nil, errors.NewError("電話番号の値が不正です。")
	}

	ad, err := NewAddress(prefecture, city, addressExtra)
	if err != nil {
		return nil, err
	}
	return &User{
		id:          ulid.NewULID(),
		email:       email,
		phoneNumber: phoneNumber,
		lastName:    lastName,
		firstName:   firstName,
		address:     ad,
	}, nil
}

func Reconstruct(
	id string,
	email string,
	phoneNumber string,
	lastName string,
	firstName string,
	address address,
) *User {
	return &User{
		id:          id,
		email:       email,
		phoneNumber: phoneNumber,
		lastName:    lastName,
		firstName:   firstName,
		address:     address,
	}
}

func (u *User) Email() string {
	return u.email
}

func (u *User) PhoneNumber() string {
	return u.phoneNumber
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) Pref() string {
	return u.address.prefecture
}

func (u *User) City() string {
	return u.address.city
}

func (u *User) AddressExtra() string {
	return u.address.extra
}

var phoneNumberDigitMap = map[int]struct{}{
	PhoneNumberDigitTen:    {},
	PhoneNumberDigitEleven: {},
}

const (
	PhoneNumberDigitTen    = 10
	PhoneNumberDigitEleven = 11
)

const (
	nameLengthMax = 255
	nameLengthMin = 1
)

type address struct {
	prefecture string
	city       string
	extra      string
}

func NewAddress(
	prefecture string,
	city string,
	extra string,
) (address, error) {
	// 1つでも空のパラメーターがあればエラー
	if prefecture == "" || city == "" || extra == "" {
		return address{}, errors.NewError("住所の値が不正です。")
	}
	return address{
		prefecture: prefecture,
		city:       city,
		extra:      extra,
	}, nil
}

type UserRepository interface {
	Save(user User) error
	FindById(id string) (User, error)
}
