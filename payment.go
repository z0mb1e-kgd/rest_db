package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func paymentCreate(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func paymentDetails(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func paymentWrongDate(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
