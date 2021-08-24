package main

import (
	"embed"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	_ "github.com/mattn/go-sqlite3"

	// "fmt"
	"log"
	"net/http"
	// "time"
)

// go:embed ./dist/*
var local embed.FS

/* -------------------------------------------------------------------------- */
/*                                    Utils                                   */
/* -------------------------------------------------------------------------- */

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func normalizeOsPath(path string) string {
	if runtime.GOOS == "windows" {
		return filepath.FromSlash(path)
	} else {
		return path
	}
}

type App struct {
	dbPath string
	db     Db
}

type Meeting struct {
	Id        string    `json:"note"`
	Note      string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type Contact struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Location string    `json:"location"`
	Emails   []string  `json:"emails"`
	Meetings []Meeting `json:"meetings"`
}

type Db struct {
	Contacts []Contact `json:"contants"`
}

func setResponseJson(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
}

func (app *App) init() {
	exPath, err := os.Executable()
	check(err)
	dirPath := filepath.Dir(exPath)
	check(err)
	var path string = normalizeOsPath(path.Join(dirPath, "db.json"))
	app.dbPath = path

	if _, err := os.Stat(path); os.IsNotExist(err) {
		app.db = Db{
			Contacts: []Contact{},
		}
		json, err := json.MarshalIndent(app.db, "", "")
		check(err)
		err = ioutil.WriteFile(path, json, 0644)
		check(err)
	} else {
		var db Db
		file, err := ioutil.ReadFile(app.dbPath)
		check(err)
		check(json.Unmarshal(file, &db))
		app.db = db
	}
	// json, err := json.MarshalIndent(app.db, "", "")
}

func decodeBody(r *http.Request, v interface{}) {
	err := json.NewDecoder(r.Body).Decode(v)
	check(err)
}

/* -------------------------------------------------------------------------- */
/*                                Notes handler                               */
/* -------------------------------------------------------------------------- */

type NotesGetBody struct {
	MeetingId string `json:"meeting_id"`
}

type NotesPostBody struct {
	MeetingId string `json:"meeting_id"`
	Note      string `json:"note"`
}

type NotesDeleteBody struct {
	MeetingId string `json:"meeting_id"`
	NoteId    string `json:"id"`
}

func (app *App) notesHandler(rw http.ResponseWriter, r *http.Request) {
	setResponseJson(rw)

	switch r.Method {
	case http.MethodGet:
		var body NotesGetBody
		decodeBody(r, body)
		break
	case http.MethodPost:
		break
	case http.MethodDelete:
		break
	}
}

/* -------------------------------------------------------------------------- */
/*                              Contacts handler                              */
/* -------------------------------------------------------------------------- */

func (app *App) contactsHandler(rw http.ResponseWriter, r *http.Request) {
	setResponseJson(rw)

	switch r.Method {
	case http.MethodGet:
		break
	case http.MethodPost:
		break
	case http.MethodDelete:
		break
	}
}

/* -------------------------------------------------------------------------- */
/*                               Meeting handler                              */
/* -------------------------------------------------------------------------- */

func (app *App) meetingsHandler(rw http.ResponseWriter, r *http.Request) {
	setResponseJson(rw)

	switch r.Method {
	case http.MethodGet:
		break
	case http.MethodPost:
		break
	case http.MethodDelete:
		break
	}
}

func main() {
	app := App{}
	app.init()

	http.HandleFunc("/notes", app.notesHandler)
	http.HandleFunc("/meetings", app.meetingsHandler)
	http.HandleFunc("/contacts", app.contactsHandler)

	log.Fatal(http.ListenAndServe(":4000", nil))
	log.Println("Server listening at port :4000")
}
