package strategy

import "testing"

func TestCreditCardStrategy_Pay(t *testing.T) {
	amount := 100.0
	creditCard := &CreditCardStrategy{
		CardNumber: "1234567890123456",
		CVV:        "123",
		ExpiryDate: "06/26",
	}
	want := "Paid $100.00 using Credit Card ending in 3456"
	got := creditCard.Pay(amount)
	if got != want {
		t.Errorf("CreditCardStrategy.Pay() = %v, want %v", got, want)
	}
}

func TestPayPalStrategy_Pay(t *testing.T) {
	amount := 50.0
	payPal := &PayPalStrategy{
		Email: "user@example.com",
	}
	want := "Paid $50.00 using PayPal account user@example.com"
	got := payPal.Pay(amount)
	if got != want {
		t.Errorf("PayPalStrategy.Pay() = %v, want %v", got, want)
	}
}

func TestShoppingCart_Checkout(t *testing.T) {
	amount := 150.0
	shoppingCart := ShoppingCart{}

	// Test CreditCardStrategy
	creditCard := &CreditCardStrategy{
		CardNumber: "1234567890123456",
		CVV:        "123",
		ExpiryDate: "06/26",
	}
	shoppingCart.SetPaymentStrategy(creditCard)
	wantCreditCard := "Paid $150.00 using Credit Card ending in 3456"
	gotCreditCard := shoppingCart.Checkout(amount)
	if gotCreditCard != wantCreditCard {
		t.Errorf("ShoppingCart with CreditCardStrategy.Checkout() = %v, want %v", gotCreditCard, wantCreditCard)
	}

	// Test PayPalStrategy
	payPal := &PayPalStrategy{
		Email: "user@example.com",
	}
	shoppingCart.SetPaymentStrategy(payPal)
	wantPayPal := "Paid $150.00 using PayPal account user@example.com"
	gotPayPal := shoppingCart.Checkout(amount)
	if gotPayPal != wantPayPal {
		t.Errorf("ShoppingCart with PayPalStrategy.Checkout() = %v, want %v", gotPayPal, wantPayPal)
	}
}
