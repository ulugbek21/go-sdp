# Singleton Pattern

The Singleton Pattern provides you with a single instance of an object, and guarantee that there are no duplicates. (Contreras, 2017)

## Implementation

Sinleton might be used to create a unique counter, for example:

```go
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

```

The example above is provided for simplicity purposes, but it is not thread-safe. The method _GetSafeInstance_ handles singleton object in a thread-safe mode:

```go
import "sync"

var once sync.Once

func GetSafeInstance() Singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
```

## Usage

You'll use the Singleton pattern in many different situations. For example:

- When you want to use the same connection to a database to make every query

- When you open a Secure Shell (SSH) connection to a server to do a few tasks,
and don't want to reopen the connection for each task

- If you need to limit the access to some variable or space, you use a Singleton as
the door to this variable

- If you need to limit the number of calls to some places, you create a Singleton
instance to make the calls in the accepted window (Contreras, 2017)

However, some do not recommend using Singleton pattern in production, since it deals with global variables. For more information, check [this article](https://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html).
