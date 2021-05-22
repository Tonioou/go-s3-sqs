package processor

import (
	"encoding/json"
	"fmt"

	models "cmd/main.go/internal/models/events"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type Processor interface {
	Process(*sqs.Message)
}

func GetProcessor() *S3Processor {
	return &S3Processor{}
}

type S3Processor struct{}

func (sp *S3Processor) Process(rawMessage *sqs.Message) {
	var message models.Message
	json.Unmarshal([]byte(*rawMessage.Body), &message)

	fmt.Println(message.Records[0].S3.Object.Key)
}
