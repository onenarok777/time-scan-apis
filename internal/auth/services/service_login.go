package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"time-scan/internal/auth/models/dtx"
	"time-scan/internal/users/repository"
	"time-scan/pkg/jwt"

	"gorm.io/gorm"
)

func LoginService(req dtx.RequestLogin) (*dtx.ResponseLogin, error) {
	// Convert req.UserID from string to uint
	userIDInt, err := strconv.ParseUint(req.UserID, 10, 32)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// find user
	user, err := repository.FindUserByID(uint(userIDInt))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if !checkPassword(req.Password, uint(user.ID)) {
		return nil, errors.New("user or password invalid")
	}

	accessToken, err := jwt.GenerateToken(user)
	if err != nil {
		return nil, errors.New("can't generate token")
	}

	res := dtx.ResponseLogin{
		AccessToken: accessToken,
	}

	return &res, nil
}

func checkPassword(reqPassword string, userID uint) bool {
	// ดึงวันเดือนปีปัจจุบัน
	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")

	expectedPassword := fmt.Sprintf("%d@%s%s%s", userID, year, month, day)

	return reqPassword == expectedPassword
}
