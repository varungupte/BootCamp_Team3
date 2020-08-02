package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	// create the api params
	params := &dynamodb.DeleteTableInput{
		TableName: aws.String("T3_Restaurant"),
	}

	// delete the table
	resp, err := db.DeleteTable(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return
	}
	// print the response data
	fmt.Printf("Response = %+v\n", resp)
}
