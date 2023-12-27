package controller

import(
	"fmt"
	"os"
	"encoding/csv"
	"github.com/AryanSingh11/carwash/models"
	// "strconv"
)

func CreateNewTechnician(inputid string, inputname string, currentinputbranch string ) (err error) {
	//copy the old ones

	var TechnicianList []models.Technician //Technician array

	//first load all data from  technician.csv file into array
	file, err := os.Open("technician.csv")
	if err != nil {
		fmt.Println("Error loading technician.csv in technicianController.go")
		return err
	}
	defer file.Close()


	reader  := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all records from technician.csv in technicianController.go")
		return err
	}

	//read all
	for _, record := range records {
		var tempTechnician models.Technician
		tempTechnician.ID = record[0]
		tempTechnician.Name = record[1]
		tempTechnician.CurrentBranch = record[2]
		//fmt.Println("temp se: ", temptechnician.TimeCreated)
		// append
		TechnicianList = append(TechnicianList, tempTechnician)
		// fmt.Println(temptechnician)
		//fmt.Println(technicianArray)
	}

	//print
	fmt.Println("reading all data from techniciansNew.csv")
	//after this for loop we have loaded everything form technician.csv to technicianArray
	//now for new technician
	var newTechnician models.Technician //new technician object
	newTechnician.ID = inputid
	newTechnician.Name = inputname
	newTechnician.CurrentBranch = currentinputbranch
	// fmt.Println(newtechnician.TimeCreated)
	//finally add
	TechnicianList = append(TechnicianList, newTechnician)
	// fmt.Println(technicianArray)
	// defer file.Close()


	FileForWrite, err := os.Create("technician.csv")
	if err != nil {
		fmt.Println("error opening technician.csv")
		return err
	}
	//Now write the new data to technician.csv	
	csvwriter := csv.NewWriter(FileForWrite)
	defer csvwriter.Flush()
	for _, record := range TechnicianList {
		row := []string{ record.ID, record.Name, record.CurrentBranch}
		//fmt.Println(row)
		csvwriter.Write(row)
	}

	//print
	fmt.Println("new Technician list is : ")
	for _, record := range TechnicianList {
		fmt.Println(record)
	}

	return nil
}

func DeleteTechnicianByID(inputid string, )(err error){

	// var newCustomer models.Customer //new Technician object
	var TechnicianList []models.Technician //customer array

	//first load all data from  technician.csv file into Technicianlist array
	file, err := os.Open("technician.csv")
	if err != nil {
		fmt.Println("Error loading technician.csv in TechnicianController.go")
		return err
	}
	reader  := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all records from technician.csv in TechnicianController.go in delete")
		return err
	}

	//read all
	//read all
	flag := false
	for _, record := range records {
		var tempTechnician models.Technician
		tempTechnician.ID = record[0]
		tempTechnician.Name = record[1]
		tempTechnician.CurrentBranch = record[2]
		// append if found
		if(tempTechnician.ID!=inputid){
			//append only when the id is different //this way its deleted
			TechnicianList = append(TechnicianList, tempTechnician)
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
		fmt.Println("no record found for given Technician id")
		return nil;
	}
	
	file.Close()

	//now write the new shorter Technicianlist
	FileForWrite, err := os.Create("technician.csv")
	if err != nil {
		fmt.Println("error opening technician.csv")
		return err
	}
	//Now write the new data to technician.csv	
	csvwriter := csv.NewWriter(FileForWrite)
	defer csvwriter.Flush()
	for _, record := range TechnicianList {
		row := []string{ record.ID, record.Name, record.CurrentBranch}
		csvwriter.Write(row)
	}


	//print the new
	fmt.Println("the new Technician list after deletion is")
	for _, record := range TechnicianList {
		fmt.Println(record)
	}

	return err
}

func UpdateTechnicianByID(inputid string, inputname string, currentinputbranch string)(err error){
	err = DeleteTechnicianByID(inputid)
	if err != nil {
		fmt.Println("error while DeletingbyID in UpdateTechnicianbyID: ", err)
		return err
	}
	err = CreateNewTechnician(inputid, inputname,currentinputbranch)
	if err != nil {
		fmt.Println("error while CreateNewTechnician in UpdateTechnicianbyID: ", err)
		return err
	}

	//print
	//print
	//print
	fmt.Println("the Technician list after Updating is: ")
	file, err := os.Open("technician.csv")
	if err != nil {
		fmt.Println("Error loading technician.csv in TechnicianController.go")
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