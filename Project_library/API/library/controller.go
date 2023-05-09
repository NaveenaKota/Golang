package library

import (
	     "net/http"
		 "Project_library/model"
		 "encoding/json"
		"fmt"
		// "log"
		 "Project_library/db/flyway/database"
		"github.com/gorilla/mux"
		 )
var members []model.Member
func getMembers (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	db := database.Connect_to_DB()
	members = database.GetMembers_DB(db, members)
	fmt.Println(members)
	json.NewEncoder(w).Encode(members)
}

func deleteMembers (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range members {
		fmt.Println(item.Person_id)
		fmt.Println(params["id"])
		if item.Person_id == params["id"] {
			members = append(members[:index], members[index+1:]...)
			fmt.Println("Delete from Database")
			db := database.Connect_to_DB()
			database.DelMembers_DB(db, item.Person_id)
			break
		}
	}
	fmt.Println(members)
	json.NewEncoder(w).Encode(members)
}

func getMember (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range members {
		if item.Person_id == params["id"] {
			db := database.Connect_to_DB()
	        database.GetMember_DB(db, members, item.Person_id)
			json.NewEncoder(w).Encode(item)
		}
	}
}

func addMember (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var member model.Member 
	err := json.NewDecoder(r.Body).Decode(&member)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
	    return
    }
	db := database.Connect_to_DB()
	database.AddMember_DB(db, member)
	members = append(members, member)
	json.NewEncoder(w).Encode(members)
}

func updateMember (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range members {
		if item.Person_id == params["id"]{
			members = append(members[:index], members[index+1:]...)
			var member model.Member 
	        err := json.NewDecoder(r.Body).Decode(&member)
            if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
	            return
            }
			db := database.Connect_to_DB()
	        database.UpdateMember_DB(db, member)
	        members = append(members, member)
			fmt.Println(members)
			json.NewEncoder(w).Encode(members)
		}
	}
}


