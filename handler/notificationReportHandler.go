package handler

import(
	"fmt"
	"github.com/AryanSingh11/carwash/report"

)

func GetAllNotificationOnID(){
	var customerIDinput string
	fmt.Println("enter the customerID for which you want list of notifications")
	fmt.Scanln(&customerIDinput)
	fmt.Println("calling FetchAllNotificationsForID form reports package in notificationreport.go ")
	report.FetchAllNotificationsForID(customerIDinput)
}

func CreateNewNotificationOnID(){
	var customerIDinput string
	fmt.Println("enter the customerID for which you want to create notification")
	fmt.Scanln(&customerIDinput)
	fmt.Println("calling CreateNotificationForID from reports package in notificationreport.go ")
	report.CreateNotificationForID(customerIDinput)
}

func MainNotificationSectionHandler() {
	fmt.Println("you have 2 methods available: ")
	fmt.Println("1. GetAllNotificationOnID: enter and ID and we will send list of all coupon notifications sent to this ID ")
	fmt.Println("2. CreateNewNotificationOnID: lets you create new notification for the CustoerID entered ")
	fmt.Println("for 1: type getallnotificationonid ")
	fmt.Println("for 2: type createnewnotificationonid ")
	fmt.Println("note that method 2 would automatically capture today's date")
	var customerIDinput string
	fmt.Scanln(&customerIDinput)
	if customerIDinput == "getallnotificationonid" {
		GetAllNotificationOnID()	
	}else if customerIDinput == "createnewnotificationonid"{
		CreateNewNotificationOnID()
	}
	


}