package backend

import (
	"context"
	"database/sql"
	"log"
)

type App struct {
	ctx context.Context
	db  *sql.DB
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	if err := a.Database(); err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	log.Println("Wails application started up again")

}

func (a *App) Close() {
	if a.db != nil {
		if err := a.db.Close(); err != nil {
			log.Println("Error closing database:", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}
}
