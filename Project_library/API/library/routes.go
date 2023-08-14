package library

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)
func Routes() {
	r := mux.NewRouter()

	r.HandleFunc("/members", getMembers).Methods("GET")
	r.HandleFunc("/members/{id}", getMember).Methods("GET")
	r.HandleFunc("/members", addMember).Methods("POST")
	r.HandleFunc("/members/{id}", updateMember).Methods("PUT")
	r.HandleFunc("/members/{id}", deleteMembers).Methods("DELETE")

	fmt.Printf("starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}

