package subject

import (
	"sync"
)

type Observer interface {
	Update(wg *sync.WaitGroup, s string)
	GetID() int
	Export() map[string]string
}
type Subject interface {
	Register(o Observer) error
	Unregister(o Observer) error
	Notify()
	SaveObservers() error
}
type saveAgent interface {
	Save(obj map[int]Observer, fileName string) error
}
