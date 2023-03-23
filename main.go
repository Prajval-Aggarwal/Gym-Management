package main

import (
	"gym-api/Routes"
	"github.com/stripe/stripe-go"
)
func main() {
	stripe.Key = "sk_test_51MnxVTSGT1jvrl9CIDO2h1vvRKS0yKYBu0MRagvAcLn9ZshNY7P5CpLLamz6U7rUhx4Bch0Onv03vsoYfg9Bitpv006VIbV229"
	routes.Routes()
}
