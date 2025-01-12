package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Run() error {
	database, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %s", err)
	}
	err = database.AutoMigrate(&User{}, &Todo{})
	if err != nil {
		return fmt.Errorf("failed to migrate database schema %s", err)
	}
	log.Println("Database setup complete!")
	return nil
}
