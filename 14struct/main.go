package main

import "fmt"
import "time"

// e-commerce order struct
type order struct{
	id string
	amount float32
	status string
	createdAt time.Time //nanosecond precision
} 

//reciever type
func (o order) changeStatus(status string){
	o.status = status
}

func (o *order) getAmount() float32{
	return o.amount
}

func main(){

	myOrder := order{
		id : "1",
		amount : 50.00,
		status : "recieved",
	}

	myOrder.status = "confirmed"

	myOrder.createdAt = time.Now()
	fmt.Println(myOrder.getAmount())
	fmt.Println("order 1 : ", myOrder)


	orderNumber2 := order {
		id : "2",
		amount : 45.50,
		status : "delievered",
		createdAt : time.Now(),
	}

	fmt.Println("order number 2 : ", orderNumber2)

}