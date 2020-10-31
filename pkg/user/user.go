// Package user provides ...
package user

import (
	"time"

	"github.com/deepzz0/appdemo/pkg/db"
)

// User user info
type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"not null;index:idx_username,unique"`
	Password string `gorm:"not null"`

	UserAgent string    `gorm:"not null"`
	CreatedIP string    `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

// InsertUser insert user
func InsertUser(u *User) error {
	return db.DB.Create(u).Error
}

// DeleteUser delete user
func DeleteUser(id int) error {
	return db.DB.Where("id=?", id).Delete(User{}).Error
}

// UpdateUser update user
func UpdateUser(id int, fields map[string]interface{}) error {
	return db.DB.Model(User{}).Where("id=?", id).
		Updates(fields).Error
}

// SelectUser select user
func SelectUser(id int) (*User, error) {
	u := new(User)
	err := db.DB.Where("id=?", id).First(u).Error
	return u, err
}

// SelectUserByUsername select user by username
func SelectUserByUsername(username string) (*User, error) {
	u := new(User)
	err := db.DB.Where("username=?", username).First(u).Error
	return u, err
}
