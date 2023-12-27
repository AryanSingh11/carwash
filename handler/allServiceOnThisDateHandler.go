package handler

import(
	"fmt"
	"github.com/AryanSingh11/carwash/report"
)

func GetAllServicesOnThisDate()(err error) {
	var inputdate string
	fmt.Println("enter the date in YYYY-MM-DD format to the services information on this date")
	fmt.Scanln(&inputdate)
	fmt.Println("calling AllServicesOnThisDate from report package ")
	fmt.Println("")
	err = report.AllServicesOnThisDate(inputdate)
	if err != nil {
		fmt.Println("error getting services on this date in allServicesOnThisDateHandler:", err)
		return err
	}

	return nil;
}