package seeders

import (
	"github.com/codeinbit/go-shop/api/models"
	"github.com/jinzhu/gorm"
	"log"
)

var admins = []models.Admin{
	{
		Model:     gorm.Model{},
		FirstName: "Olanrewaju",
		LastName:  "Abidogun",
		UUID:      "7824-38y37-358u58-3589",
		Email:     "olanrewaju.abidogun@gmail.com",
		Password:  "qqqqqq",
	},
	{
		Model:     gorm.Model{},
		FirstName: "Femi",
		LastName:  "Badmus",
		UUID:      "7819-19y-782628-08919",
		Email:     "femi.badmus@gmail.com",
		Password:  "qqqqqq",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Admin{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.Admin{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range admins {
		err = db.Debug().Model(&models.Admin{}).Create(&admins[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
