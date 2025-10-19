package main

import (
	"fmt"
	"sync"
)

func main(){

	fmt.Println("Channels in Go : ")
	numChan := make(chan int,2)
	var wg sync.WaitGroup

	//adding 2 go routines
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup){
		fmt.Println(<-numChan)
		wg.Done()
	}(numChan,&wg)


	go func(ch chan int, wg *sync.WaitGroup){
		numChan <- 17
		wg.Done()
	}(numChan,&wg)

	wg.Wait()

}