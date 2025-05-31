package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jagdeep/url-shortener-service/internal/models"
	"github.com/jagdeep/url-shortener-service/internal/services"
	"github.com/jagdeep/url-shortener-service/internal/utils"
)

type URLShortnerHandler struct {
	service *services.URLShortnerService
}

func NewURLShortnerService(service *services.URLShortnerService) *URLShortnerHandler {
	return &URLShortnerHandler{service: service}
}

func (h *URLShortnerHandler) CreateShortURL(c *gin.Context) {

	type createShortURLRequest struct {
		LongURL string `json:"url" binding:"required,url"`
	}

	var createShortURL createShortURLRequest

	if err := c.ShouldBindJSON(&createShortURL); err != nil {
		utils.BadRequestResponse(c, err)
		return
	}

	urlData := models.URL{
		LongURL: createShortURL.LongURL,
	}

	responseData, err := h.service.ShortenURL(&urlData)

	if err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, responseData)
}

func (h *URLShortnerHandler) RedirectToLongURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	urlData, err := h.service.GetOriginalURL(shortCode)

	if err != nil {
		utils.NotFoundErrorResponse(c, err)
		return
	}

	c.Redirect(http.StatusFound, urlData.LongURL)
}
