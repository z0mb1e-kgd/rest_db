package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type orderEntity struct {
	gorm.Model
	date        time.Time
	customer_id int
}

func orderWithoutDetails(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func orderWithoutInvoices(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func orderCreate(c *gin.Context) {

}
