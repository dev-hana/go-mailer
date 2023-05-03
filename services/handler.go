package services

import "github.com/dev-hana/go-mailer/database"

type Handler struct {
	db database.DBLayer
}

func NewHandler() (*Handler, error) {
	dsn := "host=localhost user=postgres password=1234 dbname=mailer port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := database.ConnectDB("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Handler{
		db: db,
	}, nil
}

func (h *Handler) InitTable() error {
	err := h.db.InitTable()
	return err
}
