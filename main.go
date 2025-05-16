package main

import (
	"commerce/db"
	"commerce/handlers"
	"commerce/middlewares"
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DRIVER = "sqlite3"
	DBNAME = "commerce.db"
)

func main() {
	database, err := sql.Open(DRIVER, DBNAME)
	if err != nil {
		log.Fatal(err)
	}

	db.Q = db.New(database)

	e := echo.New()
	e.Debug = true
	e.Static("/static", "dist")

	e.Use(middlewares.Logger())
	e.Use(middlewares.Session)

	e.GET("/", handlers.IndexHandler)

	e.Any("/login", handlers.LoginHandler)
	e.Any("/register", handlers.RegisterHandler)
	e.GET("/logout", handlers.LogoutHandler)

	e.Any("/create_listing", handlers.CreateListingHandler)
	e.Any("/edit_listing/:id", handlers.EditListingHandler)
	e.GET("/listing/:id", handlers.ListingHandler)
	e.GET("/close_listing/:id", handlers.CloseListingHandler)
	e.GET("/my_listings", handlers.MyListingsHandler)

	e.GET("/watchlist", handlers.WatchlistHandler)
	e.GET("/add_watchlist/:id", handlers.AddToWatchlistHandler)
	e.GET("/remove_watchlist/:id", handlers.RemoveFromWatchlistHandler)

	e.GET("/categories", handlers.CategoriesHandler)
	e.GET("/categories/:id", handlers.CategoryListingsHandler)

	e.POST("/bid", handlers.BidHandler)
	e.POST("/comment", handlers.CommentHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
