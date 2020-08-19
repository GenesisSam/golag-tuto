package bankutils

import (
	"fmt"
)

// User default user struct
type User struct {
	id      int
	name    string
	account Account
}

var _users = make(map[int]User)

// CreateUser create new user data and return itself
func CreateUser(name string, accountName string) User {
	id := len(_users) + 1
	var acc Account
	if accountName != "" {
		acc = CreateAccount(accountName, 1000)
	}
	_users[id] = User{id: id, name: name, account: acc}
	return _users[id]
}

// ShowAllUsers display all of users
func ShowAllUsers() {
	fmt.Println(">>>>>>> ALL USERS >>>>>>>")
	for key, item := range _users {
		fmt.Printf("[%d] %+v\n", key, item)
	}
	fmt.Println(">>>>>>> ALL USERS >>>>>>>")
}

// GetUserByID return user data by id
func GetUserByID(id int) User {
	return _users[id]
}

func (u *User) Id() int {
	return u.id
}

func (u *User) Name() string {
	return u.name
}
func (u *User) Account() Account {
	return u.account
}
