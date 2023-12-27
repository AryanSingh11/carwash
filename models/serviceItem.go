package models

//25,loki,@@newo,mustang,123,oilchange,2018-06-01
type ServiceItem struct {
	ID           int
	Name         string
	Contact      string
	VehicleModel string
	RegNumber    string
	Service 	 string
	Lastdateofservice  string
}