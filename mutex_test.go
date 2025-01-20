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

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println("Balance:", account.GetBalance())
			}
		}()
	}

	time.Sleep(3 * time.Second)
}
