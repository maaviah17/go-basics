package main
import "fmt"

func main(){

	// creating map
	m := make(map[string]string)

	// setting element inside map
	m["name"] = "xyz"
	m["attribute"] = "beautiful"

	// getting an element
	fmt.Println(m["name"], m["attribute"])


	newm := make(map[string]int)
	newm["age"] = 21
	newm["phone"] = 92323223

	fmt.Println(newm["age"], newm["phone"])
	fmt.Println(len(newm))

	delete(newm,"age")
	fmt.Println(newm)

	newm["height"] = 160
	fmt.Println(len(newm))

	clear(newm)
	fmt.Println(newm)

	//new map
	fmt.Println(".   .    NEW MAP  .   .")
	map1 := map[string]int{"price" : 200, "fax" : 23}
	k,ok := map1["price"]

	fmt.Println(k)
	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}


}