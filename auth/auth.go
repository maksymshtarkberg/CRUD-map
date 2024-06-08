package auth

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	Login    string
	Password string
	Balance  int
}

type UserManager struct {
	users map[int]User
}

func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[int]User),
	}
}

func (um *UserManager) AddNewUser(login, password string) {
	id := len(um.users) + 1
	user := User{
		ID:       id,
		Login:    login,
		Password: password,
		Balance:  0,
	}
	um.users[id] = user
}

func HandleAddUser(um *UserManager) {
	var login string
	var password string

	fmt.Print("Enter user login: ")
	reader := bufio.NewReader(os.Stdin)
	login, _ = reader.ReadString('\n')
	login = strings.TrimSpace(login)

	fmt.Print("Enter employee password: ")
	password, _ = reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if login == "" || password == "" {
		fmt.Print("\nName or password cannot be empty\n")
		return
	}

	um.AddNewUser(login, password)
}

func (um *UserManager) GetUser(id int) (User, bool) {
	user, exists := um.users[id]

	return user, exists
}

func HandleGetUser(um *UserManager) {
	var id string
	for {
		fmt.Print("Enter user id to get user: ")
		fmt.Scanln(&id)

		id, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid numeric user id.")
			continue
		}

		user, found := um.GetUser(id)
		if !found {
			fmt.Println("User not found.")
			return
		}
		fmt.Printf("User: ID:%d, Login:%s, Password:%s, Balance:%d\n", user.ID, user.Login, user.Password, user.Balance)

		return
	}
}

func (um *UserManager) EditUserBalance(user User, balance int) (User, bool) {
	user, exists := um.GetUser(user.ID)

	user.Balance = balance
	um.users[user.ID] = user

	return user, exists
}

func HandleEditUserBalance(um *UserManager) {
	var id string
	var balance string

	for {

		fmt.Print("Enter user id to edit balance: ")
		fmt.Scanln(&id)

		id, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid numeric user id.")
			continue
		}
		user, exists := um.GetUser(id)
		if !exists {
			fmt.Println("User with id", id, "not found")
			return
		}
		fmt.Printf("User: ID:%d, Login:%s, Password:%s, Balance:%d\n", user.ID, user.Login, user.Password, user.Balance)

		for {
			fmt.Print("\nEnter new balance: ")
			fmt.Scanln(&balance)

			balance, err := strconv.Atoi(balance)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid numeric balance.")
				continue
			}

			updatedUser, updated := um.EditUserBalance(user, balance)
			if !updated {
				fmt.Println("Failed to update user balance.")
				return
			}
			fmt.Printf("Updated User: ID:%d, Login:%s, Password:%s, Balance:%d\n", updatedUser.ID, updatedUser.Login, updatedUser.Password, updatedUser.Balance)
			return
		}

	}

}

func (um *UserManager) DeleteUser(id int) bool {
	user, exists := um.GetUser(id)
	if exists {
		fmt.Printf("User: ID:%d, Login:%s, Password:%s, Balance:%d\n", user.ID, user.Login, user.Password, user.Balance)
	}
	delete(um.users, id)
	return exists
}

func HandleDeleteUser(um *UserManager) {
	var id string
	for {
		fmt.Print("Enter user id to delete user: ")
		fmt.Scanln(&id)

		id, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid numeric id.")
			continue
		}

		found := um.DeleteUser(id)
		if !found {
			fmt.Println("User with id", id, "not found")
			return
		}

		fmt.Print("\nUser successfully deleted!")
		return
	}

}
