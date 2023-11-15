package repositories

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSRepository struct {
	Client   *sqs.SQS
	QueueURL string
}

func NewSQSRepository(region, queueURL string) (*SQSRepository, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	sqsClient := sqs.New(sess)

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

	_, err = s.Client.SendMessage(&sqs.SendMessageInput{
		MessageBody:  aws.String(string(messageBody)),
		QueueUrl:     aws.String(s.QueueURL),
		DelaySeconds: aws.Int64(0),
	})

	return err
}