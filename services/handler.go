package services

import (
	"net/http"

	"github.com/dev-hana/go-mailer/database"
	"github.com/dev-hana/go-mailer/smtp"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db   database.DBLayer
	smtp smtp.SMTPLayer
}

func NewHandler(dbms, dsn string) (*Handler, error) {
	db, err := database.ConnectDB(dbms, dsn)
	if err != nil {
		return nil, err
	}

	smtp, err := smtp.ConnectSMTP()
	if err != nil {
		return nil, err
	}

	return &Handler{
		db:   db,
		smtp: smtp,
	}, nil
}

func (h *Handler) InitTable() error {
	err := h.db.InitTable()
	return err
}

func (h *Handler) CheckServerConnection(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server Database error"})
		return
	}

	if h.smtp == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server SMTP error"})
		return
	}
}

func ResponseBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusAccepted, gin.H{"httpCode": http.StatusBadRequest, "error": err.Error()})
	return
}
