package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region: aws.String("us-east-1"),
	}))

	db := dynamodb.New(sess)

	awsParams := &dynamodb.CreateTableInput{
		TableName: aws.String("OrdersT3"),

		KeySchema: [] *dynamodb.KeySchemaElement {
			{AttributeName: aws.String("CustId"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("Id"), KeyType: aws.String("RANGE")},
		},

		AttributeDefinitions: [] *dynamodb.AttributeDefinition {
			{AttributeName: aws.String("CustId"), AttributeType: aws.String("N")},
			{AttributeName: aws.String("Id"), AttributeType: aws.String("N")},
		},

		ProvisionedThroughput: &dynamodb.ProvisionedThroughput {
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}
	response, err := db.CreateTable(awsParams)
	if err != nil {
		log.Fatalf("Sorry error creating table : %v", err)
	}
	// print the response
	fmt.Println(response)
}