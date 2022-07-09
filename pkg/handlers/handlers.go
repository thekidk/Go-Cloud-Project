package handlers

import(
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/thekidk/Go-Cloud-Project/pkg/user"
)

var ErrorMethodNotAllowed = "Error: Method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error, omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {

	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return ApiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return ApiResponse(http.StatusOK, result)
	}

	result, err := user.FetchUsers(tableName, dynaClient)
	if err != nil {
		return ApiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	return ApiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {

}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {

}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {

}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {

	return ApiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)

}
