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

type responsePost struct {
	Message string `json:"message"`
}

func ValidateOrders(order *models.Order) {
	fmt.Println("Validate fxn called")
	// size := len(order.OrderItems)
	for _, orderItem := range order.OrderItems {
		// requestBody := fmt.Sprintf(`{
		//     "sku_id": %v,
		//     "hub_id": %v
		// }`, orderItem.SKUID, orderItem.HubID)

		// requestBodyReader := strings.NewReader(requestBody)

		// res, _ := http.Post("http://localhost:8081/api/v1/orders/validate_order", "application/json", requestBodyReader)
		// content, _ := io.ReadAll(res.Body)

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
			Url:     url, // Use configured URL
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

		// var responsePost responsePost
		// err := json.Unmarshal(content, &responsePost)
		// if err != nil {
		// 	fmt.Println("Error unmarshalling response from Post Request.")
		// }
		// fmt.Println("response of post request: ", responsePost.Message)
		// if responsePost.Message == "Validation successful" {
		// 	log.Printf("Order with Order ID: %v having product %v from hub %v is VALID \n", order.ID, orderItem.SKUID, orderItem.HubID)

		// 	// Publish This Order Item in a message to Kafka
		// 	bytesOrderItem, _ := json.Marshal(orderItem)
		// 	oms_kafka.PublishMessageToKafka(bytesOrderItem, order.ID)

		// } else {
		// 	log.Printf("Order with Order ID: %v having product %v from hub %v is invalid \n", order.ID, orderItem.SKUID, orderItem.HubID)
		// }
	}
}
