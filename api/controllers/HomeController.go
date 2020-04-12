package controllers

import (
	"github.com/codeinbit/go-shop/api/responses"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	responses.JSON(w, http.StatusOK, "Welcome to Go Shop API")
}
