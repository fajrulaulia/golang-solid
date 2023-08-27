package srp

import (
	"fmt"
)

type User struct {
	ID      int
	Name    string
	Balance float64
}

type TransactionService struct{}

func (ts TransactionService) Deposit(user *User, amount float64) {
	user.Balance += amount
	fmt.Printf("Deposit dana sebesar %.2f ke %s\n", amount, user.Name)
}

func (ts TransactionService) Withdraw(user *User, amount float64) {
	if user.Balance >= amount {
		user.Balance -= amount
		fmt.Printf("Withdrawn %.2f from %s\n", amount, user.Name)
	} else {
		fmt.Println("Insufficient balance")
	}
}

type UserService struct {
	Users []*User
}

func (us *UserService) AddUser(user *User) {
	us.Users = append(us.Users, user)
}

func (us *UserService) GetUserByID(id int) *User {
	for _, user := range us.Users {
		if user.ID == id {
			return user
		}
	}
	return nil
}
