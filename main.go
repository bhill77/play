package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	ID   int
	Name string
}

func userHandler(c echo.Context) error {
// tes
	// user := User{ID: id, Name: "Poby"}
	// if len(id) <= 3 {
	// 	return c.JSON(http.StatusNotFound, "tidak ada")
	// }
	data := map[string]interface{}{}
	c.Bind(&data)

	if data["nama"] == "" {
		return c.String(403, "nama is required")
	}
	return c.JSON(http.StatusOK, data)
}

func main() {
	e := echo.New()

	e.POST("/user", userHandler)
	e.Start(":8000")
}
