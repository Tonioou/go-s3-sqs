package consumer

import (
	"flag"
	"fmt"

	"cmd/main.go/internal/processor"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func PoolMessages() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	urlResult, err := getQueueUrl(svc)
	if err != nil {
		fmt.Printf("Cannot get queue url", err)
	}
	getMessages(urlResult, svc)
	//go func() {  }()

}

func getQueueUrl(svc *sqs.SQS) (*sqs.GetQueueUrlOutput, error) {
	queue := "S3_CREATED_OBJECT"
	urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queue,
	})
	return urlResult, err
}

func getMessages(queueUrl *sqs.GetQueueUrlOutput, svc *sqs.SQS) {
	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueUrl.QueueUrl,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   getTimeout(),
	})

	if err != nil {
		fmt.Printf("Cannot pool messages", err)
	}
	processor := processor.GetProcessor()

	for _, message := range msgResult.Messages {
		processor.Process(message)
	}
}

func getTimeout() *int64 {
	timeout := flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	flag.Parse()
	if *timeout < 0 {
		*timeout = 0
	}

	if *timeout > 12*60*60 {
		*timeout = 12 * 60 * 60
	}
	return timeout
}
