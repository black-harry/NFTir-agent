/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: collections.service.go provides function(s) to help client process data
*/

// @package
package controllers

// @import
import (
	"NFTir/agent/utils"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// @dev Fetches data in 6 hours
//
// @param logglyClient *loggly.ClientType := jearly/loggly
//
// @param tableName string: the name of the table
//
// @param db *dynamodb.DynamoDB: dynamodb connection
func PeriodicallyFetchData(tableName string, db *dynamodb.DynamoDB) {
	for { // infinite loop. Gap time = 6 hours
		utils.SetUpTableAsync(tableName, db);
		utils.FetchDataAsync(tableName, db)
	}
}
