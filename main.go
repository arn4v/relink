package main

import (
	"embed"
	"encoding/json"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

//go:embed dist
var staticFiles embed.FS

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type App struct {
	dbPath string
}

func normalizeOsPath(path string) string {
	if runtime.GOOS == "windows" {
		return filepath.FromSlash(path)
	} else {
		return path
	}
}

func (app *App) dataHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		dbFile, err := os.Open(app.dbPath)

		if err != nil {
			rw.WriteHeader(500)
			io.WriteString(rw, "Unable to read file")
			return
		}

		defer dbFile.Close()

		io.Copy(rw, dbFile)

		break
	case http.MethodPost:
		dbFile, err := os.OpenFile(app.dbPath, os.O_WRONLY|os.O_TRUNC, 0644)

		if err != nil {
			rw.WriteHeader(500)
			io.WriteString(rw, "Unable to read file")
			return
		}

		defer dbFile.Close()

		_, err = io.Copy(dbFile, r.Body)

		if err != nil {
			rw.WriteHeader(500)
			return
		}

		break
	}

}

func (app *App) createDatabaseIfNotExists() {
	app.getDbPath()
	if _, err := os.Stat(app.dbPath); os.IsNotExist(err) {
		var db map[string]interface{}
		json.Unmarshal([]byte("{}"), &db)
		json, err := json.MarshalIndent(db, "", "")
		check(err)
		err = ioutil.WriteFile(app.dbPath, json, 0644)
		check(err)
	}
}

func (app *App) getDbPath() {
	exPath, err := os.Executable()
	check(err)
	dirPath := filepath.Dir(exPath)
	check(err)
	var path string = normalizeOsPath(path.Join(dirPath, "db.json"))
	app.dbPath = path
}

func main() {
	var app App = App{}
	app.createDatabaseIfNotExists()

	var isDev bool = os.Getenv("NODE_ENV") != "development"

	if isDev {
		fsys, err := fs.Sub(staticFiles, "dist")
		check(err)
		fs := http.FileServer(http.FS(fsys))
		http.Handle("/", fs)
	}

	http.HandleFunc("/data", app.dataHandler)

	log.Println("Server listening at port :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
