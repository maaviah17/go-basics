package main
import "fmt"

func main(){

	//while loop
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i) 
	// 	i++
	// }

	// classic for loop
	// for i:=0 ; i<=3; i++ {
	// 	fmt.Println(i)
	// }

	// printing table of 2 
	for i:=0; i<=20; i++{

		if i%2==0 {
			// continue
			fmt.Println(i)
		}
	} 

	age := 17
	if age>18 {
		fmt.Println("can vote")
	}else{
		fmt.Println("cannot vote")
	}
}