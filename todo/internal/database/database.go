package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

func Run() error {

	if err != nil {
		return fmt.Errorf("failed to connect to the database: %s", err)
	}
	err = DB.AutoMigrate(&User{}, &Todo{})
	if err != nil {
		return fmt.Errorf("failed to migrate database schema %s", err)
	}
	log.Println("Database setup complete!")
	return nil
}
