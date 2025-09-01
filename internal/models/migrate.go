package models

import (
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println("Running migrations...")
	return db.AutoMigrate(&User{}, &Task{})
}
