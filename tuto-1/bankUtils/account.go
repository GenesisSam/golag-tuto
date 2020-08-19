package bankutils

import (
	"fmt"
	"math"
)

// Account bank account
type Account struct {
	accNo   string
	name    string
	balance int
}

var _accounts = make(map[string]Account)

// CreateAccount generate new account
func CreateAccount(name string, seedMoney int) Account {
	id := fmt.Sprintf("%06d", len(_accounts)+1)
	_accounts[id] = Account{accNo: id, name: name, balance: seedMoney}

	return _accounts[id]
}

// ShowBalance 현재 남은 잔고를 보여줍니다.
func (acc *Account) ShowBalance() int {
	return acc.balance
}

// Deposit 입금
func (acc *Account) Deposit(money int) int {
	return acc.balance + int(math.Abs(float64(money)))
}

// Withdraw 출금
func (acc *Account) Withdraw(money int) int {
	return acc.balance - int(math.Abs(float64(money)))
}

func (acc *Account) IsEmpty() bool {
	return acc.name == "" && acc.accNo == ""
}
