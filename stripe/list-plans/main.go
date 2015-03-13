package main

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/plan"
	"os"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_PRIVATE_KEY")

	i := plan.List(nil)
	for i.Next() {
		fmt.Println(i.Plan())
	}
}
