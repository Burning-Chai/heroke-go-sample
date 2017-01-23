package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"net/http"
	"os"
	"time"
)

func main() {
	flag.Parse()

	port := os.Getenv("PORT")
	path := ""

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	router.LoadHTMLGlob(path + "templates/*.html")
	router.Static("/static", path+"static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/plan", func(c *gin.Context) {
		responseBody := []map[string]interface{}{}
		c.JSON(http.StatusOK, responseBody)
	})

	router.Run(":" + port)
}
