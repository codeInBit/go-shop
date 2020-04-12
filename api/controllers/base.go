package controllers


import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"

	"github.com/codeinbit/go-shop/api/models"
	"github.com/codeinbit/go-shop/api/routes"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

func (s Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string)  {
	var err error

	if Dbdriver == "mysql" {
		DbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		s.DB, err = gorm.Open(Dbdriver, DbURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == "postgres" {
		DbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		s.DB, err = gorm.Open(Dbdriver, DbURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	//database migration
	s.DB.Debug().AutoMigrate(&models.Admin{}, &models.Category{}, &models.SubCategory{}, &models.Product{})

	routes.LoadRouter()
}

func (s Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
