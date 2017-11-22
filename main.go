package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "./models"
)

var people []models.Person

// our main function
func main() {
    // Sets up the router
    router := mux.NewRouter()

    people = append(people, models.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &models.Address{City: "City X", State: "State X"}})
    people = append(people, models.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &models.Address{City: "City Z", State: "State Y"}})
    people = append(people, models.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

    // The endpoints, each calling a function to create the response
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

    port := ":8000"
    fmt.Println("The server has started and is listening")

    listenErr := http.ListenAndServe(port, router)
    if listenErr != nil {
      log.Fatal(listenErr)
    }

}


// The functions responsible for generating the response for the endpoints
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }

    json.NewEncoder(w).Encode(&models.Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var person models.Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    
    person.ID = params["id"]
    people = append(people, person)

    // Make it return 201 status
    w.WriteHeader(http.StatusCreated)

    json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }

        // Make it return 204 status
        w.WriteHeader(http.StatusNoContent)

        json.NewEncoder(w).Encode(people)
    }    
}

