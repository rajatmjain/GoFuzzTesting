package main

import (
	myFuzz "GoFuzzTesting/gofuzz"
	"fmt"

	goFuzz "github.com/google/gofuzz"

	authentication "GoFuzzTesting/subjectPrograms/authentication"
	"GoFuzzTesting/subjectPrograms/transactions"
)

func main() {
	goFuzzUsername,myFuzzUsername,goFuzzError,myFuzzError := createAccount()
	if(goFuzzError==nil){
		fmt.Println("Transaction error:",goFuzzTransactions(goFuzzUsername))
		fmt.Println("-------------------------------------------------------------")
	}
	if(myFuzzError==nil){
		fmt.Println("Transaction error:",myFuzzTransactions(myFuzzUsername))
		fmt.Println("-------------------------------------------------------------")
	}
	

		
}

////////////////////// CALLING FUNCTIONS //////////////////////
func createAccount()(string,string,error,error){
	fmt.Println("-------------------------------------------------------------")
	goFuzzUsername,goFuzzPassword := goFuzzUsernamePasswordGenerator()
	fmt.Println("Username:",goFuzzUsername)
	fmt.Println("Password:",goFuzzPassword)
	_, _, goFuzzError := authentication.CreateAccount(goFuzzUsername,goFuzzPassword)
	fmt.Println("Error:",goFuzzError)
	fmt.Println("-------------------------------------------------------------")
	myFuzzUsername,myFuzzPassword := myFuzzUsernamePasswordGenerator()
	fmt.Println("Username:",myFuzzUsername)
	fmt.Println("Password:",myFuzzPassword)
	_, _, myFuzzError := authentication.CreateAccount(myFuzzUsername,myFuzzPassword)
	fmt.Println("Error:",myFuzzError)
	fmt.Println("-------------------------------------------------------------")
	return goFuzzUsername,myFuzzUsername,goFuzzError,myFuzzError
}

func goFuzzTransactions(goFuzzUsername string)error{
	fmt.Println("-------------------------------------------------------------")
	amount := goFuzzAmountGenerator()
	return transactions.Deposit(goFuzzUsername,amount)
}

func myFuzzTransactions(myFuzzUsername string)error{
	fmt.Println("-------------------------------------------------------------")
	amount := myFuzzAmountGenerator()
	return transactions.Deposit(myFuzzUsername,amount)
}

////////////////////// GENERATORS //////////////////////
func goFuzzUsernamePasswordGenerator()(string,string){
	fmt.Println("GoFuzz")
	var username string
	unicodeRange := goFuzz.UnicodeRange{First: '0', Last: 'z'}
	stringFuzzer := goFuzz.New().Funcs(unicodeRange.CustomStringFuzzFunc())
	stringFuzzer.Fuzz(&username)
	var password string
	stringFuzzer.Fuzz(&password)
	return username,password

}

func myFuzzUsernamePasswordGenerator()(string,string){
	fmt.Println("MyFuzz")
	var username string
	usernameUnicodeRange := myFuzz.UnicodeRange{First: '0', Last: 'z',MinLength: 2,MaxLength: 16}
	usernameFuzzer := myFuzz.New().Funcs(usernameUnicodeRange.CustomStringFuzzFunc())
	usernameFuzzer.Fuzz(&username)
	var password string
	passwordUnicodeRange := myFuzz.UnicodeRange{First: '0', Last: 'z',MinLength: 7,MaxLength: 20}
	passwordFuzzer := myFuzz.New().Funcs(passwordUnicodeRange.CustomStringFuzzFunc())
	passwordFuzzer.Fuzz(&password)
	return username,password
}

func goFuzzAmountGenerator() int64{
	var B int64
	int64Fuzzer := goFuzz.New()
	int64Fuzzer.Fuzz(&B)
	return B
}

func myFuzzAmountGenerator() int64{
	var B int64
	int64Schema := myFuzz.Int64Schema{Minimum: 0,Maximum: 10000}
	int64Fuzzer := myFuzz.New().Funcs(int64Schema.CustomInt64FuzzFunc())
	int64Fuzzer.Fuzz(&B)
	return B
}

