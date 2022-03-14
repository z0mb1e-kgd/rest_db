package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func orderWithoutDetails(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func orderWithoutInvoices(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
