package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckPing(c *gin.Context) {
	if err := h.smtp.CheckSMTPConnection(); err != nil {
		ResponseBadRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
