package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //wait groups are ususually pointers but here just for example
var mut sync.Mutex //same...these are created using pointers

func main(){ 
	websiteList := []string {
		"https://github.com",
		"https://fb.com",
		"https://go.dev",
	}

	for _,web := range websiteList{
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}



func getStatusCode(endpoint string){
	defer wg.Done()
	res,err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPsie in endpoint")
	}else{
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
	
}
