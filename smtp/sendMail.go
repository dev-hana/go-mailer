package smtp

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"sync"
	"time"

	"github.com/dev-hana/go-mailer/database"
	"gopkg.in/gomail.v2"
)

type ExampleContent struct {
	AuthCode string `json:"AuthCode"`
}

func (smtp *SMTP) SendMail(mail *database.SendMail, result chan *database.SendMail, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	templateFile := fmt.Sprintf("%s.html", mail.TemplateTitle)

	var t *template.Template
	t = template.New(templateFile)
	t, err := t.ParseFiles(fmt.Sprintf("../assets/templates/%s", templateFile))
	if err != nil {
		log.Println(err.Error())
		mail.Status = false
		result <- mail
		return
	}

	exampleContent := ExampleContent{
		AuthCode: mail.Content,
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, exampleContent); err != nil {
		log.Println(err.Error())
		mail.Status = false
		result <- mail
		return
	}

	message := gomail.NewMessage()
	message.SetHeader("From", smtp.Username)
	message.SetAddressHeader("To", mail.ReceiverEmail, mail.ReceiverName)
	message.SetHeader("Subject", mail.Template.Subject)
	message.SetBody("text/html", tpl.String())
	if err := smtp.DialAndSend(message); err != nil {
		log.Println(err.Error())
		mail.Status = false
		result <- mail
		return
	}

	mail.Status = true
	fmt.Println("status", mail.Status)
	result <- mail

	defer func() {
		wg.Done()
		message.Reset()
	}()
}
