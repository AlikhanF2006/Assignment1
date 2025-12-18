package main

import (
	"fmt"

	"github.com/AlikhanF2006/Assignment1/bank"
	"github.com/AlikhanF2006/Assignment1/company"
	"github.com/AlikhanF2006/Assignment1/library"
	"github.com/AlikhanF2006/Assignment1/shapes"
)

func main() {
	lib := library.NewLibrary()

	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Borrow Book")
		fmt.Println("3. Return Book")
		fmt.Println("4. List Available Books")
		fmt.Println("5. Shapes Demo")
		fmt.Println("6. Company Demo")
		fmt.Println("7. Bank Demo")
		fmt.Println("8. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			var id, title, author string
			fmt.Print("Enter ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter Author: ")
			fmt.Scanln(&author)

			lib = lib.AddBook(id, title, author)

		case 2:
			var id string
			fmt.Print("Enter Book ID to borrow: ")
			fmt.Scanln(&id)
			lib = lib.BorrowBook(id)

		case 3:
			var id string
			fmt.Print("Enter Book ID to return: ")
			fmt.Scanln(&id)
			lib = lib.ReturnBook(id)

		case 4:
			lib.ListAvailableBooks()

		case 5:
			shapes.Demo()

		case 6:
			comp := company.NewCompany()

			comp = comp.AddEmployee(1, company.FullTimeEmployee{
				ID:     1,
				Name:   "Alice",
				Salary: 5000,
			})

			comp = comp.AddEmployee(2, company.PartTimeEmployee{
				ID:     2,
				Name:   "Bob",
				Salary: 2000,
			})

			comp.ListEmployees()

		case 7:
			account := bank.NewAccount()

			account = account.Deposit(1000)
			account = account.Withdraw(300)

			fmt.Println("Balance:", account.GetBalance())
			account.PrintTransactions()

		case 8:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
