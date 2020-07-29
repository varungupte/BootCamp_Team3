package dbClient

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"github.com/varungupte/BootCamp_Team3/pkg/services/orders/orders_server"
)

type DynamoEntity struct {
	RestaurantName string `dynamodbav:"restaurent_name"`
	DishName       string `dynamodbav:"dish_name"`
	Info orders_server.Order `dynamodbav:"orders_info"`
}
func UploadData([]orders_server.Order) {
}

func CreateTable() {
	//create an Aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}))
	//create a dynamodb instance
	db := dynamodb.New(sess)

	//create the api params
	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Orders"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("restaurent_name"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("dish_name"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("dish_name"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("restaurent_name"), AttributeType: aws.String("S")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}
	//create the table
	resp, err := db.CreateTable(params)
	errorutil.CheckError(err, "Unable to create Table")

	fmt.Println("The response from create Table", resp)
}

func UploadOrder(ord orders_server.Order)  {
	//create an Aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}))
	//create a dynamodb instance
	db := dynamodb.New(sess)
	dynamoEntity:=DynamoEntity{
		RestaurantName: ord.Restau.Name,
		DishName: ord.DishName,
		Info: ord,
	}
	fmt.Println("Inserting this do db",dynamoEntity)
	dbMap,err:=dynamodbattribute.MarshalMap(dynamoEntity)
	errorutil.CheckError(err,"Something went wrong while marhsaling db entity")
	params:=&dynamodb.PutItemInput{
		TableName: aws.String("Orders"),
		Item: dbMap,
	}
	resp,err:=db.PutItem(params)
	errorutil.CheckError(err,"Got this while inserting value to db")
	fmt.Println("Success")
	fmt.Println(resp)
}


