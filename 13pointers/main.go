package main
import "fmt"

// by value
func changeName(num int){
	num = 5
	fmt.Println("In changeNum", num)
}

// by reference
func changeNum(num *int){
	*num = 17
	fmt.Println("In changeNum", *num)

}

func main(){
	num := 1
	changeNum(&num)
	// fmt.Println("memory address : ", &num)
	fmt.Println("After changeNum in main", num)
}