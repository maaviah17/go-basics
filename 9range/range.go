package main
import "fmt"


func main(){
	nums := []int{1,3,7}

	for i:=0; i<len(nums); i++ {
		fmt.Println(nums[i])
	}
	
	sum:=0
	for _,num := range nums {
		sum += num
	} 
	fmt.Println(sum)

	for i,num := range nums {
		fmt.Println(i,num)
	}

	m:=map[string]string{"fname":"john", "lname":"doe"}
	
	for k,v := range m{
		fmt.Println(k,v)
	}

	for i,c := range "mmk" {
		fmt.Println(i,c)
	}
}