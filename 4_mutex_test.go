package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceConditionWithMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(3 * time.Second)
	// Final value of x is 100,000. Race conditions are avoided using a mutex.
	fmt.Println("Final value of x:", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (ba *BankAccount) AddBalance(amount int) {
	ba.RWMutex.Lock()
	defer ba.RWMutex.Unlock()

	ba.Balance += amount
}

func (ba *BankAccount) GetBalance() int {
	ba.RWMutex.RLock()
	defer ba.RWMutex.RUnlock()

	return ba.Balance
}

func TestRWMutexBankAccount(t *testing.T) {
	account := BankAccount{}

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println("Balance:", account.GetBalance())
			}
			wg.Done()
		}()
	}

	wg.Wait()
	//time.Sleep(3 * time.Second)
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (userBalance *UserBalance) Lock() {
	userBalance.Mutex.Lock()
}

func (userBalance *UserBalance) Unlock() {
	userBalance.Mutex.Unlock()
}

func (userBalance *UserBalance) Change(amount int) {
	userBalance.Balance += amount
}

func Transfer(user1, user2 *UserBalance, amount int, wg *sync.WaitGroup) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
	wg.Done()
}

func TestDeadlock(t *testing.T) {
	wg := sync.WaitGroup{}
	budi := &UserBalance{
		Name:    "Budi",
		Balance: 1000000,
	}
	fian := &UserBalance{
		Name:    "Fian",
		Balance: 1000000,
	}

	wg.Add(2)
	go Transfer(budi, fian, 100000, &wg)
	go Transfer(fian, budi, 200000, &wg)

	wg.Wait()

	fmt.Println("Budi:", budi.Balance) // Expected: 1100000
	fmt.Println("Fian:", fian.Balance) // Expected: 900000
}
