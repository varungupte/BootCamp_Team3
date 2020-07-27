# BootCamp_Team3

## Week 1
### Overview
Created a web server using Go programming language. 
It includes the following features:
```
1. Authentication Service for the user.
2. Add a new order
3. Update the dish in a particular order

4. Get the total number of orders.
   -> localhost:5656/order/count
    
5. Get the details of a particular order by orderID.
   -> localhost:5656/order/order_details/order_id/<orderID>

6. Get the details of N number of orders
   -> localhost:5656/order/order_details/tillorder/<N>

7. Query the most popular dish in a particular City.
   -> localhost:5656/populardish/city/<CityName>
```

### Data
**users.csv** contains the data of users:
```
type User struct {
    Id      int
    Name    string
    Street  string
    City    string
    Rating  int
}
```

**restaurants.csv** contains the data of restaurants:
```
type Restaurant struct {
    Id      int
    Name    string
    Street  string
    City    string
    Rating  int
}
```

**orders.csv** contains the data of orders:
```
type Order struct {
    Id            int
    Quantity      int
    Amount        float64
    DishName      string
    User          users.User
    Restau        restaurants.Restaurant
    DeliveryTime  string
}
```

### Steps to run the webserver:
Open Terminal and copy-paste the following commands
```
1. mkdir $HOME/GoWorkspace
2. export GOPATH=$HOME/GoWorkspace
4. mkdir $GOPATH/bin
5. export GOBIN=$GOPATH/bin
6. mkdir $GOPATH/src
7. cd $GOPATH/src
8. go get github.com/varungupte/BootCamp_Team3
9. go get -u github.com/elgs/gojq
10. go get -u github.com/gin-gonic/gin
11. cd $GOPATH/src/github.com/varungupte/BootCamp_Team3/cmd/order-prediction
12. go install .
13. $GOBIN/order-prediction
```

### Project Directory Structure
```
BootCamp_Team3
    ├── cmd
    │   └── order-prediction
    │       ├── Order.csv
    │       ├── Restaurant.csv
    │       ├── User.csv
    │       └── main.go
    └── pkg
        ├── errorutil
        │   └── errorutil.go
        ├── orders
        │   ├── orders.go
        │   └── orders.json
        ├── restaurants
        │   └── restaurants.go
        └── users
            └── users.go
```

#### /cmd
Main applications for this project.
The directory name for each application should match the name of the executable you want to have (e.g., cmd/order-prediction).
Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the /pkg directory. If the code is not reusable or if you don't want others to reuse it, 
put that code in the /internal directory.
It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and nothing else.

#### /pkg
Library code that's ok to use by external applications (e.g., pkg/orders, pkg/restaurants, pkg/users). 
