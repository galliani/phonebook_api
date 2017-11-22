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

    if response.Code != 200 {
        t.Errorf("Expected status code: %d", response.Code) //Uh-oh this means our test failed
    }
}

func TestGetPerson(t *testing.T) {
    request, _ := http.NewRequest("GET", "/people/1", nil)
    response := httptest.NewRecorder()
    
    Router().ServeHTTP(response, request)

    if response.Code != 200 {
        t.Errorf("Expected status code: %d", response.Code) //Uh-oh this means our test failed
    }
}

func TestCreatePerson(t *testing.T) {
    inputJson := `{"Firstname": "dennis", "Lastname": "The menace"}`
    //Convert string to reader because NewRequest expect reader type as the last argument
    reader := strings.NewReader(inputJson)

    request, _ := http.NewRequest("POST", "/people/100", reader)
    response := httptest.NewRecorder()
    
    Router().ServeHTTP(response, request)

    if response.Code != 201 {
        t.Errorf("Expected status code: %d", response.Code) //Uh-oh this means our test failed
    }
}

func TestDeletePerson(t *testing.T) {
    request, _ := http.NewRequest("DELETE", "/people/1", nil)
    response := httptest.NewRecorder()
    
    Router().ServeHTTP(response, request)

    if response.Code != 204 {
        t.Errorf("Expected status code: %d", response.Code) //Uh-oh this means our test failed
    }
}
