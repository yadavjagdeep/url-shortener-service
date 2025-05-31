package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data, "success": true})
}

func SuccessAcceptedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusAccepted, gin.H{"data": data, "success": true})
}

func BadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
}

func InternalServerErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "success": false})
}

func UnothorizedErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "success": false})
}

func NotFoundErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
}
