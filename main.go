package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	"backend-github-trending/helper"
	log "backend-github-trending/log"
	"backend-github-trending/repository/repo_impl"
	"backend-github-trending/router"
	"github.com/labstack/echo"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "ryan",
		Password: "postgres",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}
