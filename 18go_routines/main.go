package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(id int, w *sync.WaitGroup){
	defer w.Done()
	fmt.Println("doing task : ", id)
}

func main(){

	var wg sync.WaitGroup 
	
	for i:=0 ; i<=10; i++{
		wg.Add(1)
		go task(i,&wg)
		//this is a closure
		// go func(i int){
		// 	fmt.Println("items: ", i)
		// }(i)
	}

	wg.Wait()

	// time.Sleep(time.Second*1) <- why used ? present in .md 
}