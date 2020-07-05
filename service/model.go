package service

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const TblUser = "users"

var DB *gorm.DB

// User a table for trojan authenticator used
// see more: https://trojan-gfw.github.io/trojan/authenticator
type User struct {
	ID       int `sql:"primary_key"`
	Username string
	Password string
	Quota    int64
	Download int64
	Upload   int64
}

// Connect Dial to database
func Connect(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}

// NewUserDao .
func NewUserDao(ctx context.Context, db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

type UserDao struct {
	db *gorm.DB
}

// Create .
func (d UserDao) Create(username, pwd string, quota int64) (User, error) {
	tx := d.db.Begin()
	user := User{
		Username: username,
		Password: pwd,
		Quota:    quota,
	}
	err := tx.Create(&user).Error
	if err != nil {
		_ = tx.Rollback()
		return User{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return User{}, err
	}

	return user, nil
}

// Delete .
func (d UserDao) Delete(id int64) error {
	return d.db.Where("id=?", id).Delete(&User{}).Error
}

// Fetch .
func (d UserDao) Fetch(offset, limit int) ([]User, error) {
	rows := make([]User, 0)
	db := d.db
	err := db.Limit(limit).Offset(offset).Find(&rows).Error
	return rows, err
}
