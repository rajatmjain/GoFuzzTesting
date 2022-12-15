package authentication

import (
	"errors"
	"fmt"
	"strings"
)

func CreateAccount(username string, password string) (string,string,error){
	// check username validity
	if(len(username)<=2 || len(username)>=16){
		return username,password,errors.New("invalid username length. Length should be greater than 2 and less than 16")
	}
	if(len(password)<6){
		return username,password,errors.New("invalid password length. Length should be greater than 6")
	}
	if(strings.ContainsAny(username,`"!"#$%&'()*+,-./"):<=>?@;`)){
		return username,password,errors.New("invalid characters in username")
	}

	if(!strings.ContainsAny(password,`"!"#$%&'()*+,-./"):<=>?@0123456789;`)){
		return username,password,errors.New("missing compulsory characters in password")
	}

	if(username==password){
		return username,password,errors.New("username and password can not be the same")
	}

	// Main logic
	fmt.Println("Account created with initial balance of $100.")
	return username,password,nil
}