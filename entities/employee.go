package entities

import "gorm.io/gorm"

type Employee struct {
	ID     int    `gorm:"primary key;autoIncrement" json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

func MigrateEmployee(db *gorm.DB) error {
	err := db.AutoMigrate(&Employee{})
	return err
}
