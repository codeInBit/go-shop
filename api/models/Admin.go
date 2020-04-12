package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/codeinbit/go-shop/api/utilities"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	FirstName string `gorm:"size:100;not null" json:"firstname"`
	LastName  string `gorm:"size:100;not null" json:"lastname"`
	UUID      string `gorm:"size:255;not null;unique" json:"uuid"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
}

func (a *Admin) BeforeSave() error {
	hashedPassword, err := utilities.Hash(a.Password)
	if err != nil {
		return err
	}

	a.Password = string(hashedPassword)
	a.UUID = uuid.New().String()
	return nil
}

func (a *Admin) Prepare() {
	a.ID = 0
	a.FirstName = html.EscapeString(strings.TrimSpace(a.FirstName))
	a.LastName = html.EscapeString(strings.TrimSpace(a.LastName))
	a.Email = html.EscapeString(strings.TrimSpace(a.Email))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Admin) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if a.FirstName == "" {
			return errors.New("please enter first name")
		}
		if a.LastName == "" {
			return errors.New("please enter last name")
		}
		if a.Email == "" {
			return errors.New("please enter email address")
		}
		if a.Password == "" {
			return errors.New("please nter password")
		}
		if len(a.Password) < 6 {
			return errors.New("password should be at least 6 characters")
		}
		return nil

	case "login":
		if a.Password == "" {
			return errors.New("password is required")
		}
		if len(a.Password) < 6 {
			return errors.New("password should be at least 6 characters")
		}
		if a.Email == "" {
			return errors.New("email address is required")
		}
		if err := checkmail.ValidateFormat(a.Email); err != nil {
			return errors.New("invalid email address")
		}
		return nil

	default:
		if a.Password == "" {
			return errors.New("password is required")
		}
		if len(a.Password) < 6 {
			return errors.New("password should be at least 6 characters")
		}
		if a.Email == "" {
			return errors.New("email address is required")
		}
		if err := checkmail.ValidateFormat(a.Email); err != nil {
			return errors.New("invalid email address")
		}
		return nil
	}
}

func (a *Admin) Save(db *gorm.DB) (*Admin, error) {
	var err error
	err = db.Debug().Create(&a).Error
	if err != nil {
		return &Admin{}, err
	}

	return a, nil
}

func (a *Admin) GetAll(db *gorm.DB) (*[]Admin, error) {
	var err error
	admins := []Admin{}
	err = db.Debug().Model(&Admin{}).Limit(100).Find(&admins).Error
	if err != nil {
		return &[]Admin{}, err
	}

	return &admins, err
}

func (a *Admin) GetAdminByEmail(db *gorm.DB, email string) (*Admin, error) {
	var err error
	err = db.Debug().Model(&Admin{}).Where("email = ?", email).Error
	if err != nil {
		return &Admin{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Admin{}, errors.New("no admin found")
	}
	return a, err
}

func (a *Admin) Update(db *gorm.DB, uuid string) (*Admin, error) {
	var err error
	err = db.Debug().Model(&Admin{}).Where("uuid = ?", uuid).Take(&Admin{}).Updates(
		map[string]interface{}{
			"firstname": a.FirstName,
			"lastname":  a.LastName,
			"email":     a.Email,
			"password":  a.Password,
		}).Error
	if err != nil {
		return &Admin{}, db.Error
	}

	err = db.Debug().Model(&Admin{}).Where("uuid = ?", uuid).Take(&a).Error
	if err != nil {
		return &Admin{}, err
	}
	return a, nil
}

func (a *Admin) Delete(db *gorm.DB, uuid string) (int64, error) {
	var err error
	err = db.Debug().Model(&Admin{}).Where("uuid = ?", uuid).Take(&Admin{}).Delete(&Admin{}).Error
	if err != nil {
		return 0, db.Error
	}
	return db.RowsAffected, err
}
