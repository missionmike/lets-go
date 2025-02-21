package api

// User user type
type User struct {
	ID   int64
	Name string
	Addr *Address
}

// Address address type
type Address struct {
	City   string
	ZIP    int
	LatLng [2]float64
}

var michael = User{
	ID:   1,
	Name: "Michael",
	Addr: &Address{
		City:   "Newport",
		ZIP:    97365,
		LatLng: [2]float64{44.6368, 124.0535},
	},
}

// Hello writes a welcome string
func Hello() string {
	return "Hello, " + michael.Name
}
