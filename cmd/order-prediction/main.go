package main

import (
	"encoding/json"
	"fmt"
	"github.com/varungupte/BootCamp_Team3/pkg/orders"
	"os"
)

func main() {
	ordrs := orders.GetOrders("Order.csv")

	// Convert to JSON
	jsonData, err := json.Marshal(ordrs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("./data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
