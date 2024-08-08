package strategy

import "fmt"

// PaymentStrategy defines a common interface for different payment methods.
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CreditCardStrategy is a concrete strategy that implements payment using a credit card.
type CreditCardStrategy struct {
	CardNumber string
	CVV        string
	ExpiryDate string
}

// Pay method for CreditCardStrategy
func (cc *CreditCardStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card ending in %s", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

// PayPalStrategy is a concrete strategy that implements payment using PayPal.
type PayPalStrategy struct {
	Email string
}

// Pay method for PayPalStrategy
func (pp *PayPalStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal account %s", amount, pp.Email)
}

// ShoppingCart is the context class that will use a PaymentStrategy.
type ShoppingCart struct {
	Strategy PaymentStrategy
}

// SetPaymentStrategy allows for changing the PaymentStrategy at runtime.
func (cart *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	cart.Strategy = strategy
}

// Checkout uses the current PaymentStrategy to pay the given amount.
func (cart *ShoppingCart) Checkout(amount float64) string {
	return cart.Strategy.Pay(amount)
}
