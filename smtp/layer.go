package smtp

import (
	"sync"

	"github.com/dev-hana/go-mailer/database"
)

type SMTPLayer interface {
	// Check Connection
	CheckSMTPConnection() error
	SendMail(mail *database.SendMail, result chan *database.SendMail, wg *sync.WaitGroup)
}
