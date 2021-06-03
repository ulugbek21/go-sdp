package main

import "fmt"

var overdraftLimit = -500

// BankAccount ...
type BankAccount struct {
  balance int
}

// Deposit ...
func (b *BankAccount) Deposit(amount int) {
  b.balance += amount
  fmt.Println("Deposited", amount,
    "\b, balance is now", b.balance)
}

// Withdraw ...
func (b *BankAccount) Withdraw(amount int) bool {
  if b.balance - amount >= overdraftLimit {
    b.balance -= amount
    fmt.Println("Withdrew", amount,
      "\b, balance is now", b.balance)
    return true
  }
  return false
}

// Command ...
type Command interface {
  Call()
  Undo()
}

// Action ...
type Action int
const (
  // Deposit ...
  Deposit Action = iota
  // Withdraw ...
  Withdraw
)

// BankAccountCommand ...
type BankAccountCommand struct {
  account *BankAccount
  action Action
  amount int
  succeeded bool
}

// Call ...
func (b *BankAccountCommand) Call() {
  switch b.action {
  case Deposit:
    b.account.Deposit(b.amount)
    b.succeeded = true
  case Withdraw:
    b.succeeded = b.account.Withdraw(b.amount)
  }
}

// Undo ...
func (b *BankAccountCommand) Undo() {
  if !b.succeeded { return }
  switch b.action {
  case Deposit:
    b.account.Withdraw(b.amount)
  case Withdraw:
    b.account.Deposit(b.amount)
  }
}

// NewBankAccountCommand ...
func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
  return &BankAccountCommand{account: account, action: action, amount: amount}
}

func main() {
  ba := BankAccount{}
  cmd := NewBankAccountCommand(&ba, Deposit, 100)
  cmd.Call()
  cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
  cmd2.Call()
  fmt.Println(ba)
  cmd2.Undo()
  fmt.Println(ba)
}