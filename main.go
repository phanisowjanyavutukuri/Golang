package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rest/data"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"first,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type ResponsePerson struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"first,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
	Message   string   `json:"message"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {

	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println(router)

	//router.HandleFunc("/people", GetPeople).Methods("GET")
}
func GetPeople(w http.ResponseWriter, r *http.Request) {
	buffer, _ := json.Marshal(people)
	w.Write(buffer)
	data.DisplayAll()
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, items := range people {
		if items.ID == params["id"] {
			json.NewEncoder(w).Encode(items)
		}
	}
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var personDetails Person
	buffer, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(buffer, &personDetails)
	responsePerson := Person{
		ID:        personDetails.ID,
		Firstname: personDetails.Firstname,
		Lastname:  personDetails.Lastname,
		Address:   personDetails.Address,
	}
	buffer, _ = json.Marshal(responsePerson)
//	data.Insert(buffer)
	//w.Write(buffer)
}
