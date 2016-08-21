package main

import (
	"os"
	"log"
	"fmt"
	"net/http"
	_ "database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"jerod/webserver/socket"
)


// Standard handler
func handler(w http.ResponseWriter, r *http.Request) {
	println("standard handler")
	fmt.Fprintf(w, "Hello Echo!")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", standard.WrapHandler(http.HandlerFunc(handler)))
	e.GET("/socket.io/", standard.WrapHandler(socketIo()))
	e.Static("/css", "css")
	e.Static("/js", "js")
	e.Static("/static", "static")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}
	e.Run(standard.New(":" + port))
}
