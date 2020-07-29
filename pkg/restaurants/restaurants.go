package restaurants

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Restaurant struct {
	Id int
	Name string
	Street string
	City string
	Rating int
}

// GetRestaurants reads the restaurant information from "Restaurants.csv".
// It returns the slice of Restaurants
func GetRestaurants(filename string) []Restaurant {
	restaurantFile, err := os.Open("Restaurant.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer restaurantFile.Close()

	reader := csv.NewReader(restaurantFile)
	reader.FieldsPerRecord = -1

	restaurantData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var res Restaurant
	var restaurants []Restaurant

	for _, each := range restaurantData {
		res.Id,_ = strconv.Atoi(each[0])
		res.Name = each[1]
		res.Street= each[2]
		res.City= each[3]
		res.Rating,_=strconv.Atoi(each[4])
		restaurants = append(restaurants, res)
	}
	return restaurants
}