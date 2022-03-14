package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func invoiceExpired(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func invoiceWrongDate(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func invoiceOverpaid(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
