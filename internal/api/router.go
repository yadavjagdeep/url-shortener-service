package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jagdeep/url-shortener-service/internal/handlers"
	"github.com/jagdeep/url-shortener-service/internal/repositories"
	"github.com/jagdeep/url-shortener-service/internal/services"
	"github.com/jagdeep/url-shortener-service/internal/utils"
)

func Router() *gin.Engine {
	router := gin.Default()
	return router
}

func SetupRoutes(router *gin.Engine) {

	router.GET("/ping", Ping)

	routerShortnerGroup := router.Group("/urls")
	urlShortnerService := services.NewURLShortnerService(repositories.NewURLRepository())
	urlShortnerHandler := handlers.NewURLShortnerService(urlShortnerService)
	{
		routerShortnerGroup.POST("/generate-short-url", urlShortnerHandler.CreateShortURL)
	}

	router.GET("/:shortCode", urlShortnerHandler.RedirectToLongURL)
}

func Ping(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{
		"message": "pong",
	})
}
