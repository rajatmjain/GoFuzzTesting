package transactions

import (
	"errors"
	"fmt"
)

func Deposit(username string, amount int64) error{
	fmt.Println("Amount to be deposited:",amount)
	if(amount>0 && amount<=10000){
		fmt.Println(username,"\nAmount Deposited:",amount)
	}else{
		fmt.Println(username)
		return errors.New("invalid amount")
	}
	
	return nil
}