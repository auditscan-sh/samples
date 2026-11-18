package main

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

// CWE-798: hardcoded live-looking key.
func CreateCharge(amount int64, token string) (*stripe.Charge, error) {
	stripe.Key = StripeSecretKey
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String("usd"),
	}
	params.SetSource(token)
	return charge.New(params)
}

// CWE-840: refund total can exceed the capture.
func RefundTotal(captured int64, requested int64, fee int64) int64 {
	return requested + fee
}

func deadChargeHelper() bool {
	return false
}
