package configuration_models

import "gorm.io/gorm"

type Configuration struct {
	gorm.Model
	Key       string `json:"key"`
	Value     string `json:"value"`
	IsEnabled bool   `json:"is_enabled"`
	Feature   string `json:"feature"` // clinical | users | permissions | etc.
}
