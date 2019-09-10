package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Person struct (Model)
type Person struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Address   *Address `json:"address"`
}

// Address struct (Model)
type Address struct {
	// ID string `json:"id"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
}

// Init books var as a slice Person struct
var persons []Person

// get all Persons
func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

// get single Person
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get a params

	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// create single Person
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = strconv.Itoa(rand.Intn(100000)) // not good for production
	persons = append(persons, person)
	// returns
	json.NewEncoder(w).Encode(person)
}

// update single Person
func updatePersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			var person Person
			_ = json.NewDecoder(r.Body).Decode(&person)
			person.ID = params["id"]
			persons = append(persons, person)
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(persons)
}

// delete single Person
func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(persons)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo implement DB
	persons = append(persons, Person{"1", "Diego", "Beltran", &Address{"Winnipeg", "Manitoba", "Canada"}})
	persons = append(persons, Person{"3", "Alexa", "Martinez", &Address{"Winnipeg", "Manitoba", "Canada"}})
	persons = append(persons, Person{"2", "Sanjiv", "Suresh", &Address{"Winnipeg", "Manitoba", "Canada"}})
	persons = append(persons, Person{"3", "Jagannath", "MacBeth", &Address{"Winnipeg", "Manitoba", "Canada"}})
	persons = append(persons, Person{"4", "Vikrama", "Aradhana", &Address{"Winnipeg", "Manitoba", "Canada"}})
	persons = append(persons, Person{"5", "Demetrio", "Reis", &Address{"Winnipeg", "Manitoba", "Canada"}})
	persons = append(persons, Person{"6", "Philetus", "Debenham", &Address{"Winnipeg", "Manitoba", "Canada"}})

	// Rote Handles / Endpoints
	r.HandleFunc("/api/persons", getPersons).Methods("GET")
	r.HandleFunc("/api/person/{id}", getPerson).Methods("GET")
	r.HandleFunc("/api/person", createPerson).Methods("POST")
	r.HandleFunc("/api/person/{id}", updatePersons).Methods("PUT")
	r.HandleFunc("/api/person/{id}", deletePerson).Methods("DELETE")

	fmt.Println("Server starting...")
	// to start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
