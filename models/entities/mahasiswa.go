package entities

import (
	"time"

	"gorm.io/gorm"
)

type Mahasiswa struct {
	Id        uint   `json:"id" gorm:"primarykey"`
	Nim       string `json:"nim"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index , column:deleted_at"`
}
