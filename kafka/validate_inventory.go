package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	http2 "net/http"
	"oms-service/models"
)

func ValidateInventory(ctx context.Context, order models.KafkaResponseOrderMessage) error {

	log.Printf("Validating inventory for order ID: %s \n", order.OrderID)

	client := &http2.Client{}
	url := "http://localhost:8081/api/v1/orders/validate_inventory"

	reqBody, err := json.Marshal(order)
	if err != nil {
		return err
	}

	req, err := http2.NewRequest(http2.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http2.StatusOK {
		return fmt.Errorf("inventory validation failed with status code: %d", resp.StatusCode)
	} else {
		fmt.Println("Inventory validation successful !!!!!!!!!!!")
	}

	return nil
}
