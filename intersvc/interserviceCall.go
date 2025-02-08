package intersvc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	oms_kafka "oms-service/kafka"
	"oms-service/models"
	"time"

	"github.com/omniful/go_commons/http"

	interservice_client "github.com/omniful/go_commons/interservice-client"
)

func ValidateOrders(order *models.Order) {
	fmt.Println("Validate fxn called")

	for _, orderItem := range order.OrderItems {
		config := interservice_client.Config{
			ServiceName: "oms-service",
			BaseURL:     "http://localhost:8081/api/v1/orders/",
			Timeout:     5 * time.Second,
		}

		client, err := interservice_client.NewClientWithConfig(config)
		if err != nil {
			return
		}

		url := config.BaseURL + "validate_order"
		body := map[string]string{
			"sku_id": orderItem.SKUID,
			"hub_id": orderItem.HubID,
		}

		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return
		}

		req := &http.Request{
			Url:     url,
			Body:    bytes.NewReader(bodyBytes),
			Timeout: 7 * time.Second,
			Headers: map[string][]string{
				"Content-Type": {"application/json"},
			},
		}

		resp, intersvcErr := client.Post(req, "/")
		if intersvcErr.StatusCode.Is4xx() {
			fmt.Printf("Error making GET request to validate SKU: %v\n", err)
			return
		} else {
			fmt.Print(resp)
			log.Printf("Order with Order ID: %v having product %v from hub %v is VALID \n", order.ID, orderItem.SKUID, orderItem.HubID)

			// Publish This Order Item in a message to Kafka
			bytesOrderItem, _ := json.Marshal(orderItem)
			oms_kafka.PublishMessageToKafka(bytesOrderItem, order.ID)

			return
		}
	}
}
