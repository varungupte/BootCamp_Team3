package types

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
