package main

import (
	"PriceMonitoringService/controllers"
	"PriceMonitoringService/repository/auth"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Hello World"!}`))
	})
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")

	router.HandleFunc("/api/users", auth.ValidateTokenMiddleware(controllers.ReturnAllUsers)).Methods("GET")
	router.HandleFunc("/api/users/{id}",  auth.ValidateTokenMiddleware(controllers.GetUser)).Methods("GET")
	router.HandleFunc("/api/users",  auth.ValidateTokenMiddleware(controllers.AddUser)).Methods("POST")
	router.HandleFunc("/api/users/{id}",  auth.ValidateTokenMiddleware(controllers.UpdateUser)).Methods("PUT")
	router.HandleFunc("/api/users/{id}",  auth.ValidateTokenMiddleware(controllers.DeleteUser)).Methods("DELETE")
	// Get logged in user products
	router.HandleFunc("/api/user/products", auth.ValidateTokenMiddleware(controllers.GetUserProducts)).Methods("GET")
	// Search user products by id
	router.HandleFunc("/api/user/{id}/products", auth.ValidateTokenMiddleware(controllers.GetProductsByUserId)).Methods("GET")

	router.HandleFunc("/api/products", auth.ValidateTokenMiddleware(controllers.GetAllProducts)).Methods("GET")
	router.HandleFunc("/api/products", auth.ValidateTokenMiddleware(controllers.AddProduct)).Methods("POST")
	router.HandleFunc("/api/products/{id}", auth.ValidateTokenMiddleware(controllers.UpdateProduct)).Methods("PUT")
	router.HandleFunc("/api/products/{id}", auth.ValidateTokenMiddleware(controllers.DeleteProduct)).Methods("DELETE")

	//http.Handle("/api/users", auth.ValidateTokenMiddleware(router))
	log.Fatal(http.ListenAndServe(":8080", router))

}





