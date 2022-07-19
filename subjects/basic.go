package subject

import (
	"errors"
	"fmt"
	"sync"
)

type basic struct {
	observers map[int]Observer
	name      string // name of the subject
}

func CreateBasicSubject(name string) *basic {
	return &basic{
		observers: make(map[int]Observer),
		name:      name,
	}
}
func (b *basic) Register(o Observer) error {
	if _, ok := b.observers[o.GetID()]; ok {
		return errors.New("Observer already registered")
	}
	b.observers[o.GetID()] = o // add the observer to the list of observers
	return nil
}

func (b *basic) Unregister(o Observer) error {
	if _, ok := b.observers[o.GetID()]; !ok {
		return errors.New("no such observer")
	}
	b.observers[o.GetID()] = o // add the observer to the list of observers
	delete(b.observers, o.GetID())
	return nil

}
func (b *basic) Notify() {
	var wg sync.WaitGroup
	wg.Add(len(b.observers))
	defer fmt.Println("Finished notifying")
	defer wg.Wait()
	for _, o := range b.observers {
		go o.Update(&wg, b.name)
	}
	fmt.Println("waiting for all observers to finish")
}

func (b *basic) SaveObservers(s saveAgent, fileName string) error { //save the observers to a file

	err := s.Save(b.observers, fileName)
	if err != nil {
		return err
	}
	return nil
}
