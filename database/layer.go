package database

type DBLayer interface {
	// INIT
	InitTable() error

	GetSendMail() (mails []*SendMail, err error)
}
