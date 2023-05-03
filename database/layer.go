package database

type DBLayer interface {
	// INIT
	InitTable() error
}
