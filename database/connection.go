package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GORM struct {
	*gorm.DB
}

func ConnectDB(dbms string, dsn string) (*GORM, error) {
	var (
		db  *gorm.DB
		err error
	)

	switch dbms {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported DBMS: %s", dbms)
	}

	if err != nil {
		return nil, err
	}

	return &GORM{
		DB: db,
	}, nil
}
