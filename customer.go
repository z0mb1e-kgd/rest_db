package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func customerWithoutOrders(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func customerLastOrders(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
