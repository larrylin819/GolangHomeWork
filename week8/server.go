package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/login", login)
	e.GET("/search/:id", getUser)
	e.Logger.Fatal(e.Start(":8888"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	if (id == "larry") {
		return c.String(http.StatusOK, "用戶名：LARRY")
	}
	return c.String(http.StatusOK, "無此用戶")
}

func login(c echo.Context) error {
	user := c.FormValue("user")
	pwd := c.FormValue("pwd")
	if(user == "larry" && pwd == "1234") {
		return c.String(http.StatusOK, "登入成功")
	}
	return c.String(http.StatusForbidden, "登入失敗")
}