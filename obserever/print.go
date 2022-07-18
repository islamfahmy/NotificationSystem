package observer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type print struct {
	id int
}

func CreatePrint(id int) *print {
	return &print{
		id: id,
	}
}

func (p *print) Update(wg *sync.WaitGroup, s string) {
	defer wg.Done()
	t := time.Duration(rand.Intn(5))
	time.Sleep(time.Second * t)
	fmt.Println("observer:", p.id, "subject: ", s)
}
func (p *print) GetID() int {
	return p.id
}
