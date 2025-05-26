package tables

import "gorm.io/gorm"

type User struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(150);not null" json:"name"`
	gorm.Model
}
