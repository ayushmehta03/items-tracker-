package main

type OrderFormData struct {
	PizzaTypes []string
	PizzaSizes []string
}

type OrderRequest struct{
	Name string `form:"name" binding:"required,min=2,max=100"`
	Phone string `form:"phone" binding:"required,min=2 max=20"`
	Address string
}