package factory

import (
	"fmt"
)

// PaymentMethod ...
type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	// Cash ...
	Cash = iota
	// DebitCard ...
	DebitCard
)

// GetPaymentMethod ...
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)
	}
}

// CashPM ...
type CashPM struct{}

// DebitCardPM ...
type DebitCardPM struct{}

// Pay ...
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

// Pay ...
func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}
