package library
import (
	"database/sql"
	"log"
	"fmt"
	"Project_library/model"
)

func GetMembers_DB(db *sql.DB, members []model.Member) []model.Member{
	rows, err := db.Query(`SELECT * FROM member`)
	if err != nil {
		log.Fatal(err)
	} 
	defer rows.Close()

	for rows.Next() {
		var member model.Member
		err := rows.Scan(&member.Person_id, &member.Name, &member.PhoneNumber, &member.Email, &member.Department)
		if err != nil {
			log.Fatal(err)
		}
		members = append(members, member)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return members
	
}

func DelMembers_DB (db *sql.DB, id string) {
	fmt.Println("---in Db---")
	_, err := db.Query("Delete from member where person_id = ?", id)
	if err != nil {
		log.Fatal(err)
		fmt.Println("------In err-------------")
	} 
	fmt.Println("Deleted from members table successfully")
}

func GetMember_DB (db *sql.DB, members []model.Member, id string) {
	fmt.Println("----In DB--------")
	var member model.Member
	query := `select person_id,name,phoneNumber,email,department from member where person_id = ?`
	fmt.Println(query)
	err := db.QueryRow(query, id).Scan(&member.Person_id, &member.Name, &member.PhoneNumber, &member.Email, &member.Department)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Got a particular member  from table successfully")
}

func AddMember_DB(db *sql.DB, member model.Member){
	query := `INSERT INTO member (person_id, name, phoneNumber, email, department)
              VALUES (?, ?, ?, ?, ?)`
	fmt.Println(query)
	_, err := db.Exec(query, member.Person_id, member.Name, member.PhoneNumber, member.Email, member.Department)
	fmt.Println(member.Person_id, member.Name, member.PhoneNumber, member.Email, member.Department)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Addd new member  to table successfully")
}

func UpdateMember_DB (db *sql.DB, member model.Member){
	query := `UPDATE member SET person_id = ? , name = ?, phoneNumber = ?, email = ?, department = ? WHERE person_id = ?`
fmt.Println(query)
	_, err := db.Exec(query, member.Person_id, member.Name, member.PhoneNumber, member.Email, member.Department, member.Person_id) // check err
	if err != nil {
		log.Fatal(err)
		fmt.Println("----IN DB ERROR-----")
	}
	fmt.Println("Record updated successfully.")
}


