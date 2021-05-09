# Prototype Pattern

The Prototype Pattern uses an already created instance of some type to clone it and complete it with the particular needs of each context (Contreras, 2017)

## Implementation

Imagine that we have shirts shop. In the shop, we may have white shorts, black shorts and shirts of different colors, and each might have sub-types with some common characteristics. Instead of creating instances from zero, we might create prototypes to reuse them when creating new objects.

Here are the prototypes:

```go
var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

```

This prototypes will be reused for each individual instance of object and specific fields will be changed if necessary:

```go
    shirt1.SKU = "abbcc"
    shirt2.SKU = "fsbvf"
```

## Summary

The Prototype pattern is a powerful tool to build caches and default objects (Contreras, 2017)
