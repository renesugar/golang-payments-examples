package main

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/sub"
	"os"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_PRIVATE_KEY")

	customers := make(map[string]*stripe.Customer)
	i := customer.List(nil)
	for i.Next() {
		c := i.Customer()
		customers[c.Email] = c
	}

	if john, ok := customers["john.doe@example.com"]; ok {
		s := john.Subs.Values[0]
		sub.Cancel(s.ID, &stripe.SubParams{
			Customer: john.ID,
		})
	}

}
