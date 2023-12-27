package handler

import (
	"fmt"
	"github.com/AryanSingh11/carwash/controller"
	"os"
)

func DealerHandler() {
	fmt.Println("these are the methods available for you as a Dealer ")
	fmt.Println("1. Register new Dealer, type registerdealer to access this")
	fmt.Println("2. Create new Dealer, type newdealer to access this")
	fmt.Println("3. Delete Dealer, type deletedealer to access this")
	fmt.Println("3. Update Dealer, type updatedealer to access this")

	var dealerinput string
	fmt.Scanln(&dealerinput)
	if dealerinput == "registerdealer" || dealerinput == "newdealer" {
		//id
		var newdealerid string
		fmt.Println("enter new dealerID")
		fmt.Scanln(&newdealerid)

		//name
		var newdealername string
		fmt.Println("enter new dealer's name")
		fmt.Scanln(&newdealername)

		//current Branch
		var dealercurrentbranch string
		fmt.Println("enter new dealer's current branch")
		fmt.Scanln(&dealercurrentbranch)

		fmt.Println("calling the CreateNewDealer from dealerController.go")
		err := controller.CreateNewDealer(newdealerid, newdealername, dealercurrentbranch)
		if err != nil {
			fmt.Println("error creating new dealer: ", err)
			os.Exit(1)
		} else {
			fmt.Println("success!✅ new dealer created 😊 , checkout the dealer.csv file")
		}

	} else if dealerinput == "deletedealer" {
		//id
		var newdealerid string
		fmt.Println("enter dealerID of dealer to delete")
		fmt.Scanln(&newdealerid)
		fmt.Println("calling DeleteDealerByID from dealerController")
		err := controller.DeleteDealerByID(newdealerid)
		if err != nil {
			fmt.Println("error creating dealer: ", err)
			os.Exit(1)
		} else {
			fmt.Println("success!✅ dealer deleted 😊 , checkout the dealer.csv file")
		}

	} else if dealerinput == "updatedealer" {
		//id
		var newdealerid string
		fmt.Println("enter ID of dealer You'd like to update ")
		fmt.Scanln(&newdealerid)

		//name
		var newdealername string
		fmt.Println("enter updated dealer's name")
		fmt.Scanln(&newdealername)

		//current Branch
		var dealercurrentbranch string
		fmt.Println("enter new dealer's current branch")
		fmt.Scanln(&dealercurrentbranch)

		fmt.Println("calling the UpdateDealerByID from dealerController.go")
		err := controller.UpdateDealerByID(newdealerid, newdealername, dealercurrentbranch)
		if err != nil {
			fmt.Println("error updating dealer: ", err)
			os.Exit(1)
		} else {
			fmt.Println("success!✅ dealer Updated 😊 , checkout the dealer.csv file")
		}
	}
}
