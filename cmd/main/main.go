package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/drsimplegraffiti/gobook/pkg/routes"
	"fmt"
)

func main(){
	r := mux.NewRouter() 
	// routes.RegisterBookStoreRoutes(r) 
	// http.Handle("/", r)
	// fmt.Println("Server listening on port 8081")
	// log.Fatal(http.ListenAndServe("localhost:8081", r))

	
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	routes.RegisterBookStoreRoutes(apiV1)
	

	fmt.Println("Server listening on port 8081")
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}



