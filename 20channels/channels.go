package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println(" --Channels in Go-- ")

	//
	myChannel := make(chan int,2)
	wg := &sync.WaitGroup{}

	// myChannel <- 5
	// fmt.Println(<-myChannel)

	wg.Add(2)


//RECIEVE ONLY
	go func(ch <-chan int, wg *sync.WaitGroup){

		//this is just a way to check if the channel is open or not || it gives a bool response.
		val, isChannelOpen := <-myChannel

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myChannel)
		// fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel,wg)

//SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup){

		myChannel <- 0
		close(myChannel)
		// myChannel <- 5
		// myChannel <- 6
		// close(myChannel)
		wg.Done()
	}(myChannel,wg)

	wg.Wait()

}