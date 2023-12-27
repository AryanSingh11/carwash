package report

import (
	"encoding/csv"
	"fmt"
	"os"
	"github.com/AryanSingh11/carwash/models"
	"strconv"
)

// //25,loki,@@newo,mustang,123,oilchange,2018-06-01
// type ServiceItem struct {
// 	ID           int
// 	Name         string
// 	Contact      string
// 	VehicleModel string
// 	RegNumber    string
// 	Service 	 string
// 	lastdateofservice  string
// }

var ServiceList models.ServiceItemList // list of service performed on that paricular date 

func AllServicesOnThisDate(inputdate string)(err error){

	file, err := os.Open("customersNew.csv")
	if err != nil {
		fmt.Println("Error loading customers.csv in customerController.go")
		return err
	}
	defer file.Close()

	reader  := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all records from customer.csv in customerController.go in ReadAll")
		return err
	}

	//read all and try to match the date
	for _, record := range records {
		servicedate := record[6]
		if(servicedate == inputdate){
			var ServiceItem models.ServiceItem
			ServiceItem.ID, _ = strconv.Atoi(record[0])
			ServiceItem.Name = record[1]
			ServiceItem.Contact = record[2]
			ServiceItem.VehicleModel = record[3]
			ServiceItem.RegNumber = record[4]
			ServiceItem.Service = record[5]
			ServiceItem.Lastdateofservice = record[6]
			
			// append
			ServiceList = append(ServiceList, ServiceItem)
			// fmt.Println(tempCustomer)
			//fmt.Println(ServiceItem)
		}
	}

	//print
	fmt.Println("all services performed on this date were: ")
	for _, service := range ServiceList {
		fmt.Println(service)
	}

	return nil;
}