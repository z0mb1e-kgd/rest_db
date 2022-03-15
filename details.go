package main

import "gorm.io/gorm"

type detailsEntity struct {
	gorm.Model
	order_id   uint
	product_id uint
	quantity   uint
}
