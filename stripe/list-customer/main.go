package main

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"os"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_PRIVATE_KEY")

	// list all customers
	i := customer.List(nil)
	for i.Next() {
		c := i.Customer()
		fmt.Println(c)
	}

}
