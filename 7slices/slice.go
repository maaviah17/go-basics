package main
import "fmt"
import "slices"

func main(){

	nums := []int{1,3,5,2,8}
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// var slie = make([]int, 4)
	// fmt.Println(slie)

	nums = append(nums, 17)
	// nums = append(nums, 19)
	// fmt.Println(cap(nums))
	//fmt.Println(nums)

	new := []int{2,4}
	fmt.Println(len(new))
	fmt.Println(cap(new))

	new = append(new, 6)
	fmt.Println(new)
	fmt.Println(len(new))
	fmt.Println(cap(new))

	var numi = make([]int,0,5)
	numi = append(numi,7)

	var numi2 = make([]int,len(numi))

	// copying function
	copy(numi2,numi)
	fmt.Println(numi,numi2)
	
	// slice operation
	sum := []int{1,2,3,4,5,6}
	fmt.Println(sum[0:3])

	sums := []int{1,2,3,4,5,6}
	fmt.Println(slices.Equal(sum,sums))


	//2d slices
	dimension := [][]int{{1,2,3},{4,5,6}}
	fmt.Println(dimension)

}