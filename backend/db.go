package backend

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func (a *App) Database() {

	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Failed to get User Config Directory")
	}

	configFolder := filepath.Join(configDir, "productsusi")
	if err := os.MkdirAll(configFolder, 0755); err != nil {
		log.Fatal("Error creating folder", err)
	}

	Db, err := sql.Open("sqlite3", configFolder+"/test.db")
	if err != nil {
		log.Fatal(err)
	}

	var version string
	err = Db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) GetInfo() (string, error) {
	var version string
    err := Db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
        return "", err
	}
    return version, nil
}
