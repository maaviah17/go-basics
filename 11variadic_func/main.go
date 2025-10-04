package main
import "fmt"

func sum(nums ...int) int {
	total := 0
	for _,num := range nums{
		total += num
	}
	
	return total
}

func main(){

	nums := []int{3,66,3}
	result := sum(nums...)
	fmt.Println(result)

}