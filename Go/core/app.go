package core

import (
	"fmt"
	"os"
	"widatech_interview/golang/api"
	"widatech_interview/golang/config"

	"github.com/labstack/echo/v4"
)

type App struct {
	Config config.AppConfig
	Server *echo.Echo
	Db     *Db
}

func NewApp(configPath string) *App {
	app := &App{}

	return app
}

func (a *App) Boot() {
	a.initConfig()
	a.initDatabase(a.Config.DatabaseConfig)
}

func (a *App) Serve() {
	a.Server = api.Create()
	api.Register(a.Server, a.Db.Connection)
	api.Start(a.Server, a.Config.ServerConfig)
}

func (a *App) initConfig() {
	conf, err := config.LoadConfig(".")

	if err != nil {
		fmt.Printf("[ERR] [APP] Failed Load Config with error : %s", err.Error())

		os.Exit(0)
		return
	}

	a.Config = conf
}

func (a *App) initDatabase(conf *config.DatabaseConfig) {
	a.Db = NewDB(conf)

	if err := a.Db.MakeConnection(); err != nil {
		fmt.Printf("[ERR] [APP] Failed Make Connection : %s", err.Error())

		os.Exit(0)
		return
	}
}
