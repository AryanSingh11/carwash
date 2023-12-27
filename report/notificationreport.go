package report

import(
	"os"
	"fmt"
	"encoding/csv"
	"time"
	"github.com/AryanSingh11/carwash/models"
)

var NotificationList []string //contains list of dates at which notification was sent

func currTimeHandler() (stringTime string){
	t := time.Now()
	return t.Format("2006-01-02") //yyyy-MM-dd
}

func FetchAllNotificationsForID(inputid string) (err error){ //get all notification for this ID
	file, err := os.Open("notification.csv")
	if err != nil {
		fmt.Println("Error loading notification.csv in notificationreport.go in FetchAllNotificationsForID")
		return err
	}
	defer file.Close()


	reader  := csv.NewReader(file)
	var records [][]string
	records, err = reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all records from notification.csv in notificationreport.go in ReadAll in FetchAllNotificationsForID")
		return err
	}
	for _, record := range records {
		currentid := record[0]
		if(currentid == inputid){ //then retrive all the dates of notification
			NotificationList = append(NotificationList, record[1:]...)
		}
	}

	//print
	fmt.Println("all the notifs for this id were sent on these dates: ")
	for _, record := range records {
		currentid := record[0]
		if(currentid == inputid){ //then retrive all the dates of notification
			fmt.Println(" ")
			for _, date := range record[1:] { //we'd like to exclude 0th elem ie ID
				fmt.Print("coupon notification was sent on this date: ")
				fmt.Print(date)
				// fmt.Println(" ")
			}
			fmt.Println()
		}
	}


	return nil;
}

func CreateNotificationForID(inputid string) (err error){
	currentdate := currTimeHandler()
	var NotificationGrid []models.NotificationList
	//copy all the prev notifs

	file, err := os.Open("notification.csv")
	if err != nil {
		fmt.Println("Error opening notification.csv in notificationreport.go inside CreateNotificationForNewID")
		return err
	}
	defer file.Close()


	reader  := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all records from notification.csv in notificationreport.go in ReadAll in CreateNotificationForNewID")
		return err
	}

	for _, record := range records {
			//now append
			NotificationGrid = append(NotificationGrid, record)
		}
	

	//now append the new notifications object to the notificationsGrid
	var tempRecord []string
	tempRecord = append(tempRecord, inputid)
	tempRecord= append(tempRecord, currentdate) //appending the current date 
	//append temp to the notificationsGrid
	NotificationGrid = append(NotificationGrid, tempRecord);

	//WRITE WRITE
	FileForWrite, err := os.Create("notification.csv")
	if err != nil {
		fmt.Println("error opening notification.csv in notificationreport.go inside CreateNotification")
		return err
	}
	//Now write the new data to notification.csv	
	csvwriter := csv.NewWriter(FileForWrite)
	defer csvwriter.Flush()
	for _, record := range NotificationGrid {
		row := record
		fmt.Println(row)
		csvwriter.Write(row)
	}
	

	///print
	fmt.Println("")
	fmt.Println("the new notification grid after create operation is: ")
	for _, record := range NotificationGrid {
		fmt.Println(record)
	}

	return nil
}


//crete when all notifs are in same row
// func CreateNotificationForID(inputid string) (err error){
// 	currentdate := currTimeHandler()
// 	var NotificationGrid []models.NotificationList
// 	//copy all the prev notifs

// 	file, err := os.Open("notification.csv")
// 	if err != nil {
// 		fmt.Println("Error opening notification.csv in notificationreport.go inside CreateNotificationForNewID")
// 		return err
// 	}
// 	defer file.Close()


// 	reader  := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		fmt.Println("Error reading all records from notification.csv in notificationreport.go in ReadAll in CreateNotificationForNewID")
// 		return err
// 	}

// 	idAlreadyExisted := false
// 	var alreadyExistedNotifs []string
// 	for _, record := range records {
// 		currentid := record[0]
// 		if(currentid == inputid){ //then id already exists
// 			idAlreadyExisted = true
// 			//fmt.Println("the given id already exists in notifications, you can only retrieve it using FetchAllNotificationsForID")
// 			// var temprecord []string
// 			// currentNewDate := currTimeHandler()
// 			alreadyExistedNotifs = append(alreadyExistedNotifs, record[1:]...) //already some date is there is save it
// 			// temprecord = append(temprecord, currentNewDate) //append new date of service
// 			// NotificationGrid = append(NotificationGrid, temprecord)
// 			// for _, date := range record[1:] { //we'd like to exclude 0th elem ie ID
// 			// 	NotificationList = append(NotificationList, date)
// 			// }
// 		}else{
// 			//now append
// 			NotificationGrid = append(NotificationGrid, record)
// 		}
// 	}

// 	//now append the new notifications object to the notificationsGrid
// 	var tempRecord []string
// 	tempRecord= append(tempRecord, inputid)
// 	//check if already existed then append the rest of the data
// 	if(idAlreadyExisted){
// 		tempRecord = append(tempRecord, alreadyExistedNotifs[0:]...)//append the already existing notifs
// 	}
// 	tempRecord= append(tempRecord, currentdate) //appending the current date 
// 	//append temp to the notificationsGrid
// 	NotificationGrid = append(NotificationGrid, tempRecord);

// 	//WRITE WRITE
// 	FileForWrite, err := os.Create("notification.csv")
// 	if err != nil {
// 		fmt.Println("error opening notification.csv in notificationreport.go inside CreateNotification")
// 		return err
// 	}
// 	//Now write the new data to notification.csv	
// 	csvwriter := csv.NewWriter(FileForWrite)
// 	defer csvwriter.Flush()
// 	for _, record := range NotificationGrid {
// 		row := record
// 		fmt.Println(row)
// 		csvwriter.Write(row)
// 	}
	

// 	///print
// 	fmt.Println("")
// 	fmt.Println("the new notification grid after create operation is: ")
// 	for _, record := range NotificationGrid {
// 		fmt.Println(record)
// 	}

// 	return nil
// }