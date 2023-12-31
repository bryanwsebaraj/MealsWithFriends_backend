package models

import (
	"errors"

	"time"

	"github.com/jinzhu/gorm"
)

type College struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	College      string    `gorm:"size:100;not null" json:"college"`
	UniversityID uint32    `gorm:"primary_key;size:255;not null" json:"university"`
	City         string    `gorm:"size:100" json:"city"`
	State        string    `gorm:"size:10" json:"state"`
	Users        []User    `json:"user_list"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// just be able to read colleges. should not be able to create, update, or delete
func (c *College) FindAllColleges(db *gorm.DB) (*[]College, error) {
	var err error
	colleges := []College{}
	err = db.Debug().Model(&College{}).Limit(500).Find(&colleges).Error
	if err != nil {
		return &[]College{}, err
	}
	return &colleges, err
}

func (c *College) FindCollegesByUni(db *gorm.DB, unid uint32) (*[]College, error) {
	var err error
	colleges := []College{}
	err = db.Debug().Model(College{}).Where("university_id = ?", unid).Find(&colleges).Error
	if gorm.IsRecordNotFoundError(err) {
		return &[]College{}, errors.New("University Not Found")
	}
	if err != nil {
		return &[]College{}, err
	}

	return &colleges, err
}

func (c *College) FindCollegeByID(db *gorm.DB, cid uint32) (*College, error) {
	var err error
	err = db.Debug().Model(College{}).Where("id = ?", cid).Take(&c).Error
	if gorm.IsRecordNotFoundError(err) {
		return &College{}, errors.New("College Not Found")
	}
	if err != nil {
		return &College{}, err
	}
	return c, err
}
