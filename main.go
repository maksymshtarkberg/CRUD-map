package main

import (
	"fmt"

	"github.com/maksymshtarkberg/CRUD-map/auth"
)

func main() {
	HandleUserAuth()
}

func HandleUserAuth() {

	userManager := auth.NewUserManager()

	for {
		fmt.Println("\nUserAuthManager")
		fmt.Println("\nOptions:")
		fmt.Print("[1] Create new user\n")
		fmt.Print("[2] Get user by id\n")
		fmt.Print("[3] Edit user balance by id\n")
		fmt.Print("[4] Delete user by id\n")
		fmt.Print("[5] Exit\n")
		fmt.Print("\nChoose an option: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			auth.HandleAddUser(userManager)
		case 2:
			auth.HandleGetUser(userManager)
		case 3:
			auth.HandleEditUserBalance(userManager)
		case 4:
			auth.HandleDeleteUser(userManager)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
			return
		}
	}

}
