package main
import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gorilla/mux"    
)

// This is the router for testing endpoints
func Router() *mux.Router {
    router := mux.NewRouter()
    
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    
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