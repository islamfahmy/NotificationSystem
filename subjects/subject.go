package subject

import "sync"

type Observer interface {
	Update(wg *sync.WaitGroup, s string)
	GetID() int
}
type Subject interface {
	Register(o Observer) error
	Unregister(o Observer) error
	Notify()
}
