package main

import (
	"fmt"
	"myApp/auth"
	"myApp/user"
	"github.com/fatih/color"
)

func main(){
	auth.LoginWithCredentials("muawiyah17","hashedpassword")
	sesh := auth.GetSession()

	fmt.Println("session : ", sesh)

	user := user.User{
		Email : "mmk@mail.com",
		Name : "funky-monkey",
	}


	// fmt.Println("User info : ", user)
	color.Red(user.Email)
	color.Green(user.Name)

	// color.Blue("Prints %s in blue.", "text")
	// color.Red("We have red")
	// color.Yellow("Yellow color too!")
	color.Magenta("Hey,")

	// Hi-intensity colors
	color.HiGreen("HAPPY")
	color.Yellow("BIRTHDAY ")
	color.HiBlack("...")
	color.HiWhite("!!!")
}