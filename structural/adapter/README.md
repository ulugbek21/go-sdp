# Adapter Pattern

One of the most commonly used structural patterns is the Adapter pattern. Like in real life, where you have plug adapters and bolt adapters, in Go, an adapter will allow us to use something that wasn't built for a specific task at the beginning. (Contreras, 2017)

## Implementation

Imagine that we have old printer and new printer with newer functionalities. We need to have a different method for the new printer and we also should be able to use the old printer's method if necessary. Here _adapter_ comes in handy. _PrinterAdapter_ is a struct that can store OldPrinter

```go
type LegacyPrinter interface {
	Print(s string) string
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}
```

_PrinterAdapter_ implements modern printer and acts as an old printer if necessary:

```go
func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return
}
```

## Summary

With the Adapter design pattern, it is possible to achieve the open/close
principle in your applications. Instead of modifying your old source code (something which could not be possible in some situations), you can create a way to use the old
functionality with a new signature. (Contreras, 2017)
