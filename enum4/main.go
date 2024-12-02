package main

import "fmt"

type PaymentStatus int

const (
	Pending PaymentStatus = iota
	Completed
	Failed
	Refunded
)

func (s PaymentStatus) String() string {
	switch s {
	case Pending:
		return "Pending :-) "
	case Completed:
		return "Completed"
	case Failed:
		return "Failed"
	case Refunded:
		return "Refunded"
	default:
		return "Unknown"
	}
}
func processPayment(status PaymentStatus) {
	switch status {
	case Pending:
		fmt.Println("Processing payment...")
	case Completed:
		fmt.Println("Payment completed successfully.")
	case Failed:
		fmt.Println("Payment failed. Please try again.")
	case Refunded:
		fmt.Println("Payment refunded.")
	default:
		fmt.Println("Unknown payment status.")
	}

}

func main() {

	status := Pending
	fmt.Println("Payment Status..:", status)

	// processPayment(Pending)
	processPayment(status)

	status = Completed
	fmt.Println("Payment Status..:", status)
	processPayment(status)
}
