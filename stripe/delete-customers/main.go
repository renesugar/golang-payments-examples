package main

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"os"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_PRIVATE_KEY")

	params := stripe.CustomerListParams{}
	i := customer.List(&params)
	for i.Next() {
		c := i.Customer()
		customer.Del(c.ID)
	}
}
