package singleton

// Singleton holds method AddOne to be implemented by singleton object
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

// GetInstance gets a single instance of an object
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// AddOne ...
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
