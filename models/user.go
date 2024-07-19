package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
    ID       uuid.UUID `gorm:"type:char(36);primary_key"`
    Username string    `gorm:"unique;not null"`
    Password string    `gorm:"not null"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
    uuid := uuid.New()
    return scope.SetColumn("ID", uuid)
}
