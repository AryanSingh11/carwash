package handler

import (
	"fmt"
	"github.com/AryanSingh11/carwash/controller"
	"os"
)

func CustomerHandler() {
	fmt.Println("these are the methods available for you as a customer")
	fmt.Println("1. Register new Customer, type register to access this")
	fmt.Println("2. Create new Customer, type newcustomer to access this")
	fmt.Println("3. Delete Customer, type delete to access this")
	fmt.Println("3. Update Customer, type update to access this")

	var customerinput string
	fmt.Scanln(&customerinput)
	if customerinput == "register" || customerinput == "newcustomer" {
		//id
		var newcustomerid int
		fmt.Println("enter new customerID")
		fmt.Scanln(&newcustomerid)

		//name
		var newcustomername string
		fmt.Println("enter new customer's name")
		fmt.Scanln(&newcustomername)

		//email
		var newcustomermail string
		fmt.Println("enter new customer's contact email address")
		fmt.Scanln(&newcustomermail)

		//vehicle model
		var newcustomervehicle string
		fmt.Println("enter new customer's vehicle model")
		fmt.Scanln(&newcustomervehicle)

		//regnumber
		var newcustomerregnumber string
		fmt.Println("enter new customer's vehicle number")
		fmt.Scanln(&newcustomerregnumber)

		//service
		var newcustomerservice string
		fmt.Println("enter new service requested")
		fmt.Scanln(&newcustomerservice)

		fmt.Println("calling the CreateNewCustomer from cutomerController.go")
		err := controller.CreateNewCustomer(newcustomerid, newcustomername, newcustomermail, newcustomervehicle, newcustomerregnumber, newcustomerservice)
		if err != nil {
			fmt.Println("error creating new customer: ", err)
			os.Exit(1)
		} else {
			fmt.Println("success!✅ new customer created 😊 , checkout the customersNew.csv file")
		}
	} else if customerinput == "delete" {
		//id
		var newcustomerid int
		fmt.Println("enter customerID of customer to delete")
		fmt.Scanln(&newcustomerid)
		fmt.Println("calling deleteCustomer from customerController")
		err := controller.DeleteCustomer(newcustomerid)
		if err != nil {
			fmt.Println("error creating customer: ", err)
			os.Exit(1)
		} else {
			fmt.Println("success!✅ customer deleted 😊 , checkout the customersNew.csv file")
		}

	} else if customerinput == "update" {
		//id
		var newcustomerid int
		fmt.Println("enter the ID/customerID of the customer You'd like to update")
		fmt.Scanln(&newcustomerid)

		//name
		var newcustomername string
		fmt.Println("enter updated customer's name")
		fmt.Scanln(&newcustomername)

		//email
		var newcustomermail string
		fmt.Println("enter updated customer's contact email address")
		fmt.Scanln(&newcustomermail)

		//vehicle model
		var newcustomervehicle string
		fmt.Println("enter updated customer's vehicle model")
		fmt.Scanln(&newcustomervehicle)

		//regnumber
		var newcustomerregnumber string
		fmt.Println("enter updated customer's vehicle number")
		fmt.Scanln(&newcustomerregnumber)

		//service
		var newcustomerservice string
		fmt.Println("enter updated service requested")
		fmt.Scanln(&newcustomerservice)

		fmt.Println("calling the UpdateCustomer from cutomerController.go")
		err := controller.UpdateCustomer(newcustomerid, newcustomername, newcustomermail, newcustomervehicle, newcustomerregnumber, newcustomerservice)
		if err != nil {
			fmt.Println("error updating customer: ", err)
			os.Exit(1)
		} else {
			fmt.Println("success!✅ customer Updated 😊 , checkout the customersNew.csv file")
		}
	}
}
