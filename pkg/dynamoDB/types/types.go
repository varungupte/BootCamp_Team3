package types

type Address struct {
	HouseNo string
	Street  string
	City    string
	PIN     string
}

type Customer struct {
	Id           int
	FullName     string
	Addr         Address
	ActiveStatus bool
}