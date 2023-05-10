package services

import (
	"fmt"
	"sync"

	"github.com/dev-hana/go-mailer/database"
)

func (h *Handler) SendMailScheduler() {
	// get send mail info
	mails, err := h.db.GetSendMail()
	if err != nil {
		return
	}
	fmt.Println("Start for len: ", len(mails))

	var wg sync.WaitGroup
	wg.Add(len(mails) * 2)
	results := make(chan *database.SendMail)

	// send mail
	for _, mail := range mails {
		go h.smtp.SendMail(mail, results, &wg)
	}

	// update status
	for result := range results {
		go h.UpdateEmailStatus(result, &wg)
	}
	wg.Wait()
}

func (h *Handler) UpdateEmailStatus(mail *database.SendMail, wg *sync.WaitGroup) {
	if mail.Status {
		if err := h.db.UpdateStatus(mail); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	defer wg.Done()
}
