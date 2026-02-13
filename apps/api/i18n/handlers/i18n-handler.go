package i18n_handlers

import (
	"net/http"
	"pengi-med-saas/core/envelope"
	core_errors "pengi-med-saas/core/errors"
	message_models "pengi-med-saas/i18n/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MessageHandler struct {
	db *gorm.DB
}

func NewMessageHandler(db *gorm.DB) *MessageHandler {
	return &MessageHandler{db: db}
}

func (h *MessageHandler) GetAllMessages(c *gin.Context) envelope.Response {
	lang := c.Query("lang")
	if lang == "" {
		lang = "es" // Default language
	}

	messages := []message_models.Message{}
	if err := h.db.Where("lang = ?", lang).Find(&messages).Error; err != nil {
		return envelope.ErrorResponse(http.StatusInternalServerError, "Error obtaining messages", core_errors.ErrMessagesNotFound)
	}

	return envelope.SuccessResponse(messages, "Messages obtained successfully")
}

func (h *MessageHandler) GetMessageVersion(c *gin.Context) envelope.Response {
	version := "v0.1.0"
	return envelope.SuccessResponse(version, "Version obtained successfully")
}
