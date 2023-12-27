package controller

import(
	"github.com/AryanSingh11/carwash/models"
	"encoding/csv"
	"os"
	"fmt"
	"strconv"
	"time"
)

func currTimeHandler() (stringTime string){
	t := time.Now()
	return t.Format("2006-01-02") //yyyy-MM-dd
}

//time will automatically be added //no need to pass it as parameter
func CreateNewCustomer(id int, name string, contact string, vehicleModel string, regnumber string, service string)(err error){
	var newCustomer models.Customer //new customer object

	var CustomerArray  []models.Customer //customer array

	//first load all data from  customer.csv file into array
	file, err := os.Open("customersNew.csv")
	if err != nil {
		fmt.Println("Error loading customer.csv in customerController.go")
		return err
	}
	defer file.Close()


	reader  := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all records from customer.csv in customerController.go in ReadAll")
		return err
	}

	//read all
	for _, record := range records {
		var tempCustomer models.Customer
		tempCustomer.ID, _ = strconv.Atoi(record[0])
		tempCustomer.Name = record[1]
		tempCustomer.Contact = record[2]
		tempCustomer.VehicleModel = record[3]
		tempCustomer.RegNumber = record[4]
		tempCustomer.Service = record[5]
		tempCustomer.TimeCreated = record[6]
		//fmt.Println("temp se: ", tempCustomer.TimeCreated)
		// append
		CustomerArray = append(CustomerArray, tempCustomer)
		// fmt.Println(tempCustomer)
		//fmt.Println(CustomerArray)
	}
	//print
	fmt.Println("read all data from customersNew.csv")
	//after this for loop we have loaded everything form cusotmer.csv to CustomerArray
	//now for new customer
	newCustomer.ID = id
	newCustomer.Name = name
	newCustomer.Contact = contact
	newCustomer.VehicleModel = vehicleModel
	newCustomer.RegNumber = regnumber
	newCustomer.Service = service
	newCustomer.TimeCreated = currTimeHandler()  //returns current date
	// fmt.Println(newCustomer.TimeCreated)
	//finally add
	CustomerArray = append(CustomerArray, newCustomer)
	fmt.Println(CustomerArray)
	// defer file.Close()


	FileForWrite, err := os.Create("customersNew.csv")
	if err != nil {
		fmt.Println("error opening customers.csv")
		return err
	}
	//Now write the new data to customer.csv	
	csvwriter := csv.NewWriter(FileForWrite)
	defer csvwriter.Flush()
	for _, record := range CustomerArray {
		row := []string{ strconv.Itoa(record.ID), record.Name, record.Contact, record.VehicleModel, record.RegNumber, record.Service, record.TimeCreated}
		//fmt.Println(row)
		csvwriter.Write(row)
	}

	//print
	fmt.Println("new customer list is :")
	for _, record := range CustomerArray {
		fmt.Println(record)
	}

	return err
}

func DeleteCustomer(id int)(err error){

	// var newCustomer models.Customer //new customer object
	var CustomerArray  []models.Customer //customer array

	//first load all data from  customer.csv file into array
	file, err := os.Open("customersNew.csv")
	if err != nil {
		fmt.Println("Error loading customerNew.csv in customerController.go")
		return err
	}
	reader  := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all records from customerNew.csv in customerController.go in ReadAll")
		return err
	}

	//read all
	flag := false
	for _, record := range records {
		var tempCustomer models.Customer
		tempCustomer.ID, _ = strconv.Atoi(record[0])
		tempCustomer.Name = record[1]
		tempCustomer.Contact = record[2]
		tempCustomer.VehicleModel = record[3]
		tempCustomer.RegNumber = record[4]
		tempCustomer.Service = record[5]
		tempCustomer.TimeCreated = record[6]
		//now append
		if(tempCustomer.ID!=id){
			//append only when the id is different //this way its delted
			CustomerArray = append(CustomerArray, tempCustomer)
		}else{
			flag = true; //matlab found
		}
	}

	//check if something need to be deleted or not
	if(!flag){
		fmt.Println("no record found for given customer id")
		return nil;
	}
	
	file.Close()

	//now write the new shorter customerArray
	FileForWrite, err := os.Create("customersNew.csv")
	if err != nil {
		fmt.Println("error opening customers.csv")
		return err
	}
	//Now write the new data to customer.csv	
	csvwriter := csv.NewWriter(FileForWrite)
	defer csvwriter.Flush()
	for _, record := range CustomerArray {
		row := []string{ strconv.Itoa(record.ID), record.Name, record.Contact, record.VehicleModel, record.RegNumber}
		csvwriter.Write(row)
	}

	//print
	fmt.Print("aftr delete operation, new data set is ")
	for _, record := range CustomerArray {
		fmt.Println(record)
	}

	return nil
}

//id ->primary key
func UpdateCustomer(id int, name string, contact string, vehicleModel string, regnumber string, service string)(err error){
	fmt.Println("updating under progress")
	//first we'll delete and then insert the new customre object
	//delete
	err = DeleteCustomer(id);
	if err != nil {
		return err
	}
	//insert the new customer
	err = CreateNewCustomer(id, name, contact, vehicleModel, regnumber, service) //no need to pass date
	if err != nil {
		return err
	}

	//printing
	//printing
	//now write the new shorter customerArray
	//first load all data from  customer.csv file into array
	fmt.Println()
	fmt.Println("the customer data set after updating looks like this")
	file, err := os.Open("customersNew.csv")
	if err != nil {
		fmt.Println("Error loading customerNew.csv in customerController.go")
		return err
	}
	reader  := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all records from customerNew.csv in customerController.go in ReadAll")
		return err
	}
	//[print] all rows
	for _, record := range records {
		fmt.Println(record)
	}

	return nil;
}

