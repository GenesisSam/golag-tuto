package main

import (
	"fmt"
	BankUtils "tuto-1/bankutils"
)

func main() {
	fmt.Printf("Hello, world \n")
	BankUtils.CreateUser("Tom", "")
	BankUtils.CreateUser("Jim", "")
	BankUtils.CreateUser("Jena", "")
	BankUtils.CreateUser("Jenny", "")
	BankUtils.CreateUser("BlackPink", "")
	var menuSelect int

	for {

		fmt.Println(">>> Bank Service tuto-1")
		fmt.Println("1. Create user")
		fmt.Println("2. Select user")
		fmt.Println("3. Show all user")
		fmt.Println("4. Exit")

		fmt.Print("> ")
		fmt.Scanf("%d\n", &menuSelect)

		switch menuSelect {
		case 1:
			{
				var name, accName string
				fmt.Println(">>>> CREATE NEW USER")
				fmt.Println(">> insert your name & account name")
				fmt.Print("(Name)> ")
				fmt.Scanf("%s\n", &name)
				fmt.Print("(Account Name)> ")
				fmt.Scanf("%s\n", &accName)

				newUser := BankUtils.CreateUser(name, accName)
				fmt.Printf(">> User created: %+v", newUser)

			}
		case 2:
			{
				var id int
				fmt.Println(">>>> SELECT USER")
				fmt.Println(">> insert user id")
				fmt.Print("> ")
				fmt.Scanf("%d", &id)
				user := BankUtils.GetUserByID(id)
				acc := user.Account()
				if !acc.IsEmpty() {
					fmt.Println(">> has Account")
				}
				fmt.Printf(">> Result: %+v", user)
			}

		case 3:
			{
				fmt.Println(">>>> SHOW ALL USERS")
				BankUtils.ShowAllUsers()
			}
		case 4:
			{
				goto END
			}
		}

		fmt.Println("")
	}
END:
}
