package main

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	customer := router.Group("/customer")
	{
		customer.GET("/without_orders")
		customer.GET("/last_orders")
	}

	order := router.Group("/order")
	{
		order.POST("/")
		order.GET("/details?order_id=:order_id")
		order.GET("/without_details")
		order.GET("/without_invoices")
	}

	invoice := router.Group("/invoice")
	{
		invoice.GET("/expired")
		invoice.GET("/wrong_date")
		invoice.GET("/overpaid")
	}

	payment := router.Group("/payment")
	{
		payment.GET("/")
		payment.GET("/details?payment_id=:payment_id")
		payment.GET("/wrong_date")
	}

	product := router.Group("/product")
	{
		product.GET("/")
		product.GET("/list")
		product.GET("/details?product_id=:product_id")
		product.GET("/high_demand")
		product.GET("/bulk")
		product.GET("/number_in_year")
	}

	category := router.Group("/category")
	{
		category.GET("/list")
		category.GET("/details?product_id=:product_id")
		category.GET("/wrong_date")
	}
	return router
}
