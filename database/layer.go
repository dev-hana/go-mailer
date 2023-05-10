package database

type DBLayer interface {
	// INIT
	InitTable() error

	CreateMail(mail SendMail) error
	GetSendMail() (mails []*SendMail, err error)
	UpdateStatus(mail *SendMail) error
}
