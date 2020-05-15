package modal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

// GetDynamoDBSession : To return the dynamoDB session based on environment
func GetDynamoDBSession() (db *dynamo.DB) {
	var awsSession = session.Must(session.NewSession())
	var config = aws.Config{}
	localDev := os.Getenv("LocalDevelopment")
	if len(localDev) != 0 {
		config.Endpoint = aws.String("http://localhost:8000")
		config.Region = aws.String("us-west-2")
	} else {
		config.Region = aws.String("us-east-2")
	}
	db = dynamo.New(awsSession, &config)
	return
}
