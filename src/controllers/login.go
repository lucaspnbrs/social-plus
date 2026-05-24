package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"social-plus/src/auth"
	"social-plus/src/db"
	"social-plus/src/models"
	"social-plus/src/repositories"
	"social-plus/src/responses"
	"social-plus/src/security"
)

//Login Route is responsible for authenticating a responsible a party in the API
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.ERROR(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryFromUsers(db)
	userSaveInDb, erro := repository.SearchUserByEmail(user.Email)

	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
	}

	if erro = security.VerifyPass(userSaveInDb.Pass, user.Pass); erro != nil {
		responses.ERROR(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.CreateToken(userSaveInDb.ID)
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return 
	}

	w.Write([]byte(token))
}