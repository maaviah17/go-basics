package main
import "fmt"

// func printSlice(items []int){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

func printSlice[T any](items []T){
	for _,item := range items{
		fmt.Println(item)
	}
}

func printStringSlice(items []string){
	for _,item := range items{
		fmt.Println(item)
	}
}

func main(){
	// nums := []int{1,17,21}
	names := []string{"golang", "javascript", "cpp"}
	printSlice(names)
	// printStringSlice(names)

}