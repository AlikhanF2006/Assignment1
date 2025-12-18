package bank

import "fmt"

type BankAccount struct {
	Balance      float64
	Transactions []float64
}

func NewAccount() BankAccount {
	return BankAccount{
		Balance:      0,
		Transactions: []float64{},
	}
}

func (b BankAccount) Deposit(amount float64) BankAccount {
	if amount <= 0 {
		fmt.Println("Invalid deposit amount")
		return b
	}

	b.Balance += amount
	b.Transactions = append(b.Transactions, amount)
	fmt.Println("Deposited:", amount)
	return b
}

func (b BankAccount) Withdraw(amount float64) BankAccount {
	if amount <= 0 {
		fmt.Println("Invalid withdrawal amount")
		return b
	}

	if amount > b.Balance {
		fmt.Println("Insufficient funds")
		return b
	}

	b.Balance -= amount
	b.Transactions = append(b.Transactions, -amount)
	fmt.Println("Withdrawn:", amount)
	return b
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

func (b BankAccount) PrintTransactions() {
	fmt.Println("Transactions:")
	for _, t := range b.Transactions {
		fmt.Println(t)
	}
}
