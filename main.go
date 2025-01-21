package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"music-exchange/db"
	"music-exchange/handlers/auth"
	load "music-exchange/handlers/templates"
	"music-exchange/middleware"
	"music-exchange/templates"
	"music-exchange/templates/sign"
	"music-exchange/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var DB *db.Postgres

func main() {
	
	// Environment Variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}	

	// DB Inits
	DB, err = db.NewPG(context.Background(), os.Getenv("DB_URL"))

	fmt.Println(DB.Ping(context.Background()))

	var userDB db.UserDB
	userDB = db.UserDB{ Postgres: DB }
	songDB := db.SongDB{ Postgres: DB }

	if err != nil {
		panic(err)
	}

	// Server
	r := gin.Default()
	r.Static("./static", "static")

	// Error
	// 404
	r.NoRoute(func(ctx *gin.Context) { util.Render(ctx, 404, templates.NoRoutePage()) })

	// Pages

	r.GET("/", load.Index)
	r.GET("/dashboard", load.Dashboard)
	r.GET("/leaderboard", middleware.JWTAuthMiddleware, load.Leaderboard)
	r.GET("/stats", middleware.JWTAuthMiddleware, load.Stats)

	r.GET("/sign", func(ctx *gin.Context) { util.Render(ctx, 200, sign.Sign() ) })

	// Components
	r.GET("/set-sign-up", func(ctx *gin.Context) { util.Render(ctx, 200, sign.SignUp()) })
	r.GET("/set-login", func(ctx *gin.Context) { util.Render(ctx, 200, sign.Login()) })
	
	// Auth
	var oAuthHandlers auth.OAuthHandlers

	oAuthHandlers = auth.OAuthHandlers{ 
		DB: &userDB,
		SongDB: &songDB,
	}

	r.POST("/spotify-oauth", oAuthHandlers.CreateSpotifyOAuth)
	r.GET("/spotify-callback", oAuthHandlers.SpotifyCallback)

	/*
		r.POST("/apple-music-oauth", auth.AppleMusicOAuth)
		r.GET("/apple-music-callback", auth.AppleMusicCallback)
	*/

	r.Run(":8080")
}
