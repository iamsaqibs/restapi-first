package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name:"`
	Dob string `json:"dob"`
	Phone string `json:"phone"`
}

type Users []User

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func aboutPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "About Page hit")
}

func allUsers(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: All Users")
	user := Users{
		User{
			Name:  "Saqib",
			Dob:   "1995",
			Phone: "03355600939",
		},
	}

	json.NewEncoder(w).Encode(user)
}

func handleRequest(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/about", aboutPage).Methods("GET")
	router.HandleFunc("/allUsers", allUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main(){
	handleRequest()
}
