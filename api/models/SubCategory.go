package models

import (
	"errors"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type SubCategory struct {
	gorm.Model
	Category    Category `json:"category"`
	CategoryId  int32    `gorm:"not null" json:"category_id"`
	Name        string   `gorm:"size:100;not null" json:"name"`
	Slug        string   `gorm:"size:100;not null" json:"slug"`
	Description string   `gorm:"size:100;not null" json:"description"`
	Visible     bool     `gorm:"default:true" json:"visible"`
}

func (sc SubCategory) BeforeSave() error {
	sc.Slug = slug.Make(sc.Slug)
	return nil
}

func (sc *SubCategory) Prepare() {
	sc.ID = 0
	sc.Name = html.EscapeString(strings.TrimSpace(sc.Name))
	sc.Description = html.EscapeString(strings.TrimSpace(sc.Description))
	sc.Category = Category{}
}

func (sc *SubCategory) Validate() error {
	if sc.Name == "" {
		return errors.New("title field is required")
	}
	if sc.Description == "" {
		return errors.New("description field is required")
	}
	if sc.CategoryId < 1 {
		return errors.New("category is required")
	}

	return nil
}

func (sc *SubCategory) Save(db *gorm.DB) (*SubCategory, error) {
	var err error
	err = db.Debug().Model(&SubCategory{}).Create(&sc).Error
	if err != nil {
		return &SubCategory{}, err
	}

	if sc.ID != 0 {
		err = db.Debug().Model(&Category{}).Where("id = ?", sc.CategoryId).Take(&sc.Category).Error
		if err != nil {
			return &SubCategory{}, err
		}
	}

	return sc, nil
}

func (sc *SubCategory) GetAll(db *gorm.DB) (*[]SubCategory, error) {
	var err error
	subCategories := []SubCategory{}
	err = db.Debug().Model(&SubCategory{}).Limit(100).Find(&subCategories).Error
	if err != nil {
		return &[]SubCategory{}, err
	}

	if len(subCategories) > 0 {
		for i, _ := range subCategories {
			err = db.Debug().Model(&Category{}).Where("id = ?", subCategories[i].CategoryId).Take(&subCategories[i].Category).Error
			if err != nil {
				return &[]SubCategory{}, err
			}
		}
	}

	return &subCategories, nil
}

func (sc *SubCategory) Update(db *gorm.DB, slug string) (*SubCategory, error) {
	var err error
	err = db.Debug().Model(&SubCategory{}).Where("slug = ?", slug).Take(&SubCategory{}).Updates(
		map[string]interface{}{
			"name":        sc.Name,
			"description": sc.Description,
		}).Error
	if err != nil {
		return &SubCategory{}, err
	}

	if sc.ID != 0 {
		err = db.Debug().Model(&Category{}).Where("id = ?", sc.CategoryId).Take(&sc.Category).Error
		if err != nil {
			return &SubCategory{}, err
		}
	}

	return sc, nil
}

func (sc *SubCategory) Delete(db *gorm.DB, catId uint64, slug string) (int64, error) {
	var err error
	err = db.Debug().Model(&SubCategory{}).Where("slug = ? and category_id = ?", slug, catId).Take(&SubCategory{}).Delete(&SubCategory{}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, errors.New("subcategory not found")
		}
		return 0, err
	}
	return db.RowsAffected, err
}
