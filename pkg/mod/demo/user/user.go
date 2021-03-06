// Package user provides ...
package user

import (
	"time"

	"github.com/deepzz0/appdemo/pkg/connector/db"
)

func init() {
	db.DB.AutoMigrate(User{})
}

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

// SimpleUser simple user
type SimpleUser struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ForShow user for show
func (u *User) ForShow() SimpleUser {
	return SimpleUser{
		ID:        u.ID,
		Username:  u.Username,
		UpdatedAt: u.UpdatedAt,
		CreatedAt: u.CreatedAt,
	}
}

// IsExistUser exist user?
func IsExistUser(username string) (bool, error) {
	var count int64 = 0
	err := db.DB.Model(User{}).Where("username=?", username).
		Count(&count).Error
	return count > 0, err
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
