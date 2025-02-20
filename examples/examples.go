package examples

import (
	"fmt"
	"log"
	"syscall"

	"github.com/rebelnato/gosqlite3/crud"
	"golang.org/x/term"
)

/*
 Below function can be called from parent program to initiate interactive examples of CRUD operations
 How to call ?
 import the "github.com/rebelnato/examples" package
 Call the flow like examples.entryExamplesFlow()
 This function isn't gonna return any value , it will only initiate an interactive loop
 Loop progression depends on users inputs
*/

func entryExamplesFlow() {

	var userId, selectedOption string
	log.Print("Please provide your name : ")
	fmt.Scanln(&userId)
	log.Printf("Hello ! %q , would you like to initiate sqlite CRUD examples ? (Please respond with yes/no)\n> ", userId)
	fmt.Scanln(&selectedOption)
	if selectedOption == "no" {
		log.Println("Examples test program interupted")
		return
	} else {
		initiateExampleTest()
		return
	}
}

/*
 Below function will be called from within entryExamplesFlow if user wishes to proceed with CRUD examples
 How to call ? - Gets called automatically based on input provided in entryExamplesFlow function call

 This function isn't gonna return any value , it will only initiate an interactive loop
 Loop progression depends on users inputs
 As part of the loop user will receive option to perform read , insert , updaet and delete on available data in db
*/

func initiateExampleTest() {

	var exampleInputOption, inputUserName, inputPassword string
	log.Printf("Please select one of the CRUD operations (read,insert,update,delete). Enter exit if you want to exit the testing.\n> ")
	fmt.Scanln(&exampleInputOption)
	switch exampleInputOption {
	case "read":
		var readType string
		log.Println("Starting user data fetch example")
		log.Printf("What would you like to fetch from db ? (alluserslist/singleuserdata) \n> ")
		fmt.Scanln(&readType)
		if readType == "alluserslist" {
			log.Println("Starting process to fetch all available users list")
			users, err := crud.QueryUserList()
			if err != nil {
				log.Printf("Failed to fetch data from db due to error message :%q\n", err)
				initiateExampleTest()
				return
			} else {
				log.Printf("List of users from db is : %q \n\n", users)
				initiateExampleTest()
				return
			}
		} else if readType == "singleuserdata" {
			log.Println("Please provide username of the user to perform search")
			log.Print("User name : ")
			fmt.Scanln(&inputUserName)
			id, username, passwordFromDb, err := crud.QueryData(inputUserName)
			if err != nil {
				log.Printf("Failed to fetch data from db due to error message :%q\n", err)
				log.Println("Reinitiating the example test flow as the provided user is not found")
				initiateExampleTest()
			} else {
				log.Println("Fetched user data is as follows :")
				log.Printf("User name : %q\nIndex ID : %d\nUser password : %q", username, id, passwordFromDb)
			}
		} else {
			log.Println("Invalid input , please type exact command.")
			log.Println("Reinitiating the example test flow as the provided user is not found")
			initiateExampleTest()
			return
		}

		log.Println("") // Just adding an extra blank line for better clarity in terminal output
		initiateExampleTest()
		return
	case "insert":
		log.Println("Starting user data insertion example")
		log.Println("Please provide user name and password in same flow")
		log.Print("User name : ")
		fmt.Scanln(&inputUserName)
		log.Print("Password : ")
		pass, passReadError := term.ReadPassword(int(syscall.Stdin))
		if passReadError != nil {
			log.Printf("Password read failed with error %q", passReadError)
			return
		}
		inputPassword = string(pass)

		insertStatus := crud.InsertData(inputUserName, inputPassword)
		if insertStatus != nil {
			log.Printf("\nFailed to insert with error : %q", insertStatus)
		} else {
			log.Printf("\nSuccessfully inserted data for user : %q", inputUserName)
		}
		log.Println("") // Just adding an extra blank line for better clarity in terminal output
		initiateExampleTest()
		return
	case "update":
		var option, newUsername, newPassword string
		var updateStatus error
		log.Println("Starting user data update example")
		log.Println("Please provide username of the user to perform update")
		log.Print("User name : ")
		fmt.Scanln(&inputUserName)
		log.Printf("please provide the purpose of update task ? (username/password)\n> ")
		fmt.Scanln(&option)
		if option == "username" {
			log.Printf("Proceeding with username update .\nPlease provide new proposed username : ")
			fmt.Scanln(&newUsername)
			tmp := crud.UpdateUsername(inputUserName, newUsername)
			updateStatus = tmp
		} else if option == "password" {
			log.Printf("Proceeding with password update .\nPlease provide new proposed password for username %q : ", inputUserName)
			pass, passReadError := term.ReadPassword(int(syscall.Stdin))
			if passReadError != nil {
				log.Printf("Password read failed with error %q", passReadError)
				return
			}
			newPassword = string(pass)
			tmp := crud.UpdatePassword(inputUserName, newPassword)
			updateStatus = tmp
		}

		if updateStatus != nil {
			log.Printf("Failed to fetch data from db due to error message :%q\n", updateStatus)
			log.Println("Reinitiating the example test flow as the provided user is not found")
			initiateExampleTest()
			return
		} else {
			switch option {
			case "username":
				log.Printf("User name for old user name %q has successfully been updated to new user name %q .", inputUserName, newUsername)
			case "password":
				log.Printf("Password successfully updated for user %q.", inputUserName)
			default:
				log.Println("Invalid option entered, reinitiating the example test flow.")
				initiateExampleTest()
				return
			}
		}
		log.Println("") // Just adding an extra blank line for better clarity in terminal output
		initiateExampleTest()
		return

	case "delete":
		var confirmation string
		log.Println("Starting user data deletion example")
		log.Println("Please provide user name and password for the user whose data needs to be deleted")
		log.Print("User name : ")
		fmt.Scanln(&inputUserName)
		log.Print("Password : ")
		pass, passReadError := term.ReadPassword(int(syscall.Stdin))
		if passReadError != nil {
			log.Printf("Password read failed with error %q", passReadError)
			return
		}
		inputPassword = string(pass)
		log.Printf("Please confim whether you want to delete the data associated with user %q ? (yes/no)", inputUserName)
		fmt.Scanln(&confirmation)
		if confirmation == "yes" {
			deleteStatus := crud.DeleteUser(inputUserName, inputPassword)
			if deleteStatus != nil {
				log.Printf("\nFailed to delete user data with error : %q", deleteStatus)
			}
			log.Println("") // Just adding an extra blank line for better clarity in terminal output
			initiateExampleTest()
			return
		} else if confirmation == "no" {
			log.Printf("User denied deletion confimration, restarting the example test flow\n")
			initiateExampleTest()
			return
		} else {
			log.Printf("Invalid input provided by user, restarting the example test flow\n")
			initiateExampleTest()
			return
		}

	case "exit":
		log.Println("Exiting script based on user input")
		return
	default:
		log.Println("Provided input is invalid reinitiating the process.")
		initiateExampleTest()
		return
	}
}
