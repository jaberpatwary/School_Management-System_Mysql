package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection() {
	db := db.Connectdb()
	a.DB = db
}
func (a *App) Routes() {
	r := gin.Default()
	controller := controller.NewMangaController(a.DB)
	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":1000")
}
