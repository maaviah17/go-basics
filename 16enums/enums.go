package main
import "fmt"


type OrderStatus string

// type OrderStatus int
 
// const (
// 	Recieved OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

const (
	Recieved OrderStatus = "recieved"
	Confirmed = "confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("changing order status to ", status)
}
func main(){ 
	changeOrderStatus(Prepared)
}