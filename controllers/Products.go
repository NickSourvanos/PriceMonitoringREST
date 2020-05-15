package controllers

import (
	"PriceMonitoringService/config"
	"PriceMonitoringService/models"
	"PriceMonitoringService/repository/product"
	"PriceMonitoringService/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	if user.Role == config.ProductsAdministrator {
		db, err := config.GetDatabaseConnection()
		if err != nil {
			errResponse := models.ResponseMessage{Message: "Database connection error"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		productsRepository := product.NewProductRepositoryPostgres(db)

		w.Header().Set("Content-Type", "application/json")
		var product models.Product
		_ = json.NewDecoder(r.Body).Decode(&product)

		newProduct := models.Product{
			ModelName:      product.ModelName,
			Version:        product.Version,
			Price:          product.Price,
			Description:    product.Description,
			ProductionDate: product.ProductionDate,
		}

		fmt.Println(newProduct)

		err = productsRepository.Save(&newProduct)

		if err != nil {
			errResponse := models.ResponseMessage{Message: "Error inserting product: " + err.Error()}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		errResponse := models.ResponseMessage{Message: "Product inserted successfully"}
		json.NewEncoder(w).Encode(errResponse)
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	db, err := config.GetDatabaseConnection()
	if err != nil {
		errResponse := models.ResponseMessage{Message: "Database connection error"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
	productRepository := product.NewProductRepositoryPostgres(db)

	vars := mux.Vars(r)
	key := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	if user.Role == config.ProductsAdministrator {

		var product models.Product
		_ = json.NewDecoder(r.Body).Decode(&product)

		newProduct := models.Product{
			ModelName:      product.ModelName,
			Version:        product.Version,
			Price:          product.Price,
			Description:    product.Description,
			ProductionDate: product.ProductionDate,
		}

		err = productRepository.Update(key, &newProduct)

		if err != nil {
			errResponse := models.ResponseMessage{Message: "Error updating product: " + err.Error()}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		response := models.ResponseMessage{Message: "Product updated successfully"}
		json.NewEncoder(w).Encode(response)

	} else if user.Role == config.Retailer {

		var product models.Product
		_ = json.NewDecoder(r.Body).Decode(&product)

		newProduct := models.Product{
			Price:          product.Price,
		}

		err = productRepository.UpdatePrice(key, &newProduct)

		if err != nil {
			errResponse := models.ResponseMessage{Message: "Error updating product: " + err.Error()}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		response := models.ResponseMessage{Message: "Product updated successfully"}
		json.NewEncoder(w).Encode(response)
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	if user.Role == config.ProductsAdministrator {

		db, err := config.GetDatabaseConnection()
		if err != nil {
			errResponse := models.ResponseMessage{Message: "Database connection error"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		productsRepository := product.NewProductRepositoryPostgres(db)

		vars := mux.Vars(r)
		key := vars["id"]

		err = productsRepository.Delete(key)

		if err != nil {
			errResponse := models.ResponseMessage{Message: "Product with id " + key + " does not exist"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		response := models.ResponseMessage{Message: "Product deleted successfully"}
		json.NewEncoder(w).Encode(response)
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func GetUserProducts(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	if user.Role == config.Retailer {
		db, err := config.GetDatabaseConnection()
		if err != nil {
			errResponse := models.ResponseMessage{Message: "Database connection error"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		productsRepository := product.NewProductRepositoryPostgres(db)

		products, err := productsRepository.FindProductsByUserId(user.UserId)

		if err != nil {
			errResponse := models.ResponseMessage{Message: "Error fetching products"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		if len(products) == 0 {
			errResponse := models.ResponseMessage{Message: "User with id " + user.UserId + " has no products"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func GetProductsByUserId(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	if user.Role == config.ProductsAdministrator || user.Role == config.SystemAdministrator {
		db, err := config.GetDatabaseConnection()
		if err != nil {
			errResponse := models.ResponseMessage{Message: "Database connection error"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		productsRepository := product.NewProductRepositoryPostgres(db)

		vars := mux.Vars(r)
		key := vars["id"]

		products, err := productsRepository.FindProductsByUserId(key)

		if err != nil {
			errResponse := models.ResponseMessage{Message: "User with id " + key + " does no t exist"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		if len(products) == 0 {
			errResponse := models.ResponseMessage{Message: "User with id " + key + " has no products"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	val := context.Get(r, "decoded")
	user := utils.MapToStruct(val)

	if user.Role == config.ProductsAdministrator || user.Role == config.SystemAdministrator {
		db, err := config.GetDatabaseConnection()
		if err != nil {
			errResponse := models.ResponseMessage{Message: "Database connection error"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		productsRepository := product.NewProductRepositoryPostgres(db)

		products, err := productsRepository.FindAll()
		if err != nil {
			errResponse := models.ResponseMessage{Message: "Error fetching products"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		if len(products) == 0 {
			errResponse := models.ResponseMessage{Message: "No products found"}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	} else {
		errResponse := models.ResponseMessage{Message: "This role does not have permission to access this endpoint"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}
