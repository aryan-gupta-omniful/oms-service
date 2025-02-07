package intersvc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	oms_kafka "oms-service/kafka"
	"oms-service/models"
	"strings"
)

type responsePost struct {
	Message string `json:"message"`
}

func ValidateOrders(order *models.Order) {
	fmt.Println("Validate fxn called")
	// size := len(order.OrderItems)
	for _, orderItem := range order.OrderItems {
		requestBody := fmt.Sprintf(`{
            "sku_id": %v,
            "hub_id": %v
        }`, orderItem.SKUID, orderItem.HubID)

		requestBodyReader := strings.NewReader(requestBody)

		res, _ := http.Post("http://localhost:8081/api/v1/orders/validate_order", "application/json", requestBodyReader)
		content, _ := io.ReadAll(res.Body)

		var responsePost responsePost
		err := json.Unmarshal(content, &responsePost)
		if err != nil {
			fmt.Println("Error unmarshalling response from Post Request.")
		}
		fmt.Println("response of post request: ", responsePost.Message)
		if responsePost.Message == "Validation successful" {
			log.Printf("Order with Order ID: %v having product %v from hub %v is VALID \n", order.ID, orderItem.SKUID, orderItem.HubID)

			// Publish This Order Item in a message to Kafka
			bytesOrderItem, _ := json.Marshal(orderItem)
			oms_kafka.PublishMessageToKafka(bytesOrderItem, order.ID)

		} else {
			log.Printf("Order with Order ID: %v having product %v from hub %v is invalid \n", order.ID, orderItem.SKUID, orderItem.HubID)
		}
	}
}
