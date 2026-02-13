package tenant_models

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	Name string `gorm:"not null"`
	Slug string `gorm:"not null;unique"`
}

func NewTenant(name string) *Tenant {
	return &Tenant{
		Name: name,
	}
}

func (t *Tenant) Save(db *gorm.DB) error {
	return db.Save(t).Error
}
