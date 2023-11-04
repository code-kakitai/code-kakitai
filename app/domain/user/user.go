package user

import (
	"net/mail"
	"unicode/utf8"

	"github.com/code-kakitai/go-pkg/strings"
	"github.com/code-kakitai/go-pkg/ulid"

	errDomain "github/code-kakitai/code-kakitai/domain/error"
)

type User struct {
	id          string
	email       string
	phoneNumber string
	lastName    string
	firstName   string
	address     address
}

// 永続化層から取得したデータをドメインに変換
func Reconstruct(
	id string,
	email string,
	phoneNumber string,
	lastName string,
	firstName string,
	prefecture string,
	city string,
	addressExtra string,
) (*User, error) {
	return newUser(
		id,
		email,
		phoneNumber,
		lastName,
		firstName,
		prefecture,
		city,
		addressExtra,
	)
}

func NewUser(
	email string,
	phoneNumber string,
	lastName string,
	firstName string,
	prefecture string,
	city string,
	addressExtra string,
) (*User, error) {
	return newUser(
		ulid.NewULID(),
		email,
		phoneNumber,
		lastName,
		firstName,
		prefecture,
		city,
		addressExtra,
	)
}

func newUser(
	id string,
	email string,
	phoneNumber string,
	lastName string,
	firstName string,
	prefecture string,
	city string,
	addressExtra string,
) (*User, error) {
	// 名前のバリデーション
	if utf8.RuneCountInString(lastName) < nameLengthMin || utf8.RuneCountInString(lastName) > nameLengthMax {
		return nil, errDomain.NewError("名前（姓）の値が不正です。")
	}
	if utf8.RuneCountInString(firstName) < nameLengthMin || utf8.RuneCountInString(firstName) > nameLengthMax {
		return nil, errDomain.NewError("名前（名）の値が不正です。")
	}

	// メールアドレスのバリデーション
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errDomain.NewError("メールアドレスの値が不正です。")
	}

	// phoneNumberからハイフンを除く
	phoneNumber = strings.RemoveHyphen(phoneNumber)
	// 電話番号のバリデーション
	if _, ok := phoneNumberDigitMap[utf8.RuneCountInString(phoneNumber)]; !ok {
		return nil, errDomain.NewError("電話番号の値が不正です。")
	}

	ad, err := newAddress(prefecture, city, addressExtra)
	if err != nil {
		return nil, err
	}

	return &User{
		id:          id,
		email:       email,
		phoneNumber: phoneNumber,
		lastName:    lastName,
		firstName:   firstName,
		address:     ad,
	}, nil
}

func (u *User) ID() string {
	return u.id
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
		return address{}, errDomain.NewError("住所の値が不正です。")
	}
	return address{
		prefecture: prefecture,
		city:       city,
		extra:      extra,
	}, nil
}
