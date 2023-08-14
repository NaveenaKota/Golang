package dataservice

import (
	"Project_shopping/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "Naveena@123"
	dbname   = "shopping"
)

func ConnectToDb() *sql.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	//defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetAlbumsFromDb(db *sql.DB, albums []models.Album) []models.Album {
	fmt.Println("in database")

	rows, err := db.Query("SELECT id,title,artist,price FROM public.albums")
	fmt.Println(rows, err)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var a models.Album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			log.Fatal(err)
		}
		albums = append(albums, a)
		fmt.Println(albums)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return albums
}

func GetAlbumByIdFromDb(db *sql.DB, albums []models.Album, id string) {
	fmt.Println("----In DB--------")
	var album models.Album
	query := "SELECT id, title, artist, price FROM public.albums WHERE id = $1"
	fmt.Println(query)
	err := db.QueryRow(query, id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Got a particular album  from table successfully")
}

func PostAddAlbumToDb(db *sql.DB, albums models.Album) {
	query, err := db.Prepare("INSERT INTO public.albums (id, title, artist, price) VALUES ($1, $2, $3, $4)")
	if err != nil {
		fmt.Println("Error in the query statement preparation")
		log.Fatal(err)
	}
	defer query.Close()

	if _, err := query.Exec(albums.ID, albums.Title, albums.Artist, albums.Price); err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Added New album successfully")
}

func DeleteAlbumFromDb(db *sql.DB, id string) {
	query := fmt.Sprintf("DELETE FROM public.albums WHERE id = %s", id)

	// Execute the DELETE query
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Album with ID:", id, "deleted successfully")
}

func UpdateAlbumInDb(db *sql.DB, albums models.Album) {
	query, err := db.Prepare("UPDATE public.albums SET id = $1 , title = $2, artist = $3, Price = $4 WHERE id = $1")
	if err != nil {
		fmt.Println("Error in the query statement preparation")
		log.Fatal(err)
	}
	defer query.Close()

	if _, err := query.Exec(albums.ID, albums.Title, albums.Artist, albums.Price); err != nil {
		fmt.Println("Error in executing query")
		log.Fatal(err)
	}
	fmt.Println("Record updated successfully.")
}
