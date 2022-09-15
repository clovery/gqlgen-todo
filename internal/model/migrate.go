package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	DB.AutoMigrate(&Todo{}, &User{})
}
