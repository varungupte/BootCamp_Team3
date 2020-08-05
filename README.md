# BootCamp_Team3

## Week 1
### Overview
Created a web server using Go programming language. 
It includes the following features:
```
POST /admin/login

POST /restaurant
DELETE /restaurant/{restId}/item/{itemId}
DELETE /restaurant/{restId}
PUT /restaurant/{restId}/item
GET /restaurant/count
GET /restaurant/{restName}
GET /restaurant/{restId}
GET /restaurant/{restId}/items
GET /restaurant/{restId}/items?priceMin={min}&priceMax={max}

POST /order
GET /order/count
PUT /order/{orderId}/item/{itemId}/quantity/{quantity}
DELETE /order/{orderId}/item/{itemId}/quantity/{quantity}

GET /customers/all
GET /customers/id/{custId}
GET /customers/count
POST /customers/new
DELETE /customers/id/{custId}
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

### Steps to run the GIN and gRPC servers:
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
11. go get -u google.golang.org/grpc
12. cd $GOPATH/src/github.com/varungupte/BootCamp_Team3/cmd/order-prediction
13. go run main.go              <-- this will run gin server
14. Open a new terminal tab
15. cd $GOPATH/src/github.com/varungupte/BootCamp_Team3/pkg/services/orders/orders_server
16. go run orders_server.go     <-- this will run gRPC server
```

### Project Directory Structure
```
BootCamp_Team3
    ├── Dockerfile
    ├── Jenkinsfile
    ├── README.md
    ├── cmd
    │   └── gin_server
    │       └── main.go
    ├── go.mod
    ├── go.sum
    └── pkg
        ├── auth
        │   └── jwtAuth.go
        ├── customers_client
        │   ├── customers_client.go
        │   └── customers_client_test.go
        ├── errorutil
        │   └── errorutil.go
        ├── grpcUtil
        │   └── grpcUtil.go
        ├── orders_client
        │   ├── orders_client.go
        │   └── orders_client_test.go
        └── restaurants_client
            ├── restaurants_client.go
            └── restaurants_client_test.go
```

#### /cmd
Main applications for this project.
The directory name for each application should match the name of the executable you want to have (e.g., cmd/order-prediction).
Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the /pkg directory. If the code is not reusable or if you don't want others to reuse it, 
put that code in the /internal directory.
It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and nothing else.

#### /pkg
Library code that's ok to use by external applications (e.g., pkg/orders, pkg/restaurants, pkg/users). 

### Request Flow

![image](https://user-images.githubusercontent.com/59866066/88765441-1c657e80-d194-11ea-9f5c-17054dfc3d5e.png)

Architecture
![image](https://user-images.githubusercontent.com/59866066/89395956-9d3fef80-d72b-11ea-9f67-ae5c8d82f6db.jpeg)
