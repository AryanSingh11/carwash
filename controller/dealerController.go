package controller

import(
	"fmt"
	"os"
	"encoding/csv"
	"github.com/AryanSingh11/carwash/models"
	// "strconv"
)

// var DealerList []models.Dealer

func CreateNewDealer(inputid string, inputname string, currentinputbranch string ) (err error) {
	//copy the old ones

	var DealerList []models.Dealer //dealer array

	//first load all data from  dealer.csv file into array
	file, err := os.Open("dealer.csv")
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
		var tempDealer models.Dealer
		tempDealer.ID = record[0]
		tempDealer.Name = record[1]
		tempDealer.CurrentBranch = record[2]
		//fmt.Println("temp se: ", tempCustomer.TimeCreated)
		// append
		DealerList = append(DealerList, tempDealer)
		// fmt.Println(tempCustomer)
		//fmt.Println(CustomerArray)
	}

	//print
	fmt.Println("reading all data from customersNew.csv")
	//after this for loop we have loaded everything form cusotmer.csv to CustomerArray
	//now for new customer
	var newDealer models.Dealer //new customer object
	newDealer.ID = inputid
	newDealer.Name = inputname
	newDealer.CurrentBranch = currentinputbranch
	// fmt.Println(newCustomer.TimeCreated)
	//finally add
	DealerList = append(DealerList, newDealer)
	// fmt.Println(CustomerArray)
	// defer file.Close()


	FileForWrite, err := os.Create("dealer.csv")
	if err != nil {
		fmt.Println("error opening dealer.csv")
		return err
	}
	//Now write the new data to dealer.csv	
	csvwriter := csv.NewWriter(FileForWrite)
	defer csvwriter.Flush()
	for _, record := range DealerList {
		row := []string{ record.ID, record.Name, record.CurrentBranch}
		//fmt.Println(row)
		csvwriter.Write(row)
	}

	//print
	fmt.Println("new dealer list is :")
	for _, record := range DealerList {
		fmt.Println(record)
	}

	return nil
}

func DeleteDealerByID(inputid string, )(err error){

	// var newCustomer models.Customer //new dealer object
	var DealerList []models.Dealer //customer array

	//first load all data from  dealer.csv file into Dealerlist array
	file, err := os.Open("dealer.csv")
	if err != nil {
		fmt.Println("Error loading dealer.csv in dealerController.go")
		return err
	}
	reader  := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all records from dealer.csv in dealerController.go in delete")
		return err
	}

	//read all
	//read all
	flag := false
	for _, record := range records {
		var tempDealer models.Dealer
		tempDealer.ID = record[0]
		tempDealer.Name = record[1]
		tempDealer.CurrentBranch = record[2]
		// append if found
		if(tempDealer.ID!=inputid){
			//append only when the id is different //this way its deleted
			DealerList = append(DealerList, tempDealer)
		}else{
			flag=true; //means something found
		}
	}
	// flag := false
	// for _, record := range records {
	// 	var  models.Customer
	// 	tempCustomer.ID, _ = strconv.Atoi(record[0])
	// 	tempCustomer.Name = record[1]
	// 	tempCustomer.Contact = record[2]
	// 	tempCustomer.VehicleModel = record[3]
	// 	tempCustomer.RegNumber = record[4]
	// 	tempCustomer.Service = record[5]
	// 	tempCustomer.TimeCreated = record[6]
	// 	//now append
	// 	if(tempCustomer.ID!=id){
	// 		//append only when the id is different //this way its delted
	// 		CustomerArray = append(CustomerArray, tempCustomer)
	// 	}else{
	// 		flag = true; //matlab found
	// 	}
	// }

	//check if something need to be deleted or not
	if(!flag){
		fmt.Println("no record found for given dealer id")
		return nil;
	}
	
	file.Close()

	//now write the new shorter dealerlist
	FileForWrite, err := os.Create("dealer.csv")
	if err != nil {
		fmt.Println("error opening dealer.csv")
		return err
	}
	//Now write the new data to dealer.csv	
	csvwriter := csv.NewWriter(FileForWrite)
	defer csvwriter.Flush()
	for _, record := range DealerList {
		row := []string{ record.ID, record.Name, record.CurrentBranch}
		csvwriter.Write(row)
	}


	//print the new
	fmt.Println("the new Dealer list after deletion is")
	for _, record := range DealerList {
		fmt.Println(record)
	}

	return err
}

func UpdateDealerByID(inputid string, inputname string, currentinputbranch string)(err error){
	err = DeleteDealerByID(inputid)
	if err != nil {
		fmt.Println("error while DeletingbyID in UpdateDealerbyID: ", err)
		return err
	}
	err = CreateNewDealer(inputid, inputname,currentinputbranch)
	if err != nil {
		fmt.Println("error while CreateNewDealer in UpdateDealerbyID: ", err)
		return err
	}

	//print
	//print
	//print
	fmt.Println("the dealer list after Updating is: ")
	file, err := os.Open("dealer.csv")
	if err != nil {
		fmt.Println("Error loading dealer.csv in dealerController.go")
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
		fmt.Println(record)
	}
	//printover

	return nil
	
}