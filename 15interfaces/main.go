package main
import "fmt"

type paymenter interface{
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct{
	//provide interface in the struct not a concrete implementation
	gateway paymenter
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

type fakePayment struct{}

func (f fakePayment) pay(amount float32){
	fmt.Println("making payment using fake payment")
}

type stripe struct{}

func (s stripe) pay (amount float32){
	fmt.Println("making payment using stripe : ", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("making payment using paypal :", amount)
}

func (p paypal) refund(amount float32, account string){
	fmt.Println("refunding amount")
}

func main(){
	// stripePayment := stripe{}
	// razorPayment := razor{}
	// fakeGw := fakePayment{}
	paypalGw := paypal{}
	newPayment := payment{
		// gateway : stripePayment,
		// gateway : fakeGw,
		gateway : paypalGw,
	}
	newPayment.makePayment(1200)
}