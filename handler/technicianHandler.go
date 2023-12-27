package handler

import (
	"fmt"
	"github.com/AryanSingh11/carwash/controller"
	"os"
)

func TechnicianHandler() {
	fmt.Println("these are the methods available for you as a Technician ")
	fmt.Println("1. Register new Technician, type registerTechnician to access this")
	fmt.Println("2. Create new Technician, type newTechnician to access this")
	fmt.Println("3. Delete Technician, type deleteTechnician to access this")
	fmt.Println("3. Update Technician, type updateTechnician to access this")

	var Technicianinput string
	fmt.Scanln(&Technicianinput)
	if Technicianinput == "registerTechnician" || Technicianinput == "newTechnician" {
		//id
		var newTechnicianid string
		fmt.Println("enter new TechnicianID")
		fmt.Scanln(&newTechnicianid)

		//name
		var newTechnicianname string
		fmt.Println("enter new Technician's name")
		fmt.Scanln(&newTechnicianname)

		//current Branch
		var Techniciancurrentbranch string
		fmt.Println("enter new Technician's current branch")
		fmt.Scanln(&Techniciancurrentbranch)

		fmt.Println("calling the CreateNewTechnician from technicianController.go")
		err := controller.CreateNewTechnician(newTechnicianid, newTechnicianname, Techniciancurrentbranch)
		if err != nil {
			fmt.Println("error creating new Technician: ", err)
			os.Exit(1)
		} else {
			fmt.Println("success!✅ new Technician created 😊 , checkout the technician.csv file")
		}

	} else if Technicianinput == "deleteTechnician" {
		//id
		var newTechnicianid string
		fmt.Println("enter TechnicianID of Technician to delete")
		fmt.Scanln(&newTechnicianid)
		fmt.Println("calling DeleteTechnicianByID from TechnicianController")
		err := controller.DeleteTechnicianByID(newTechnicianid)
		if err != nil {
			fmt.Println("error creating Technician: ", err)
			os.Exit(1)
		} else {
			fmt.Println("success!✅ Technician deleted 😊 , checkout the technician.csv file")
		}

	} else if Technicianinput == "updateTechnician" {
		//id
		var newTechnicianid string
		fmt.Println("enter ID of Technician You'd like to update ")
		fmt.Scanln(&newTechnicianid)

		//name
		var newTechnicianname string
		fmt.Println("enter updated Technician's name")
		fmt.Scanln(&newTechnicianname)

		//current Branch
		var Techniciancurrentbranch string
		fmt.Println("enter new Technician's current branch")
		fmt.Scanln(&Techniciancurrentbranch)

		fmt.Println("calling the UpdateTechnicianByID from TechnicianController.go")
		err := controller.UpdateTechnicianByID(newTechnicianid, newTechnicianname, Techniciancurrentbranch)
		if err != nil {
			fmt.Println("error updating Technician: ", err)
			os.Exit(1)
		} else {
			fmt.Println("success!✅ Technician Updated 😊 , checkout the technician.csv file")
		}
	}
}
