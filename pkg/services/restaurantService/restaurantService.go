package restaurantService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/varungupte/BootCamp_Team3/pkg/dynamoDB/types"
	"github.com/varungupte/BootCamp_Team3/pkg/errorutil"
	"strconv"
)


func SaveRestaurant(entity types.Restaurant) (types.Restaurant, error) {
	db := MakeNewDbSession()
	restaurantMap, err := dynamodbattribute.MarshalMap(entity)
	errorutil.CheckError(err, "Error occured While Marshalling Restaurant Entity")
	_, err = db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("Restaurant"),
		Item:      restaurantMap,
	})
	errorutil.CheckError(err, "Error Occured While putting Restaurent in Db")
	fmt.Println("Put item successful")
	return entity, nil
}

func GetRestaurant(id int64) (types.Restaurant, error) {
	db := MakeNewDbSession()
	resp, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Restaurant"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(int(id))),
			},
		},
	})
	if err != nil {
		return types.Restaurant{}, err
	}
	var restaurantEntity types.Restaurant
	err = dynamodbattribute.UnmarshalMap(resp.Item, &restaurantEntity)
	return restaurantEntity, nil
}

func DeleteRestaurant(id int64) error{
	db := MakeNewDbSession()
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("Restaurant"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(int(id))),
			},
		},
	}
	// delete the item
	resp, err := db.DeleteItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return err
	}
	// print the response data
	fmt.Println("Success")
	fmt.Println(resp)
	return nil

}

func GetRestaurantItems(id int64)([]types.Item,error)  {
	res,err:=GetRestaurant(id)
	if err!=nil {
		return []types.Item{},err
	}
	return res.Items, nil
}

func GetItemsBetweenRange(min float32,max float32,id int64)([]types.Item,error)  {
	res,err:=GetRestaurant(id)
	if err!=nil {
		return []types.Item{},err
	}
	reqItems:=make([]types.Item,0,5)
	for _,item:=range res.Items{
		if item.Cost>=min&&item.Cost<=max {
			reqItems=append(reqItems, item)
		}
	}
	return reqItems,nil
}

func DeleteItemFromRestaurant(restaurantId int64,itemName string) error {
	restaurantEntity,err:=GetRestaurant(restaurantId)
	if err!=nil{
	   return err
	}
	var itemIndex int
	for i,item:=range restaurantEntity.Items{
		if item.Name==itemName{
			itemIndex=i
			break
		}
	}
	restaurantEntity.Items= append(restaurantEntity.Items[:itemIndex], restaurantEntity.Items[itemIndex+1:]...)
	_,err=SaveRestaurant(restaurantEntity)
	if err!=nil{
		return err
	}
	return nil
}
func UpdateItemInRestaurant(restaurantId int64,item types.Item) error {
	restaurantEntity,err:=GetRestaurant(restaurantId)
	if err!=nil{
		return err
	}
	itemPresent:=false
	for i,val:=range restaurantEntity.Items{
		if val.Name==item.Name{
			restaurantEntity.Items[i]=item
			itemPresent=true
			break
		}
	}
	if !itemPresent{
		restaurantEntity.Items=append(restaurantEntity.Items,item)
	}
	_,err=SaveRestaurant(restaurantEntity)
	if err!=nil{
		return err
	}
	return nil
}
func GetRestaurantCount()(*int64,error)  {
	db:=MakeNewDbSession()
	// create the api params
	params := &dynamodb.DescribeTableInput{
		TableName: aws.String("Restaurant"),
	}
	// get the table description
	resp, err := db.DescribeTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
   return resp.Table.ItemCount,nil
}

func MakeNewDbSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
		//EndPoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
	}))

	// create a dynamodb instance
	return dynamodb.New(sess)
}

