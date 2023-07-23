package user

import (
	"net/mail"
	"unicode/utf8"

	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/strings"
	"github.com/code-kakitai/go-pkg/ulid"
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
	if utf8.RuneCountInString(lastName) < nameLengthMin || utf8.RuneCountInString(lastName) > nameLengthMax {
		return nil, errors.NewError("名前（姓）の値が不正です。")
	}
	if utf8.RuneCountInString(firstName) < nameLengthMin || utf8.RuneCountInString(firstName) > nameLengthMax {
		return nil, errors.NewError("名前（名）の値が不正です。")
	}

	// メールアドレスのバリデーション
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errors.NewError("メールアドレスの値が不正です。")
	}

	// phoneNumberからハイフンを除く
	phoneNumber = strings.RemoveHyphen(phoneNumber)
	// 電話番号のバリデーション
	if _, ok := phoneNumberDigitMap[utf8.RuneCountInString(phoneNumber)]; !ok {
		return nil, errors.NewError("電話番号の値が不正です。")
	}

	ad, err := newAddress(prefecture, city, addressExtra)
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

func CreateUser()

// 永続化層から取得したデータをドメインに変換
func BuilderUser(
	id string,
	email string,
	phoneNumber string,
	lastName string,
	firstName string,
	prefecture string,
	city string,
	addressExtra string,
) (*User, error) {
	return NewUser(
		lastName,
		firstName,
		email,
		phoneNumber,
		prefecture,
		city,
		addressExtra,
	)
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

func newAddress(
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
