package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title" form:"title" valid:"required~Title of your photo is required"`
	Caption   string `json:"caption" form:"caption"`
	PhotoUrl  string `json:"photo_url" form:"photo_url" valid:"required~Photo URL of your photo is required"`
	UserID    uint   `json:"user_id"`
	User      *User
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"-,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeDelete(tx *gorm.DB) (err error) {
	// Set the DeletedAt field to the current time
	now := time.Now()
	p.DeletedAt = &now

	// Update the object to set the DeletedAt field and mark it as "not deleted"
	if err := tx.Model(p).Update("deleted_at", now).Error; err != nil {
		return err
	}

	// Return nil to indicate that the delete operation should proceed
	return nil
}
