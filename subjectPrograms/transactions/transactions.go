package transactions

import (
	"errors"
	"fmt"
)

func Deposit(username string, amount int64) error{
	fmt.Println("Amount to be deposited:",amount)
	// Assertion
	if(amount>0 && amount<=10000){
		fmt.Println(username,"\nAmount Deposited:",amount)
	}else{
		fmt.Println(username)
		return errors.New("invalid amount")
	}
	
	return nil
}


// task 1. coverage over time.
// task 2: error/assertion injection, and report activated assertions over time for both gofuzz and "myfuzz"
// testing goal: generate inputs to exercise as many branches or assertions as possible.
// if (x is out of the range) assert("invalid amount of money")
// each assertion represents one posible error
// cumulative coverage over time
// cumulative # of assertions over time