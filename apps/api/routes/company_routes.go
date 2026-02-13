package routes

import (
	"pengi-med-saas/core/envelope"
	"pengi-med-saas/core/logger"
	company_handlers "pengi-med-saas/features/companies/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCompanyRoutes(router *gin.RouterGroup, db *gorm.DB) {
	companyHandler := company_handlers.NewCompanyHandler(db, logger.Log)

	group := router.Group("/companies")
	{
		group.GET("", envelope.Handle(companyHandler.GetCompanies))
	}
}
