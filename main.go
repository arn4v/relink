package main

import (
	"database/sql"
	"embed"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"

	// "fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	// "time"
)

// go:embed ./dist/*
var local embed.FS

type App struct {
	db *sql.DB
}

type ContactsDeleteReq struct {
	id string
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getExecutablePath() string {
	ex, err := os.Executable()
	checkErr(err)
	exPath := filepath.Dir(ex)
	return exPath
}

func (a *App) openDb() *sql.DB {
	dbPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkErr(err)
	dbPath = filepath.Join(dbPath, "relink.sqlite3")
	if a.db != nil {
		return a.db
	} else {
		_db, err := sql.Open("sqlite3", dbPath)
		a.db = _db
		checkErr(err)
		return a.db
	}
}

func (a *App) createTables() {
	db := a.openDb()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS contacts(id INTEGER AUTO INCREMENT PRIMARY KEY)")
	checkErr(err)
	rows, err := db.Query("SELECT * FROM contacts")
	log.Println("Test", rows)
	defer rows.Close()
}

func setResponseJson(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
}

func main() {
	a := &App{}
	a.openDb()
	a.createTables()

	upsertContactQuery, err := a.db.Prepare("")
	checkErr(err)

	updateContactQuery, err := a.db.Prepare("")
	checkErr(err)

	insertNoteQuery, err := a.db.Prepare("")
	checkErr(err)

	http.HandleFunc("/notes", func(rw http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/contacts", func(rw http.ResponseWriter, r *http.Request) {
		setResponseJson(rw)

		switch r.Method {
		case http.MethodGet:
			break
		case http.MethodPost:
			break
		case http.MethodDelete:
			var body ContactsDeleteReq
			err := json.NewDecoder(r.Body).Decode(&body)
			checkErr(err)
			break
		}
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
}
