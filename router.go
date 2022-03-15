package main

import (
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	customer := router.Group("/customer")
	{
		customer.GET("/list", customerList)
		customer.GET("/without_orders", customerWithoutOrders)
		customer.GET("/last_orders", customerLastOrders)
	}

	order := router.Group("/order")
	{
		order.POST("/", orderCreate)
		order.GET("/details")
		order.GET("/without_details", orderWithoutDetails)
		order.GET("/without_invoices", orderWithoutInvoices)
	}

	invoice := router.Group("/invoice")
	{
		invoice.GET("/expired", invoiceExpired)
		invoice.GET("/wrong_date", invoiceWrongDate)
		invoice.GET("/overpaid", invoiceOverpaid)
	}

	payment := router.Group("/payment")
	{
		payment.GET("/", paymentCreate)
		payment.GET("/details", paymentDetails)
		payment.GET("/wrong_date", paymentWrongDate)
	}

	product := router.Group("/product")
	{
		product.GET("/")
		product.GET("/list")
		product.GET("/details")
		product.GET("/high_demand")
		product.GET("/bulk")
		product.GET("/number_in_year")
	}

	category := router.Group("/category")
	{
		category.GET("/list")
		category.GET("/details")
		category.GET("/wrong_date")
	}
	return router
}
