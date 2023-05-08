package services

import (
	"net/http"

	"github.com/dev-hana/go-mailer/database"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db database.DBLayer
}

func NewHandler(dbms, dsn string) (*Handler, error) {
	db, err := database.ConnectDB(dbms, dsn)
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

func (h *Handler) CheckDBConnection(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server Database error"})
		return
	}
}

func ResponseBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusAccepted, gin.H{"httpCode": http.StatusBadRequest, "error": err.Error()})
	return
}
