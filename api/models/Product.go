package models

import (
	"errors"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type Product struct {
	gorm.Model
	SubCategory   SubCategory `json:"category"`
	SubCategoryId int32       `gorm:"not null" json:"category_id"`
	Name          string      `gorm:"size:100;not null" json:"name"`
	Slug          string      `gorm:"size:100;not null" json:"slug"`
	Description   string      `gorm:"size:100;not null" json:"description"`
	Price         float64     `gorm:"size:255;not null" json:"price"`
	Visible       bool        `gorm:"default:true" json:"visible"`
}

func (p Product) BeforeSave() error {
	p.Slug = slug.Make(p.Slug)
	return nil
}

func (p *Product) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	//p.Price = html.EscapeString(strings.TrimSpace(p.Price))
	p.SubCategory = SubCategory{}
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("title field is required")
	}
	if p.Description == "" {
		return errors.New("description field is required")
	}
	if p.Price == 0.0 {
		return errors.New("price field is required")
	}
	if p.SubCategoryId < 1 {
		return errors.New("category is required")
	}

	return nil
}

func (p *Product) Save(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&SubCategory{}).Where("id = ?", p.SubCategoryId).Take(&p.SubCategory).Error
		if err != nil {
			return &Product{}, err
		}
	}

	return p, nil
}

func (p *Product) GetAll(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}

	if len(products) > 0 {
		for i, _ := range products {
			err = db.Debug().Model(&SubCategory{}).Where("id = ?", products[i].SubCategoryId).Take(&products[i].SubCategory).Error
			if err != nil {
				return &[]Product{}, err
			}
		}
	}

	return &products, nil
}

func (p *Product) Update(db *gorm.DB, slug string) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Where("slug = ?", slug).Take(&Product{}).Updates(
		map[string]interface{}{
			"name":        p.Name,
			"price":       p.Price,
			"description": p.Description,
		}).Error
	if err != nil {
		return &Product{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&SubCategory{}).Where("id = ?", p.SubCategoryId).Take(&p.SubCategory).Error
		if err != nil {
			return &Product{}, err
		}
	}

	return p, nil
}

func (p *Product) Delete(db *gorm.DB, subCatId uint64, slug string) (int64, error) {
	var err error
	err = db.Debug().Model(&Product{}).Where("slug = ? and sub_category_id = ?", slug, subCatId).Take(&Product{}).Delete(&Product{}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, errors.New("product not found")
		}
		return 0, err
	}
	return db.RowsAffected, err
}
