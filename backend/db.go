package backend

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func (a *App) Database() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Failed to get User Config Directory")
	}
	Db, err := sql.Open("sqlite3", configDir+"/test.db")
	if err != nil {
		log.Fatal(err)
	}

	var version string
	err = Db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}
}
