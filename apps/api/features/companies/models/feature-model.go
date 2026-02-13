package company_models

import (
	permission_models "pengi-med-saas/features/permissions/models"

	"gorm.io/gorm"
)

type Feature struct {
	gorm.Model
	Code        string                         `gorm:"not null;unique" json:"code"`
	Name        string                         `gorm:"not null" json:"name"`
	Permissions []permission_models.Permission `gorm:"many2many:feature_permissions;" json:"permissions"`
}

func (f *Feature) Save(db *gorm.DB) error {
	return db.Save(f).Error
}
