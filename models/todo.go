package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Todo struct {
    ID        uuid.UUID `gorm:"type:char(36);primary_key"`
    UserID    uuid.UUID `gorm:"type:char(36)"`
    Title     string    `gorm:"not null"`
    Completed bool      `gorm:"default:false"`
}

func (todo *Todo) BeforeCreate(scope *gorm.Scope) error {
    uuid := uuid.New()
    return scope.SetColumn("ID", uuid)
}
