<p align="center">
  <img src="/gopher.jpg" height="400">
  <h1 align="center">
    Software Design Patterns in Go
  </h1>
</p>

A list of Software Design Patterns discussed in the Exadel's Go Meetup. All sources are referenced according to _Harvard Referencing Style_ through [this website](https://www.citethisforme.com/). Bibliography citation is given in this file, whereas in-text citations are provided in each individual `.md` file in the sub-directories.

## Creational Patterns

| Pattern | Description |
|:-------:|:----------- |
| [Singleton](/creational/singleton) | Creates only a single instance without duplicates |
| [Builder](/creational/builder) | Helps to construct complex objects from simple objects |
| [Factory Method](/creational/factory) | Abstracts the user from the knowledge of the struct |
| [Abstract Factory](/creational/abstract_factory) | Is a _factory of factories_ that groups related families of objects |
| [Prototype](/creational/prototype) | Builds caches and default objects |


## Structural Patterns

| Pattern | Description |
|:-------:|:----------- |
| [Adapter](/structural/adapter) | Allows to use objects that weren't built for a specific purpose at the beginning |
| [Bridge](/structural/bridge) | Decouples object from what it does |
| [Composite](/structural/composite) | Creates hierarchies and trees of objects |
| [Decorator](/structural/decorator) | Decorates an already existing type with more functional features without actually touching it |
| [Facade](/structural/facade) | Hides the complexity scope from the user |
| [Flyweight](/structural/flyweight) | Reuses existing instances of objects with similar/identical state to minimize resource usage |
| [Proxy](/structural/proxy) | Wraps an object to hide some of its characteristics |

## Concurrency Patterns

| Pattern | Description |
|:-------:|:----------- |
| [Barrier](/concurrency/barrier) | Puts up a barrier so that
nobody passes until all the results are retrieved |
| [Future](/concurrency/future) | Implements asynchronous programming in Go |
| [Pipeline](/concurrency/pipeline) | Creates a concurrent structure of a multistep algorithm |
| [Workers Pool](/concurrency/workers_pool) | Bounds the amount of Goroutines available |
| [Publisher/Subscriber](/concurrency/publish_subscriber) | Concurrent-safe Observer Pattern |

## References

- Contreras, M., 2017. Go Design Patterns. 1st ed. Birmingham: Packt Publishing Ltd.