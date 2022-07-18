package main

import (
	observer "main/obserever"
	subject "main/subjects"
)

func main() {

	obs1 := observer.CreatePrint(1)
	obs2 := observer.CreatePrint(2)
	obs3 := observer.CreatePrint(3)
	obs4 := observer.CreatePrint(4)

	subj := subject.CreateBasic("subj 1")
	subj.Register(obs1)
	subj.Register(obs2)
	subj.Register(obs3)
	subj.Register(obs4)
	for {
		subj.Notify()
	}
}
