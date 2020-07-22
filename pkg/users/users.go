package users

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	Id int
	Name string
	Street string
	City string
	Rating int
}

func GetUsers(filename string) []User {
	userFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer userFile.Close()

	reader := csv.NewReader(userFile)
	reader.FieldsPerRecord = -1

	userData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var user User
	var users []User

	for _, each := range userData {
		user.Id,_ = strconv.Atoi(each[0])
		user.Name = each[1]
		user.Street= each[2]
		user.City= each[3]
		user.Rating,_=strconv.Atoi(each[4])
		users = append(users, user)
	}
	return users
}