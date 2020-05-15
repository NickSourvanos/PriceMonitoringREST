package controllers

import (
	"PriceMonitoringService/config"
	"PriceMonitoringService/models"
	"PriceMonitoringService/repository/auth"
	user2 "PriceMonitoringService/repository/user"
	"encoding/json"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var userModel models.User
	//var dbUser models.User

	db, err := config.GetDatabaseConnection()
	if err != nil {
		errResponse := models.ResponseMessage{Message: "Database connection error"}
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}

	usersRepository := user2.NewUserRepositoryPostgres(db)

	err = json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		errResponse := models.ResponseMessage{Message: err.Error()}
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}

	user, err := usersRepository.FindByUsername(userModel.Username)
	if err != nil {
		errResponse := models.ResponseMessage{Message: "User with username " + userModel.Username + " does not exist"}
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}

	if user.Password != userModel.Password {
		errResponse := models.ResponseMessage{Message: "Invalid password provided"}
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}


	accessToken := auth.TokenHandler(strconv.FormatInt(user.Id, 10), user.Username, user.Role)
	if accessToken == "" {
		errResponse := models.ResponseMessage{Message: "Error providing access token"}
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	token := models.Token{ AccessToken: accessToken}
	_ = json.NewEncoder(w).Encode(token)
}
