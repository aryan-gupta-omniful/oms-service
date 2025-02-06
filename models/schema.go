package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName              string             `bson:"first_name" json:"first_name"`
	LastName               string             `bson:"last_name" json:"last_name"`
	EmailID                string             `bson:"email_id" json:"email_id"`
	CountryCode            string             `bson:"country_code" json:"country_code"`
	PhoneNumber            string             `bson:"phone_number" json:"phone_number"`
	Gender                 string             `bson:"gender" json:"gender"`
	DOB                    time.Time          `bson:"dob" json:"dob"`
	Status                 string             `bson:"status" json:"status"`
	DefaultShippingAddress interface{}        `bson:"default_shipping_address" json:"default_shipping_address"`
	CreatedAt              time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt              time.Time          `bson:"updated_at" json:"updated_at"`
}

type Address struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntityID     primitive.ObjectID `bson:"entity_id" json:"entity_id"`
	EntityType   string             `bson:"entity_type" json:"entity_type"`
	AddressLine1 string             `bson:"address_line1" json:"address_line1"`
	AddressLine2 string             `bson:"address_line2" json:"address_line2"`
	Pincode      string             `bson:"pincode" json:"pincode"`
	City         string             `bson:"city" json:"city"`
	State        string             `bson:"state" json:"state"`
	Country      string             `bson:"country" json:"country"`
}

type OrderItem struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"order_items_id"`
	SKUID           primitive.ObjectID `bson:"sku_id" json:"sku_id"`
	QuantityOrdered int                `bson:"quantity_ordered" json:"quantity_ordered"`
	HubID           primitive.ObjectID `bson:"hub_id" json:"hub_id"`
	SellerID        primitive.ObjectID `bson:"seller_id" json:"seller_id"`
}

type Order struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"order_id"`
	CustomerID      primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	Currency        string             `bson:"currency" json:"currency"`
	TotalAmount     int                `bson:"total_amount" json:"total_amount"`
	TransactionID   string             `bson:"transaction_id" json:"transaction_id"`
	OrderItems      []OrderItem        `bson:"order_items" json:"order_items"`
	ModeOfPayment   string             `bson:"mode_of_payment" json:"mode_of_payment"`
	Status          string             `bson:"status" json:"status"`
	BillingAddress  interface{}        `bson:"billing_address" json:"billing_address"`
	ShippingAddress interface{}        `bson:"shipping_address" json:"shipping_address"`
	InvoiceID       interface{}        `bson:"invoice_id" json:"invoice_id"`
	TenantID        primitive.ObjectID `bson:"tenant_id" json:"tenant_id"`
}

type Credentials struct {
	EntityID     primitive.ObjectID `bson:"entity_id" json:"entity_id"`
	EntityType   string             `bson:"entity_type" json:"entity_type"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
	ExpiryAt     time.Time          `bson:"expiry_at" json:"expiry_at"`
	PasswordHash string             `bson:"password_hash" json:"password_hash"`
}
