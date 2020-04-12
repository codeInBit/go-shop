package models

import (
	"errors"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"size:100;not null" json:"name"`
	Slug        string `gorm:"size:100;not null" json:"slug"`
	Description string `gorm:"size:100;not null" json:"description"`
	Visible     bool   `gorm:"default:true" json:"visible"`
}

func (c Category) BeforeSave() error {
	c.Slug = slug.Make(c.Slug)
	return nil
}

func (c *Category) Prepare() {
	c.ID = 0
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.Description = html.EscapeString(strings.TrimSpace(c.Description))
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.New("title field is required")
	}
	if c.Description == "" {
		return errors.New("description field is required")
	}
	return nil
}

func (c *Category) Save(db gorm.DB) (*Category, error) {
	var err error
	err = db.Debug().Model(&Category{}).Create(&c).Error
	if err != nil {
		return &Category{}, err
	}

	return c, nil
}

func (c *Category) GetAll(db gorm.DB) (*[]Category, error) {
	var err error
	categories := []Category{}
	err = db.Debug().Model(&Category{}).Limit(100).Find(&categories).Error
	if err != nil {
		return &[]Category{}, err
	}

	return &categories, err
}

func (c *Category) Update(db gorm.DB, slug string) (*Category, error) {
	var err error
	err = db.Debug().Model(&Category{}).Where("slug = ?", slug).Take(&Category{}).Updates(
		map[string]interface{}{
			"name":        c.Name,
			"description": c.Description,
		}).Error
	if err != nil {
		return &Category{}, err
	}

	err = db.Debug().Model(&Category{}).Where("slug = ?", slug).Take(&c).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil
}

func (c Category) Delete(db gorm.DB, slug string) (int64, error) {
	var err error
	err = db.Debug().Model(&Admin{}).Where("slug = ?", slug).Take(&Admin{}).Delete(&Admin{}).Error
	if err != nil {
		return 0, db.Error
	}
	return db.RowsAffected, err
}
