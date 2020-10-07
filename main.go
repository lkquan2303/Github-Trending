package main
import (
	"GitHub-Trending/db"
	"GitHub-Trending/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	sql := db.Sql{
		Host: "localhost",
		Port: 5432,
		UserName: "postgres",
		Password: "ahdayne1",
		DbName: "golang",
	}
	sql.Connect()
	defer sql.Close()
	e := echo.New()
	e.GET("/", welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)
	e.Logger.Fatal(e.Start(":2000"))
}
func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to my App")
}
