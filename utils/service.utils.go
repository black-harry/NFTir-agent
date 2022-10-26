/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: error.utils.go provides function that helps handle exceptions
*/

// @package
package utils

// @import
import (
	"NFTir/agent/models"
	"encoding/json"
	"log"

	"github.com/jamespearly/loggly"
)

// @dev Loads environment varables
//
// @param e error
func HandleException(e error) {
	if (e != nil) {
		log.Fatal(e);
	}
}


// @dev Handles interacting with Loggly
// 
// @param httpLogglyMessage models.HttpLogglyMessage
// 
// @param level string
func HandleLoggly(logglyMessage models.LogglyMessage, level string) {
	logglyClient := loggly.New("NFTir")

	// stringify struct to prepare for jearly/loggly.Send()
	stringifiedLogglyMessage, marshalErr := json.Marshal(logglyMessage)
	HandleException(marshalErr)

	// Send message to Loggly
	logglyErr := logglyClient.Send(level, string(stringifiedLogglyMessage)); 
	HandleException(logglyErr)
}