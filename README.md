# Go gin server

### Overview
Created a web server using Go programming language. 
Refer to the corresponding gRPC server implementation in [gRPC repo](https://github.com/BhaviD/BootCamp_Team3_gRPC)

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

### Technologies used
```
- Go gin
- gRPC
- dynamoDB
- Docker
- Jenkins
- Prometheus
- Grafana
- Godocs
- JWT Auth
- Sonarqube
```

### Models
```
type Address struct {
	HouseNo string
	Street  string
	City    string
	PIN     string
}

type Customer struct {
	Id           uint32
	FullName     string
	Addr         Address
	ActiveStatus bool
}

type Item struct {
	Id       uint32
	Name     string
	Cuisine  string
	Cost     float32
	Quantity uint32
}

type Order struct {
	Id           uint32
	ResId        uint32
	CustId       uint32
	Items        []Item
	Discount     float32
	DeliveryAddr Address
}

type Restaurant struct {
	Id           int64
	Name         string
	Items        []Item
	Addr         Address
	ActiveStatus bool
}
```

### Steps to run the gin server:
Open Terminal and copy-paste the following commands
```
1. mkdir $HOME/GoWorkspace
2. export GOPATH=$HOME/GoWorkspace
4. mkdir $GOPATH/bin
5. export GOBIN=$GOPATH/bin
6. mkdir $GOPATH/src
7. cd $GOPATH/src
8. go get github.com/varungupte/BootCamp_Team3
9. docker build -t img-go-gin:1.1.1 .
10. docker run --rm -p 7878:7878 --net=my_bridge --name=cont-go-gin img-go-gin:1.1.1
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

### Architecture
![image](https://user-images.githubusercontent.com/59866066/89395956-9d3fef80-d72b-11ea-9f67-ae5c8d82f6db.jpeg)
