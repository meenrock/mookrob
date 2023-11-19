package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSRepository struct {
	Client   *sqs.Client
	QueueURL string
}

func NewSQSRepository(region, queueURL string) (*SQSRepository, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	sqsClient := sqs.NewFromConfig(cfg)

	return &SQSRepository{
		Client:   sqsClient,
		QueueURL: queueURL,
	}, nil
}

func (s *SQSRepository) SendMessage(data interface{}) error {
	messageBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = s.Client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		MessageBody:  aws.String(string(messageBody)),
		QueueUrl:     aws.String(s.QueueURL),
		DelaySeconds: 0,
	})

	return err
}

func ReceiveMessage(ctx context.Context, queueURL string) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load AWS SDK configuration: %v", err)
	}

	client := sqs.NewFromConfig(cfg)

	for {
		result, err := client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:              aws.String(queueURL),
			MaxNumberOfMessages:   1,
			VisibilityTimeout:     10,
			WaitTimeSeconds:       5,
			MessageAttributeNames: []string{"All"},
		})

		if err != nil {
			return fmt.Errorf("failed to receive message from SQS: %v", err)
		}

		for _, message := range result.Messages {
			fmt.Printf("Received message: %v\n", *message.Body)

			_, err := client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: message.ReceiptHandle,
			})

			if err != nil {
				return fmt.Errorf("failed to delete message from SQS: %v", err)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
