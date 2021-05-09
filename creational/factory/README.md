# Factory Pattern

The Factory method pattern (or simply, Factory) is probably the second-best known and used design pattern in the industry. Its purpose is to abstract the user from the knowledge of the struct he needs to achieve for a specific purpose, such as retrieving some value, maybe from a web service or a database. The user only needs an interface that provides him this value. By delegating this decision to a Factory, this Factory can provide an interface that fits the user needs. It also eases the process of downgrading or upgrading of the implementation of the underlying type if needed. (Contreras, 2017)

## Implementation

The example could be a _PaymentMethod_ factory with two types of payment methods - _debit card_ and _cash_

```go
type PaymentMethod interface {
	Pay(amount float32) string
}

type CashPM struct{}

type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}

```

Indeed, methods for cash and debit card will be much more complex and different.

## Summary

Factory method pattern allows to group families of objects so that their implementation is outside of our scope. (Contreras, 2017)
