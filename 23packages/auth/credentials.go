package auth

import "fmt"

//if want to export this function to other folder, then just name the function starting with capital
// func login(){...} -> private scope
// func Login(){...} -> global scope

func LoginWithCredentials(username string, password string){
	fmt.Println("logging user using ->  \n", "username : " , username, "\n password : ", password)
}

