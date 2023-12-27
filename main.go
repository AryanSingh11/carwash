package main

import (
	// "encoding/csv"
	// "log"
	"fmt"
	"moul.io/banner"
	// "os"

	// "github.com/AryanSingh11/carwash/controller"
	"github.com/AryanSingh11/carwash/handler"
	"github.com/AryanSingh11/carwash/emailservice"
	// "github.com/AryanSingh11/carwash/report"
	// "github.com/AryanSingh11/carwash/report"
	// "github.com/AryanSingh11/controller"
)

func main() {
	fmt.Println(banner.Inline("amena garage"))
	fmt.Println("You can log in as: customer , dealer, technician")
	fmt.Println("You can also access these 2 types of reports: ")
	//reports
	fmt.Println("1)Given a date, a report showing customer information, car, type of service and last date of service will be geneated ")
	fmt.Println("2)Given customer id, a report showing details of last coupon notification for each of the service will be geenerated ")

	fmt.Println("")
	fmt.Println("to get report 1 : type reportbydate")
	fmt.Println("to get report 2 : type notificationreport")

	//login profiles
	fmt.Println("")
	fmt.Println("to login as customer, type : customer")
	fmt.Println("to login as dealer, type : dealer")
	fmt.Println("to login as technician, type : technician")

	//email
	fmt.Println("")
	fmt.Println("To Run the Email service that send Previous Customers coupon notification")
	fmt.Println("if the last oilchange date was 6 months ago, type : email")

	var userinput string
	fmt.Scanln(&userinput)
	//tell them how to exit
	fmt.Println("")
	fmt.Println("you have entered the program, to exit, type x and press enter ")

	if userinput == "customer" {
		handler.CustomerHandler()
	}else if userinput == "dealer" {
		handler.DealerHandler()
	}else if userinput == "technician" {
		handler.TechnicianHandler()
	}else if userinput == "reportbydate"{
		handler.GetAllServicesOnThisDate()
	}else if userinput == "notificationreport"{
		handler.MainNotificationSectionHandler()
	}else if userinput == "email" {
		emailservice.EmailFinder()
	}

}

// func main(){
// 	d := "2023-11-28"
// 	report.AllServicesOnThisDate(d)
// 	fmt.Println(report.ServiceList)
// }

// func main(){ // TODO: implement correct
// 	id := "5678"
// 	report.FetchAllNotificationsForID(id)
// 	fmt.Println(report.NotificationList)
// }

// func main(){
// 	// id := "5678"
// 	// report.CreateNotificationForNewID(id)
// 	id2 := "87865"
// 	report.CreateNotificationForNewID(id2)
// }

// func main(){
// 	//Create a program
// 	//that will send a promotion (or reminder) to the vehicle owner for oil change 6 months
// 	//since the last oil change.
// 	// emailservice.EmailFinder() //automatically find the email past 6 months
// 	// fmt.Println("final email list is ")
// 	// fmt.Println(emailservice.EmailAddrList)

// 	 id := 345345;
// 	 name := "bob"
// 	 contact := "@@bob"
// 	 vehicleModel := "jeep"
// 	 regnumber := "9000"
// 	 service := "oilchange"

// 	 err := controller.CreateNewCustomer(id, name, contact, vehicleModel, regnumber, service)
// 	 if err == nil {
// 		fmt.Println("new customer created successfully")
// 	 }else{
// 		fmt.Println(err)
// 		os.Exit(1);
// 	}

// 	//  err = controller.DeleteCustomer(255)
// 	//  if err != nil {
// 	// 	fmt.Print("delete customer failed")
// 	//  }

// 	//  err = controller.UpdateCustomer(25, "hawai", "email", "bolero", "mp0432", "paint change")
// 	//  if err != nil {
// 	// 	fmt.Print("update customer failed, error is ", err)
// 	// 	os.Exit(1);
// 	//  }

// }

// func main(){
// 	id:="11"
// 	// name:= "akeal"
// 	// branch:="iowa"
// 	err := controller.DeleteDealerByID(id)
// 	if err != nil {
// 		fmt.Println("delete newDealer failed, error is ", err)
// 	}

// 	name:= "bhavesh"
// 	branch:="minnesota"
// 	err = controller.UpdateDealerByID(id, name, branch)

// }

// func main(){
// 	id:="11"
// 	// name:= "akeal"
// 	// branch:="iowa"
// 	err := controller.DeleteTechnicianByID(id)
// 	if err != nil {
// 		fmt.Println("delete newDealer failed, error is ", err)
// 	}

// 	id2:="33"
// 	name:= "bhavesh"
// 	branch:="minnesota"
// 	err = controller.UpdateTechnicianByID(id2, name, branch)
// 	if err != nil {
// 		fmt.Println("update Technician failed, error is ", err)
// 	}

// }
