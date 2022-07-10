package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/thekidk/Go-Cloud-Project/config"
	"github.com/thekidk/Go-Cloud-Project/pkg/handlers"
)

var(
	dynaClient dynamodbiface.DynamoDBAPI
)

const tableName = "LambdaInGoUser"


func main() {

	err := config.AwsConfig(dynaClient)
	if err != nil {
		return
	}
	
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser(req, tableName, dynaClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynaClient)
	case "PUT":
		return handlers.UpdateUser(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynaClient)
	default:
		return handlers.UnhandledMethod()
	}

}