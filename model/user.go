package model

import (
	"btpn-backend-go/config"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"not null" json:"username" form:"username" valid:"required~Your username is required"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string     `gorm:"not null" json:"-" form:"password" valid:"required~Your password is required,minstringlength(8)~Password has to have a minimum length of 8 characters"`
	Photos    []Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"-,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = config.HashPass(u.Password)

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = config.HashPass(u.Password)

	err = nil
	return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	// Set the DeletedAt field to the current time
	now := time.Now()
	u.DeletedAt = &now

	// Update the object to set the DeletedAt field and mark it as "not deleted"
	if err := tx.Model(u).Update("deleted_at", now).Error; err != nil {
		return err
	}

	// Return nil to indicate that the delete operation should proceed
	return nil
}
