package singleton

import "sync"

var once sync.Once

// GetSafeInstance is similar to GetInstance but safe to threads
func GetSafeInstance() Singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
