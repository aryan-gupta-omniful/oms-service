package models

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
