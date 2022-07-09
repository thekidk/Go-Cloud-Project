package aws_config

import(
	"os"
	"github.com/aws/aws-sdk-go"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-lambda-go/lambda"
)

func AwsConfig() error {

	region := os.getEnv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)
		},)

	if err != nil {
		return err
	}

	dynamoClient = dynamodb.New(awsSession)
	lambda.Start(handler)

	return nil

}
