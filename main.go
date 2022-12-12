package main

import (
	myFuzz "GoFuzzTesting/gofuzz"
	"fmt"

	goFuzz "github.com/google/gofuzz"

	authentication "GoFuzzTesting/subjectPrograms"
)

func main() {
	createAccount()
	
		
}

func createAccount(){
	fmt.Println("-------------------------------------------------------------")
	goFuzzUsername,goFuzzPassword := goFuzzGenerator()
	fmt.Println("Username:",goFuzzUsername)
	fmt.Println("Password:",goFuzzPassword)
	_, _, goFuzzError := authentication.CreateAccount(goFuzzUsername,goFuzzPassword)
	fmt.Println("Error:",goFuzzError)
	fmt.Println("-------------------------------------------------------------")
	myFuzzUsername,myFuzzPassword := myFuzzGenerator()
	fmt.Println("Username:",myFuzzUsername)
	fmt.Println("Password:",myFuzzPassword)
	_, _, myFuzzError := authentication.CreateAccount(myFuzzUsername,myFuzzPassword)
	fmt.Println("Error:",myFuzzError)
	fmt.Println("-------------------------------------------------------------")
}

func goFuzzGenerator()(string,string){
	fmt.Println("GoFuzz")
	var username string
	unicodeRange := goFuzz.UnicodeRange{First: '0', Last: 'z'}
	stringFuzzer := goFuzz.New().Funcs(unicodeRange.CustomStringFuzzFunc())
	stringFuzzer.Fuzz(&username)
	var password string
	stringFuzzer.Fuzz(&password)
	return username,password

}

func myFuzzGenerator()(string,string){
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

