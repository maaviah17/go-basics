package main

import (
	"fmt"
	"sync"
)

func main(){

	fmt.Println("-- Race Condition --")

	var score = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("1st Routine")
		mut.Lock()	
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("2nd Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("3rd Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
