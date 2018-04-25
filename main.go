package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"mdNote/handler"
	"mdNote/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

var (
	PORT         string
	DATABASE_URL string
)

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	if URL, exist := os.LookupEnv("DATABASE_URL"); !exist {
		panic("Cannot found DATABASE_URL in environment variable")
	} else {
		if URL, err := pq.ParseURL(URL); err != nil {
			panic(err)
		} else {
			DATABASE_URL = strings.Join([]string{URL, "sslmode=require"}, " ")
		}
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{echo.OPTIONS, echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.HEAD, echo.PATCH},
	}))

	h := new(handler.Handler)
	if db, err := model.OpenAndCreate("postgres", DATABASE_URL); err != nil {
		panic(fmt.Errorf("OpenAndCreate Error : %v\n", err))
	} else {
		h.DB = db
		defer h.DB.Close()
	}
	h.SecretKey = handler.GenerateRandomKey(64)

	e.Static("/static", filepath.Join("MdNote", "dist", "static"))

	e.GET("/", func(c echo.Context) error {
		return c.File(filepath.Join("MdNote", "dist", "index.html"))
	})

	e.GET("/note", h.GetNote, h.AuthRequired())
	e.GET("/note/list", h.GetNotes, h.AuthRequired())

	e.PUT("/note", h.SaveNote, h.AuthRequired())
	e.DELETE("/note", h.DeleteNote, h.AuthRequired())

	e.POST("/auth", h.Login)
	e.POST("/signup", h.SignUp)

	e.Logger.Fatal(e.Start(":" + PORT))
}
