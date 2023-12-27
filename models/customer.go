package models

// type Customer struct {
// 	ID           int
// 	Name         string
// 	Contact      string
// 	VehicleMake  string
// 	VehicleModel string
// 	RegNumber    string
// 	// Other attributes relevant to a customer
// }

type Customer struct {
	ID           int
	Name         string
	Contact      string
	VehicleModel string
	RegNumber    string
	Service 	 string
	TimeCreated  string 
	// Other attributes relevant to a customer
}

// func NewCustomer(id int, name, contact, make, model, reg string) *Customer {
// 	return &Customer{
// 		ID:           id,
// 		Name:         name,
// 		Contact:      contact,
// 		VehicleMake:  make,
// 		VehicleModel: model,
// 		RegNumber:    reg,
// 	}
// }