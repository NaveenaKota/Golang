package bubblesort

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)
func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/sort", bubblesort).Methods("GET")
	fmt.Printf("starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}
