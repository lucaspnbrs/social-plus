package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"social-plus/src/auth"
	"social-plus/src/db"
	"social-plus/src/models"
	"social-plus/src/repositories"
	"social-plus/src/responses"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//Fetch Users
func FetchUsers( w http.ResponseWriter, r *http.Request) {
    nomeOuNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := db.Connect()
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return 	
	}
	defer db.Close()

	repository := repositories.NewRepositoryFromUsers(db)
	users, erro := repository.Search(nomeOuNick)
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return 
	}

	responses.JSON(w, http.StatusOK,users)
}

//Fetch Users By ID
func FetchUsersByID( w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
		responses.ERROR(w, http.StatusBadRequest, erro)
		return 
	}
	
	db, erro := db.Connect()
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	respository := repositories.NewRepositoryFromUsers(db)
	user, erro := respository.FetchUserID(userID)
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

//Create Users
func CreateUser( w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := ioutil.ReadAll((r.Body))
	if erro != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, erro)
		return 
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.ERROR(w, http.StatusBadRequest, erro)
		return
	}

	if erro := user.Prepare("register"); erro != nil {
		responses.ERROR(w, http.StatusBadRequest, erro)
	}

	database, erro := db.Connect()
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewRepositoryFromUsers(database)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return 
	}
	
	responses.JSON(w, http.StatusCreated, user)
	
}

//Updating Users
func UpUsers( w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
		responses.ERROR(w, http.StatusBadRequest, erro)
		return 
	}

	userIDinToken, erro := auth.ExtractUserWithID(r)
	if erro != nil {
		responses.ERROR(w, http.StatusUnauthorized, erro)
		return 
	}

	if userID != userIDinToken {
		responses.ERROR(w, http.StatusForbidden, erro)
		return
	}

	fmt.Println(userIDinToken)

	bodyRequest, erro := ioutil.ReadAll(r.Body) //it use to read the body of the request
	if erro != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.ERROR(w, http.StatusBadRequest, erro)
		return 
	}

	if erro = user.Prepare("edit"); erro != nil {
		responses.ERROR(w, http.StatusBadRequest, erro)
		return 
	}

	database, erro := db.Connect()
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return 
	}
	defer database.Close()

	repository := repositories.NewRepositoryFromUsers(database)
	if erro = repository.UpdateUser(userID, user); erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

//Delete Users
func DeleteUsers( w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return 
	}

	userIDinToken, erro := auth.ExtractUserWithID(r)
	if erro != nil {
		responses.ERROR(w, http.StatusUnauthorized, erro)
		return
	}

	if userIDinToken != userID {
		responses.ERROR(w, http.StatusForbidden, errors.New("It's not possible delete a user that not yours!"))
		return 
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.ERROR(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repository := repositories.NewRepositoryFromUsers(db)
    if erro = repository.DeleteUser(userID); erro != nil {
		responses.ERROR(w, http.StatusBadRequest, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
