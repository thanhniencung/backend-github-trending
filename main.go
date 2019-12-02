package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	log "backend-github-trending/log"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"os"
)

func init() {
	fmt.Println("init package main")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {
	fmt.Println("main function")
	sql := &db.Sql{
		Host: "localhost",
		Port: 5432,
		UserName: "ryan",
		Password: "postgres",
		DbName: "golang",
	}
	sql.Connect()
	defer sql.Close()

	var email string
	err := sql.Db.GetContext(context.Background(), &email, "SELECT email FROM users WHERE email=$1", "abc@gmail.com")
	if err != nil {
		log.Error(err.Error())
	}

	e := echo.New()
	e.GET("/", handler.Welcome)

	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}

