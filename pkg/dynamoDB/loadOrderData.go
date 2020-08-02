package main

import (
	"encoding/json"
	"fmt"
	"github.com/varungupte/BootCamp_Team3/pkg/dynamoDB/types"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)


func main() {
	// read the json data file
	f, err := ioutil.ReadFile(string(os.Getenv("GOPATH")) + "/src/github.com/varungupte/BootCamp_Team3/assets/orderData.json")
	if err != nil {
		panic("Could not read movie json data file")
	}

	// parse the json movie data
	var orders []types.Order
	if err := json.Unmarshal(f, &orders); err != nil {
		panic("Could not parse json movie data")
	}

	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://127.0.0.1:8000"),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	// iterate through each movie
	for _, order := range orders {
		// marshal the movie struct into an aws attribute value map
		orderMap, err := dynamodbattribute.MarshalMap(order)
		if err != nil {
			panic("Cannot marshal movie into AttributeValue map")
		}

		// create the api params
		params := &dynamodb.PutItemInput{
			TableName: aws.String(orders_table),
			Item:      orderMap,
		}

		// put the item
		resp, err := db.PutItem(params)
		if err != nil {
			fmt.Printf("Unable to add order: %v\n", err.Error())
		} else {
			// print the response data
			fmt.Printf("Put item successful: '%s' (resp = '%+v')\n", order.Id, resp)
		}
	}
}