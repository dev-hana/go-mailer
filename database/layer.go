package database

type DBLayer interface {
	// INIT
	InitTable() error

	//SMTP
	CreateSMTP(smtp *SMTP) error
}
