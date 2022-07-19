package main

import (
	"log"
	observer "main/obserever"
	saveAgent "main/storage"
	subject "main/subjects"

	"github.com/spf13/viper"
)

func testPrint() {
	obs1 := observer.CreatePrintObservable(1)
	obs2 := observer.CreatePrintObservable(2)
	obs3 := observer.CreatePrintObservable(3)
	obs4 := observer.CreatePrintObservable(4)

	subj := subject.CreateBasicSubject("subj 1")
	subj.Register(obs1)
	subj.Register(obs2)
	subj.Register(obs3)
	subj.Register(obs4)
	for {
		subj.Notify()
	}

}
func testEmail() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
		return
	}
	email, ok := viper.Get("email").(string)
	if !ok {
		log.Fatalf("Email not set")
		return
	}
	password, ok := viper.Get("password").(string)

	if !ok {
		log.Fatalf("password not set")
		return
	}
	obs1 := observer.CreateEmailObservable(1, email, password)
	subj := subject.CreateBasicSubject("subj 1")
	subj.Register(obs1)
	subj.Notify()

}
func testSave() {
	subj := subject.CreateBasicSubject("subj 1")

	obs1 := observer.CreatePrintObservable(1)
	obs2 := observer.CreateEmailObservable(2, "fakemail.com", "fakepassword", "email@test", "email2@test")

	subj.Register(obs1)
	subj.Register(obs2)

	a := saveAgent.CreateJsonSaveAgent()
	subj.SaveObservers(a, "observers")
}
func main() {

	testSave()

}
