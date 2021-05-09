# Bridge Pattern

The Bridge pattern is a design with a slightly cryptic definition from the original _Gang of Four_ book. It decouples an abstraction from its implementation so that the two can vary independently. This cryptic explanation just means that you could even decouple the most basic form of functionality: decouple an object from what it does. (Contreras, 2017)

## Implementation

A specific printer might have one functionality, but after some time you might change your mind and change what it actually does. Instead of modifying the old code, you might need to just add new functionality and decouple that object.

Here it is achieved by having one interface _PrinterAPI_, two _PrinterImpl1_ and _PrinterImpl2_  struct types, which implement the interface:

```go
type PrinterAPI interface {
	PrintMessage(string) error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}

	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

```

And here is an example of decoupling the object and changing its functionality:

```go
normal := NormalPrinter{
    Msg:     expectedMessage,
    Printer: &PrinterImpl1{},
}

normal = NormalPrinter{
    Msg: expectedMessage,
    Printer: &PrinterImpl2{
        Writer: &testWriter,
    },
}
```

## Summary

Bridge Pattern allows you to reuse object's abstractions as well as its implementations.
