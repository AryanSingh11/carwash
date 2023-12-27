package emailservice

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/AryanSingh11/carwash/report"
	// "math"
)

// type EmailAddr struct{
// 	Emailaddress  string
// }



func greaterThan6months(lastdateofservice string) (ans bool){
	//lastlastdateofservice is of format 2018-06-01
	//find the month part
	month := lastdateofservice[5:7]
	year := lastdateofservice[0:4]
	if(month[0]	== '0'){
		month = month[1:]
	}
	//fmt.Println("month: ", month)  //06 -> 6
	//fmt.Println("year: ", year)  //06 -> 6

	prevmonthinInt, err := strconv.Atoi(month) //atoi converts string to int
	if err != nil {
		fmt.Println("error converting prevmonth to int in Emailsender.go greaterThan6months function")
	}
	prevYearinInt, err := strconv.Atoi(year) //atoi converts string to int
	if err != nil {
		fmt.Println("error converting prevyear to int in Emailsender.go greaterThan6months function")
	}
	// prevmonthinInt, err := strconv.ParseInt(month, 10, 32);
	// prevmonthinInt, err := time.Parse("2006-01-02", lastdateofservice);
	// prevmonthinString := prevmonthinInt.Month().String()
	// if err != nil {
	// 	fmt.Println("error converting month to int in Emailsender.go using parseInt function")
	// }

	//find the current time and date
	// t := time.Now()
	// currMonth := t.Month().String()
	// currMonthinInt, err := strconv.ParseInt(currMonth, 10, 32);
	currentYearinInt := time.Now().Year()
	//fmt.Println("currentYear: ", currentYearinInt)
	// if err != nil {
	// 	fmt.Println("error converting currYearinInt to int in Emailsender.go using time.Now().Year()")
	// }
	currentMonthinInt := int(time.Now().Month())
	//fmt.Println("currentMonth: ", currentMonthinInt)
	// if err != nil {
	// 	fmt.Println("error converting currMonthinInt to int in Emailsender.go using time.Now().Month()")
	// }

	//check
	ansbool := false;
	if(currentYearinInt==prevYearinInt){
		var diff int = currentMonthinInt-prevmonthinInt
		if(diff>6){
			ansbool = true
		}
	}else if(currentYearinInt>prevYearinInt){
		//prev year
		ansbool=true;
	}

	return ansbool
}

var EmailAddrList []string //array of email addresses

func EmailFinder()(err error){
	file, err := os.Open("customersNew.csv")
	if err != nil {
		fmt.Println("error opening customerNew.csv in Emailsender.go ", err)
		return err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("errror reading customerNew.csv in Emailsender.go ", err)
		return err
	}
	for _, record := range records {
		lastdateofservice := record[6]
		isrequired := greaterThan6months(lastdateofservice)
		if(isrequired){ //isrequierd==true then send email since date > 6 months
			currentEmailAddr := record[2] //third field is the email address
			EmailAddrList = append(EmailAddrList, currentEmailAddr)
			fmt.Println("sent coupon email to ",record[3], " at: ", currentEmailAddr)

			//update the notification csv
			report.CreateNotificationForID(record[0])
		}
	}

	fmt.Println("")
	fmt.Println("since new coupon notifications are sent, the notification.csv file is updated with new coupon record")
	

	return nil;
}