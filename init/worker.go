package init

// import (
// 	"context"
// 	"log"

// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/sqs"
// 	"go.mongodb.org/mongo-driver/mongo"

// 	"oms-service/workers"
// )

// type WorkerManager struct {
// 	sqsWorker *workers.SQSWorker
// }

// func NewWorkerManager(ctx context.Context, mongoClient *mongo.Client) (*WorkerManager, error) {
// 	// Load AWS configuration
// 	cfg, err := config.LoadDefaultConfig(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create SQS client
// 	sqsClient := sqs.NewFromConfig(cfg)

// 	// Initialize BulkOrderWorker with dependencies
// 	bulkWorker := &workers.BulkOrderWorker{
// 		// Add required dependencies
// 		// MongoDB client, etc.
// 	}

// 	// Create SQS worker
// 	sqsWorker := workers.NewSQSWorker(
// 		sqsClient,
// 		"YOUR_SQS_QUEUE_URL", // Get this from config
// 		bulkWorker,
// 	)

// 	return &WorkerManager{
// 		sqsWorker: sqsWorker,
// 	}, nil
// }

// func (wm *WorkerManager) StartWorkers(ctx context.Context) {
// 	wm.sqsWorker.Start(ctx)
// 	log.Println("All workers started successfully")
// }

// func (wm *WorkerManager) StopWorkers() {
// 	wm.sqsWorker.Stop()
// 	log.Println("All workers stopped successfully")
// }
