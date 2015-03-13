package main

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/plan"
	"os"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_PRIVATE_KEY")

	plan.New(&stripe.PlanParams{
		ID:       "Free",
		Amount:   0,
		Interval: "month",
		Name:     "Free Plan",
		Currency: "usd",
	})

	plan.New(&stripe.PlanParams{
		ID:        "Basic Monthly",
		Amount:    1999,
		Interval:  "month",
		Name:      "Basic Monthly Plan",
		Currency:  "usd",
		Statement: "Basic Monthly Plan",
	})

	plan.New(&stripe.PlanParams{
		ID:        "Basic Annual",
		Amount:    19999,
		Interval:  "year",
		Name:      "Basic Annual Plan",
		Currency:  "usd",
		Statement: "Basic Annual Plan",
	})

	plan.New(&stripe.PlanParams{
		ID:        "Premium Monthly",
		Amount:    4999,
		Interval:  "month",
		Name:      "Premium Monthly Plan",
		Currency:  "usd",
		Statement: "Premium Monthly Plan",
	})

	plan.New(&stripe.PlanParams{
		ID:        "Premium Annual",
		Amount:    49999,
		Interval:  "year",
		Name:      "Premium Annual Plan",
		Currency:  "usd",
		Statement: "Premium Annual Plan",
	})
}
