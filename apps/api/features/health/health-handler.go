package health

import (
	"pengi-med-saas/core/envelope"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	response := envelope.SuccessResponse(nil, "ok")
	c.JSON(response.Code, response)
}
