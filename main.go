package main

import (
	"GitHub-Trending/db"
	"GitHub-Trending/handler"
	"GitHub-Trending/repository/repo_impl"
	"GitHub-Trending/router"
	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "ahdayne1",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()
	e := echo.New()
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":2000"))
}
