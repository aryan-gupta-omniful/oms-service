package models

import (
	"time"
)

type Customer struct {
	ID                     string      `bson:"_id,omitempty" json:"id"`
	FirstName              string      `bson:"first_name" json:"first_name"`
	LastName               string      `bson:"last_name" json:"last_name"`
	EmailID                string      `bson:"email_id" json:"email_id"`
	CountryCode            string      `bson:"country_code" json:"country_code"`
	PhoneNumber            string      `bson:"phone_number" json:"phone_number"`
	Gender                 string      `bson:"gender" json:"gender"`
	DOB                    time.Time   `bson:"dob" json:"dob"`
	Status                 string      `bson:"status" json:"status"`
	DefaultShippingAddress interface{} `bson:"default_shipping_address" json:"default_shipping_address"`
	CreatedAt              time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt              time.Time   `bson:"updated_at" json:"updated_at"`
}

type Address struct {
	ID           string `bson:"_id,omitempty" json:"id"`
	EntityID     string `bson:"entity_id" json:"entity_id"`
	EntityType   string `bson:"entity_type" json:"entity_type"`
	AddressLine1 string `bson:"address_line1" json:"address_line1"`
	AddressLine2 string `bson:"address_line2" json:"address_line2"`
	Pincode      string `bson:"pincode" json:"pincode"`
	City         string `bson:"city" json:"city"`
	State        string `bson:"state" json:"state"`
	Country      string `bson:"country" json:"country"`
}

type Credentials struct {
	EntityID     string    `bson:"entity_id" json:"entity_id"`
	EntityType   string    `bson:"entity_type" json:"entity_type"`
	RefreshToken string    `bson:"refresh_token" json:"refresh_token"`
	ExpiryAt     time.Time `bson:"expiry_at" json:"expiry_at"`
	PasswordHash string    `bson:"password_hash" json:"password_hash"`
}
