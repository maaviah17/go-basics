package main
import "fmt"



type payment struct{
	gateway stripe
}

func (p payment) makePayment (amount float32){
	// razorpayGateway := razorpay{}
	// razorpayGateway.pay(amount)
	// stripePayment := stripe{}
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay (amount float32){
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}


type stripe struct{}

func (s stripe) pay (amount float32){
	fmt.Println("making payment using stripe : ", amount)
}

func main(){
	stripePayment := stripe{}
	r
	newPayment := payment{
		gateway : stripePayment,
	}
	newPayment.makePayment(1200)
}