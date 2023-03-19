package backend

import (
	"context"
	"database/sql"
	"log"
)

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.Database()
	log.Println("Wails application started up again")

}

func (a *App) Close() {
	if a.db != nil {
		a.db.Close()
	}
}
