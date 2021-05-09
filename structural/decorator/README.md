# Bridge Pattern

The Decorator design pattern allows you to decorate an already existing type with more functional features without actually touching it. How is it possible? Well, it uses an approach similar to _matryoshka dolls_, where you have a small doll that you can put inside a doll of the same shape but bigger, and so on and so forth.

The Decorator type implements the same interface of the type it decorates, and stores an instance of that type in its members. This way, you can stack as many decorators (dolls) as you want by simply storing the old decorator in a field of the new one. (Contreras, 2017)

## Implementation

In our example, we will prepare a Pizza type, where the core is the pizza and the
ingredients are the decorating types. We will have a couple of ingredients for our
pizza–onion and meat.

```go
type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type Meat struct {
	Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Onion")
	}

	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", s, "onion"), nil
}
```

As a result, we could _simulate matryoshka dolls_:

```go
pizza := &Onion{&Meat{&PizzaDecorator{}}}
```
