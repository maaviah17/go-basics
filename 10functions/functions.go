package main
import "fmt"

func add(a int,b int) int{
	return a+b
}

func getNames () (string,string, int){
	return "muawiyah","khalid", 32
}

func processIt(fn func(a int)int){
	fn(17)
}

func main(){
	sum := add(4,5)
	fmt.Println(sum)

	// _ used to ignore the value
	fname,lname,_ := getNames()
	_,_,age := getNames();
	// fmt.Println(getNames())
	fmt.Println(lname, fname)
	fmt.Println(age)

	processIt()
}