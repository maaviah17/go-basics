package main

import (
	"fmt"
	// "time"
)

func task(id int){
	fmt.Println("doing task : ", id)
}

func main(){

	
	for i:=0 ; i<=10; i++{
		go task(i)

		//this is a closure
		// go func(i int){
		// 	fmt.Println("items: ", i)
		// }(i)
	}

	// time.Sleep(time.Second*1) <- why used ? present in .md 
}