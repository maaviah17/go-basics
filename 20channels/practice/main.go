package main

import (
	"fmt"
	// "math/rand"
	// "sync"
	// "time"
)


//for sending
// func processNum(ch chan int, wg *sync.WaitGroup){
// 	defer wg.Done()

// 	for num := range ch{
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Millisecond*3000)
// 	}
// }

//for recieving
func sum(res chan int, num1 int, num2 int){
	nums := num1 + num2
	res <- nums
}

func main(){
	
	result := make(chan int)
	go sum(result,33,36)

	res := <-result
	fmt.Println(res)
	// numChan := make(chan int)
	// wg := &sync.WaitGroup{}


	// wg.Add(1)
	// go processNum(numChan, wg)

	// for {
	// 	numChan <- rand.Intn(10)
	// }

	// wg.Wait()

	// messageChan := make(chan string,2)
	// wg := &sync.WaitGroup{}
	// messageChan <- "ping"
	// msg := <-messageChan
	// wg.Add(1)
	// go func(wg *sync.WaitGroup){
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(wg)
	// wg.Wait()
}