package errors

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

/*
Implements error handling for the BGP connector
*/

// For now, this kills the program but in the future this should signal the broker
// that this particular goroutine is screwed
func RaiseError(e interface{}) {
	//log.Print(e)
	//log.Fatal(e)
	panic(e)
}

// Check to see if the passed err value is non-nil and handle accordingly
func CheckError(err error) {
	if err != nil {
		RaiseError(err.Error())
	}
}

// Same as CheckError but checks for AWS error codes specifically
func CheckAwsError(err error) {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		RaiseError(err)
		return
	}
}
