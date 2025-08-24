package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	mutex := sync.Mutex{}

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("x:", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func TransferBalance(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	fmt.Println("Unlock", user1.Name)
	user2.Unlock()
	fmt.Println("Unlock", user2.Name)
}

func TestTransferBalance(t *testing.T) {
	user1 := UserBalance{
		Mutex:   sync.Mutex{},
		Name:    "Ucup 1",
		Balance: 10000,
	}

	user2 := UserBalance{
		Mutex:   sync.Mutex{},
		Name:    "Ucup 2",
		Balance: 10000,
	}

	go TransferBalance(&user1, &user2, 5000)
	go TransferBalance(&user1, &user2, 1000)
	time.Sleep(3 * time.Second)

	fmt.Println("user 1 balance:", user1.Balance)
	fmt.Println("user 2 balance:", user2.Balance)
}
