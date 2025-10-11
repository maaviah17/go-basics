package main

import (
	"fmt"
	"sync"
)

func randomGenerator(num int, w *sync.WaitGroup){
	defer w.Done()
	fmt.Println("number : ", num)
}

func main(){

	var wg sync.WaitGroup

	for i:=0;i<=5;i++{
		wg.Add(1)
		go randomGenerator(i,&wg)
	}

	wg.Wait()
}