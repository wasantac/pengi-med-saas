package envelope

import (
	core_errors "pengi-med-saas/core/errors"

	"github.com/gin-gonic/gin"
)

type Action func(r *gin.Context) Response

func Handle(action Action) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := action(c)

		// Translate response if translator is available
		if val, exists := c.Get("translator"); exists {
			if translate, ok := val.(func(string) string); ok {
				response.Message = translate(response.Message)
				if appErr, ok := response.Data.(core_errors.AppError); ok {
					appErr.ErrorMessage = translate(appErr.ErrorCode)
					response.Data = appErr
				}
			}
		}

		if response.Code > 399 {
			c.JSON(response.Code, response)
			return
		}

		c.JSON(response.Code, response)
	}
}
