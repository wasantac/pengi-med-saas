package company_models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	Name       string            `gorm:"not null" json:"name"`
	Code       string            `gorm:"not null;unique" json:"code"`
	Features   []Feature         `gorm:"many2many:plan_features;"`
	Price      float64           `gorm:"not null" json:"price"`
	Properties datatypes.JSONMap `gorm:"type:jsonb;default:'{}'::jsonb"`
}

func (p *Plan) Save(db *gorm.DB) error {
	return db.Save(p).Error
}
