package services

import (
	"net/http"

	"github.com/dev-hana/go-mailer/database"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMail(c *gin.Context) {
	var mail database.SendMail
	if err := c.ShouldBindJSON(&mail); err != nil {
		ResponseBadRequest(c, err)
		return
	}

	if err := h.db.CreateMail(mail); err != nil {
		ResponseBadRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
