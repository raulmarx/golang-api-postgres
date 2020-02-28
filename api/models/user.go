package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	Id    uint   `gorm:"primary_key"`
	Email string `gorm:"type:varchar(100);unique_index" json:"email"`
	Name  string `gorm:"size:100;not null"              json:"name"`
	Phone string `gorm:"size:100;not null"              json:"phone"`
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Table("usuarios").Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}
// GetAllUsers returns a list of all the user
func GetAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	if err := db.Debug().Table("usuarios").Find(&users).Error; err != nil {
		return &[]User{}, err
	}
	return &users, nil
}
