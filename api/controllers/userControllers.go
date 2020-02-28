package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"testApi/api/models"
	"testApi/api/responses"
)
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := models.GetAllUsers(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
	return
}

func (a *App)GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	GetUser, err := user.FindUserByID(a.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, GetUser)
}