package observer

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
)

type emailObserver struct {
	id            int
	senderEmail   string
	reciverEmails []string
	auth          smtp.Auth
}

var smtpHost = "smtp.gmail.com"
var smtpPort = "587"

func CreateEmailObservable(id int, email, password string, emails ...string) *emailObserver {
	// Create authentication
	auth := smtp.PlainAuth("", email, password, smtpHost)

	return &emailObserver{
		id:            id,
		senderEmail:   email,
		reciverEmails: emails,
		auth:          auth,
	}
}

func (e *emailObserver) Export() map[string]string {
	return map[string]string{
		"id":            fmt.Sprintf("%d", e.id),
		"senderEmail":   e.senderEmail,
		"reciverEmails": fmt.Sprintf("%v", e.reciverEmails),
		"typeKey":       e.getType(),
	}
}

func (e *emailObserver) AddEmail(s string) {
	e.reciverEmails = append(e.reciverEmails, s)
}
func (e *emailObserver) Update(wg *sync.WaitGroup, s string) {
	defer wg.Done()

	message := []byte(s)

	err := smtp.SendMail(smtpHost+":"+smtpPort, e.auth, e.senderEmail, e.reciverEmails, message)
	if err != nil {
		log.Fatal(err)
	}
}
func (e *emailObserver) GetID() int {
	return e.id
}
func (e *emailObserver) getType() string {
	return "2"
}
