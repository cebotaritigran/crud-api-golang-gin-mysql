package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"not null;primaryKey;autoIncrement"`
	FirstName string    `gorm:"not null;type:varchar(115);default:''" json:"firstname"`
	LastName  string    `gorm:"not null;type:varchar(115);default:''" json:"lastname"`
	Email     string    `gorm:"not null;type:varchar(115);default:'';unique" json:"email"`
	Password  string    `gorm:"not null;type:varchar(115);default:''" json:"password"`
	Authlevel uint      `gorm:"not null;type:int(11);default:0" json:"authlevel"`
	IsActive  bool      `gorm:"not null;type:tinyint(4);default:1;" json:"isactive"`
	CreatedAt time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp;autoUpdateTime:true;default:CURRENT_TIMESTAMP()"`
}
