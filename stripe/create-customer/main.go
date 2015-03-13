package main

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"os"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_PRIVATE_KEY")

	customer.New(&stripe.CustomerParams{
		Desc:  "Example Customer",
		Email: "john.doe@example.com",
	})

	customer.New(&stripe.CustomerParams{
		Desc:  "Example Customer 2",
		Email: "jane.doe@example.com",
	})
}
