package repository

import (
	"time-scan/internal/users/models/tables"
	"time-scan/pkg/gorm"
)

func FindUserByID(id uint) (*tables.User, error) {
	// get db
	db := gorm.DB

	// find db
	var user tables.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
