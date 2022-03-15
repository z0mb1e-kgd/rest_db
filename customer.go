package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerEntity struct {
	// gorm.Model

	ID      uint   `gorm:"primarykey"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

// FOR TESTING PURPOSES {
func customerCreate(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}

// }

// FOR TESTING PURPOSES {
func customerList(c *gin.Context) {
	var customers []customerEntity
	db.Table("customer").Find(&customers)
	c.JSON(http.StatusOK, &customers)
}

// }

func customerWithoutOrders(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func customerLastOrders(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
