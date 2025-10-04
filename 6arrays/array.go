package main
import "fmt"

func main(){

	var nums [4]int  
	nums[2] = 12
	
	var value[3]bool
	value[1] = true
	fmt.Println(value)

	fmt.Println(nums)
	// arrays length
	fmt.Println(len(nums))


	var name[3]string
	name[1] = "dj"
	fmt.Println(name)

	num := [3]int{1,5,24}
	fmt.Println(num)

	numdimension := [2][2]int{}
}