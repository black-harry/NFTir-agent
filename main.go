/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
*/

// @package
package main

// @import
import (
	"NFTir/agent/controllers"
	"NFTir/agent/utils"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// @dev initialize global variables
var (
	db *dynamodb.DynamoDB
)

// @dev Run before main()
func init()  {
	if (os.Getenv("APP_MODE") != "release") {
		utils.LoadEnvVars()
	}
	db = utils.EstablishAwsDynamodbSession()
}

// @dev Root function 
func main()  {
	controllers.PeriodicallyFetchData(os.Getenv("TABLE_NAME"), db)
}