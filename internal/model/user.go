package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null" json:"name,omitempty"`
	Username string `gorm:"not null; unique" json:"username,omitempty"`
	Password string `gorm:"not null" json:"-"`
}

func (u User) Create(db *gorm.DB) (User, error) {
	err := db.Create(&u).Error
	return u, err
}

func (u User) Update(db *gorm.DB, newUser User) error {
	return db.Model(&u).Updates(newUser).Error
}

func (u User) Delete(db *gorm.DB) error {
	return db.Unscoped().Delete(&u).Error
}

func (u User) Get(db *gorm.DB) (User, error) {
	var user User
	var err error

	if u.Username != "" {
		db = db.Where("username = ?", u.Username)
	}

	if u.ID != 0 {
		err = db.First(&user, u.ID).Error
	} else {
		err = db.First(&user).Error
	}

	if err != nil {
		return user, err
	}

	return user, nil
}