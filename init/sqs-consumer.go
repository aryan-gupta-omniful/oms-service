package init

import (
	"context"
	"encoding/json"
	"fmt"
	"oms-service/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/omniful/go_commons/log"
)

func getConfig() *aws.Config {
	if awsConfig == nil {
		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("eu-north-1"))
		if err != nil {
			panic("Unable to connect to AWS")
		}
		awsConfig = &cfg
		return awsConfig
	}
	return awsConfig
}

func initialiseSQSConsumer(ctx context.Context) {

	sqsClient := sqs.NewFromConfig(*getConfig())

	sqURL := getSQSUrl(ctx)
	fmt.Println("Queue URL: ", *sqURL)

	// This will constantly listen to the SQS queue and print the messages
	for {
		messagesResult, err := sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl: sqURL,
		})
		if err != nil {
			fmt.Println("Unable to receive mesaages from SQS: ", err)
		}

		for _, message := range messagesResult.Messages {
			fmt.Println("Message: ", *message.Body)
			var sqsMessage models.SQSMessage
			if err := json.Unmarshal([]byte(*message.Body), &sqsMessage); err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}
		}
	}
}
