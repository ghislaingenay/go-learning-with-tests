package syncgo

import "sync"

type Counter struct {
	mu sync.Mutex // A Mutex is a mutual exclusion lock. 
	// The zero value for a Mutex is an unlocked mutex.
	value int
}

func (c *Counter) Inc() {
	// any goroutine calling Inc will acquire the lock on Counter
	c.mu.Lock()
	defer c.mu.Unlock()
	//  All the other goroutines will have to wait for it to be Unlocked before getting access.
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func NewCounter() *Counter {
	return &Counter{}
}