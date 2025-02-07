package models

import "time"

type Order struct {
	ID              string      `bson:"_id,omitempty" json:"order_id"`
	CustomerID      string      `bson:"customer_id" json:"customer_id"`
	CreatedAt       time.Time   `bson:"created_at" json:"created_at"`
	Currency        string      `bson:"currency" json:"currency"`
	TotalAmount     int         `bson:"total_amount" json:"total_amount"`
	TransactionID   string      `bson:"transaction_id" json:"transaction_id"`
	OrderItems      []OrderItem `bson:"order_items" json:"order_items"`
	ModeOfPayment   string      `bson:"mode_of_payment" json:"mode_of_payment"`
	Status          string      `bson:"status" json:"status"`
	BillingAddress  interface{} `bson:"billing_address" json:"billing_address"`
	ShippingAddress interface{} `bson:"shipping_address" json:"shipping_address"`
	InvoiceID       interface{} `bson:"invoice_id" json:"invoice_id"`
	TenantID        interface{} `bson:"tenant_id" json:"tenant_id"`
}

type OrderItem struct {
	ID              string `bson:"_id,omitempty" json:"order_items_id"`
	OrderID         string
	SKUID           string `bson:"sku_id" json:"sku_id"`
	QuantityOrdered int    `bson:"quantity_ordered" json:"quantity_ordered"`
	HubID           string `bson:"hub_id" json:"hub_id"`
	SellerID        string `bson:"seller_id" json:"seller_id"`
}

type BulkOrderQueueMessage struct {
	CustomerID string `json:"customer_id"`
	FilePath   string `json:"file_path"`
}

type BulkOrderRequest struct {
	SellerID int    `json:"sellerID"`
	FilePath string `json:"filePath"`
}

type SQSMessage struct {
	SellerID int    `json:"sellerID"`
	FilePath string `json:"filePath"`
}
