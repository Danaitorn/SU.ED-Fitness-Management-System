package main

import (
    "backend/config"
    "backend/routes"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time" 
    "os"
)

func main() {

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOriginFunc: func(origin string) bool { return true },
        AllowMethods:     []string{"GET", "POST", "PUT","PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
    r.Static("/uploads", "./uploads")
    loc, _ := time.LoadLocation("Asia/Bangkok")
	time.Local = loc
    config.ConnectDatabase()
    routes.SetupRoutes(r)
   	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}