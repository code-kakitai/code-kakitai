// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package dbgen

import (
	"time"
)

type Owner struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID          string `json:"id"`
	OwnerID     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int32  `json:"stock"`
}

type User struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Firebaseuid string `json:"firebaseuid"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	// 郵便番号
	PostalCode string `json:"postal_code"`
	// 都道府県
	Prefecture string `json:"prefecture"`
	// 市区町村
	City string `json:"city"`
	// 住所詳細
	AddressExtra string    `json:"address_extra"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}