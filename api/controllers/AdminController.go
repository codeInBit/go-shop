package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/codeinbit/go-shop/api/models"
	"github.com/codeinbit/go-shop/api/responses"
	"github.com/codeinbit/go-shop/api/utilities"
	"io/ioutil"
	"net/http"

)

func (s *Server) Create(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	admin := models.Admin{}

	err = json.Unmarshal(body, &admin)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	admin.Prepare()
	err = admin.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	adminCreated, err := admin.Save(s.DB)
	if err != nil {
		formattedError := utilities.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, adminCreated.ID))
	responses.JSON(w, http.StatusCreated, adminCreated)
}

func (s Server) GetAll(w http.ResponseWriter, r http.Request) {
	admin := models.Admin{}

	admins, err := admin.GetAll(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, admins)
}
