package services

import (
	"github.com/dev-hana/go-mailer/database"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSMTP(c *gin.Context) {
	var smtp database.SMTP
	if err := c.ShouldBindJSON(&smtp); err != nil {
		ResponseBadRequest(c, err)
	}
}
