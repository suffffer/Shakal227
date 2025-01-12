package db

import (
	"time"

	"gorm.io/gorm"
)

// Note represents a user in the system

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"not null"`
	Todos     []Todo `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Note represents a single to-do item

type Todo struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	IsCompleted bool   `gorm:"default:false"`
	UserID      uint   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
