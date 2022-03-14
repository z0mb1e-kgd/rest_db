package main

import (
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	customer := router.Group("/customer")
	{
		customer.GET("/without_orders", customerWithoutOrders)
		customer.GET("/last_orders", customerLastOrders)
	}

	order := router.Group("/order")
	{
		order.POST("/")
		order.GET("/details?order_id=:order_id")
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
		payment.GET("/", paymentMake)
		payment.GET("/details?payment_id=:payment_id", paymentDetails)
		payment.GET("/wrong_date", paymentWrongDate)
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
