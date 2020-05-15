package controllers

import (
	"PriceMonitoringService/config"
	"PriceMonitoringService/models"
	user2 "PriceMonitoringService/repository/user"
	"PriceMonitoringService/utils"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	if user.Role == config.SystemAdministrator {
		db, err := config.GetDatabaseConnection()
		if err != nil {
			errResponse := models.ResponseMessage{Message: "Database connection error"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		usersRepository := user2.NewUserRepositoryPostgres(db)

		w.Header().Set("Content-Type", "application/json")
		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)

		newUser := models.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Password:  user.Password,
			Email:     user.Email,
			Role:      user.Role,
			Active:    user.Active,
		}

		err = usersRepository.Save(&newUser)

		if err != nil {
			errResponse := models.ResponseMessage{Message: "Error creating user, " + err.Error()}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		errResponse := models.ResponseMessage{Message: "User created successfully"}
		json.NewEncoder(w).Encode(errResponse)
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	if user.Role == config.SystemAdministrator {
		db, err := config.GetDatabaseConnection()
		if err != nil {
			errResponse := models.ResponseMessage{Message: "Database connection error"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		usersRepository := user2.NewUserRepositoryPostgres(db)

		vars := mux.Vars(r)
		key := vars["id"]

		w.Header().Set("Content-Type", "application/json")
		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)

		newUser := models.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Password:  user.Password,
			Email:     user.Email,
			Role:      user.Role,
			Active:    user.Active,
		}

		err = usersRepository.Update(key, &newUser)

		if err != nil {
			errResponse := models.ResponseMessage{Message: "Error updating user, " + err.Error()}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		response := models.ResponseMessage{Message: "User updated successfully"}
		json.NewEncoder(w).Encode(response)
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	if user.Role == config.SystemAdministrator {
		db, err := config.GetDatabaseConnection()
		if err != nil {
			errResponse := models.ResponseMessage{Message: "Database connection error"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		usersRepository := user2.NewUserRepositoryPostgres(db)

		vars := mux.Vars(r)
		key := vars["id"]

		err = usersRepository.Delete(key)

		if err != nil {
			errResponse := models.ResponseMessage{Message: "User with id " + key + " does not exist"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		response := models.ResponseMessage{Message: "User deleted successfully"}
		json.NewEncoder(w).Encode(response)
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	if user.Role == config.SystemAdministrator {

		db, err := config.GetDatabaseConnection()
		if err != nil {
			errResponse := models.ResponseMessage{Message: "Database connection error"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		usersRepository := user2.NewUserRepositoryPostgres(db)

		vars := mux.Vars(r)
		key := vars["id"]

		user, err := usersRepository.FindByID(key)

		if err != nil {
			errResponse := models.ResponseMessage{Message: "User with id " + key + " does not exist"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func ReturnAllUsers(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	if user.Role == config.SystemAdministrator {

		//var test = models.Claims{}
		val := context.Get(r, "decoded")
		user := utils.MapToStruct(val)

		if user.Role == "system-administrator" {
			db, err := config.GetDatabaseConnection()
			if err != nil {
				errResponse := models.ResponseMessage{Message: "Database connection error"}
				json.NewEncoder(w).Encode(errResponse)
				return
			}
			usersRepository := user2.NewUserRepositoryPostgres(db)

			results, err := usersRepository.FindAll()
			if err != nil {
				errResponse := models.ResponseMessage{Message: "Error fetching users list"}
				json.NewEncoder(w).Encode(errResponse)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(results)
		} else {
			errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}
