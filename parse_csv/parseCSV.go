package parse_csv

import (
	"context"
	"fmt"
	"log"
	"oms-service/intersvc"
	"oms-service/models"
	"os"
	"strconv"
	"time"

	"github.com/omniful/go_commons/csv"
)

func ParseCSV(filePath string) ([]*models.Order, error) {
	fmt.Println("Parse CSV function called successfull!")
	fmt.Println("This is the file beign opened: ", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error in opening file path.")
	}
	defer file.Close()

	// Map to group items by order_no
	orderGroups := make(map[string]*models.Order)

	// Initialize the CSV reader (based on your previous implementation)
	CSV, err := csv.NewCommonCSV(
		csv.WithBatchSize(100),
		csv.WithSource(csv.Local),
		csv.WithLocalFileInfo(filePath),
		csv.WithHeaderSanitizers(csv.SanitizeAsterisks, csv.SanitizeToLower),
		csv.WithDataRowSanitizers(csv.SanitizeSpace, csv.SanitizeToLower),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize CSV reader: %v", err)
	}

	err = CSV.InitializeReader(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize CSV reader: %v", err)
	}

	// Process the records and group them by order_no and customer_name
	for !CSV.IsEOF() {
		var records csv.Records
		records, err := CSV.ReadNextBatch()
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Processing records:")
		fmt.Println(records)
		for _, record := range records {
			orderID := record[0]  // order_id
			skuID := record[1]    // sku_id
			quantity := record[2] // quantity
			sellerID := record[3] // seller_id
			hubID := record[4]    // hub_id

			// Convert quantity to integer
			IntQuantity, err := strconv.Atoi(quantity)
			if err != nil {
				return nil, fmt.Errorf("invalid quantity %s: %v", quantity, err)
			}

			// Check if the  group forderor this order_id already exists
			orderKey := orderID
			order, exists := orderGroups[orderKey]
			if !exists {
				// If order doesn't exist, create a new order
				now := time.Now()
				order = &models.Order{
					// SellerID:     sellerID,
					// HubID:        hubID,
					ID:              orderID,
					CustomerID:      sellerID,
					CreatedAt:       now,
					Currency:        "INR",
					TotalAmount:     0,
					TransactionID:   "SAMPLE_TXN_ID",
					ModeOfPayment:   "PAYPAL",
					Status:          "on_hold",
					BillingAddress:  "sample address",
					ShippingAddress: "sample address",
					InvoiceID:       "999",
					TenantID:        "111",
					OrderItems:      []models.OrderItem{}, // Start with an empty slice of items
				}
				// Add the new order to the map
				orderGroups[orderKey] = order
			}

			// Create a new OrderItem and append it to the order's OrderItems
			orderItem := models.OrderItem{
				OrderID:         orderID,
				SKUID:           skuID,
				QuantityOrdered: IntQuantity,
				HubID:           hubID,
				SellerID:        sellerID,
			}
			order.OrderItems = append(order.OrderItems, orderItem)
		}
	}

	var orders []*models.Order
	for _, order := range orderGroups {
		orders = append(orders, order)
	}

	fmt.Println("Final orders:")
	for _, order := range orders {
		fmt.Printf("Order No: %s, Total Items: %d\n", order.ID, len(order.OrderItems))
		go intersvc.ValidateOrders(order)
	}

	return orders, nil

}
