package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

//Declaring a slice
var people []Person

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<!DOCTYPE html><html><body><hr />Hello Go...<hr /></body></html>")
}

//CRUD operation
//Getting all the value from the slice named people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w,"<h1>Hello GO...</h1>")
	json.NewEncoder(w).Encode(people)
}

//Displaying a single person value
func GetPerson(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w,"<h1>Hello GO...</h1>")
	param := mux.Vars(r)

	for _, item := range people {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

//Creating a new user
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w,"<h1>Hello GO...</h1>")
	param := mux.Vars(r)
	var personStruct Person

	_ = json.NewDecoder(r.Body).Decode(&personStruct)
	personStruct.ID = param["id"]
	//appending with the slice
	people = append(people, personStruct)
	json.NewEncoder(w).Encode(people)
}

//Deletion
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w,"<h1>Hello GO...Tested withe Docker</h1>")
	param := mux.Vars(r)
	for index, item := range people {
		if item.ID == param["id"] {
			people = append(people[:index], people[index+1:]...)
		}
		json.NewEncoder(w).Encode(people)
	}
}
func main() {

	//Configuring the Listener PORT
	PORT := os.Getenv("PORT")
	if PORT == "" {
		//then manually set the port
		PORT = "3001"
	}

	router := mux.NewRouter()
	//Manually populating struct
	people = append(people, Person{
		ID:        "1",
		Firstname: "ABC",
		Lastname:  "XYZ",
		Address: &Address{
			City:  "Bangalore",
			State: "Karnataka"}})
	people = append(people, Person{
		ID:        "2",
		Firstname: "Soumyabrata",
		Lastname:  "Sen",
		Address: &Address{
			City:  "Pune",
			State: "Maharashtra"}})
	//log.Fatal(http.ListenAndServe(":8000", router))
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	router.HandleFunc("/goapi/test", viewHandler).Methods("GET")
	http.ListenAndServe(":"+PORT, router)
}
