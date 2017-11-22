package main
import (
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
    "github.com/gorilla/mux"    
)

// This is the router for testing endpoints
func Router() *mux.Router {
    router := mux.NewRouter()
    
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

    return router
}

func TestGetPeople(t *testing.T) {
    request, _ := http.NewRequest("GET", "/people", nil)
    response := httptest.NewRecorder()
    
    Router().ServeHTTP(response, request)

    assertResponseCode(200, response.Code, t)
}

func TestGetPerson(t *testing.T) {
    request, _ := http.NewRequest("GET", "/people/1", nil)
    response := httptest.NewRecorder()
    
    Router().ServeHTTP(response, request)

    assertResponseCode(200, response.Code, t)
}

func TestCreatePerson(t *testing.T) {
    inputJson := `{"Firstname": "dennis", "Lastname": "The menace"}`
    //Convert string to reader because NewRequest expect reader type as the last argument
    reader := strings.NewReader(inputJson)

    request, _ := http.NewRequest("POST", "/people/100", reader)
    response := httptest.NewRecorder()
    
    Router().ServeHTTP(response, request)

    assertResponseCode(201, response.Code, t)
}

func TestDeletePerson(t *testing.T) {
    request, _ := http.NewRequest("DELETE", "/people/1", nil)
    response := httptest.NewRecorder()
    
    Router().ServeHTTP(response, request)

    assertResponseCode(204, response.Code, t)
}

func assertResponseCode(expectedCode int, responseCode int, t *testing.T) {
    if responseCode != expectedCode {
        t.Errorf("Expected status code: %d", responseCode) //Uh-oh this means our test failed
    }
}