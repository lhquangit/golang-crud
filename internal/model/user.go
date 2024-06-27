package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primaryKey" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Role     string `json:"role" gorm:"column:role"`
	Posts    []Post `gorm:"foreignKey:UserId"`
}

func (User) TableName() string {
	return "user"
}

// Encode password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// Check password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
