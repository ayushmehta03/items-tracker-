package models
import(
	"gorm.io/gorm"
)
 var(
	OrderStatuses=[]string{"Order placed ","Preparing","Baking","Quality Check","Ready"}

	PizzaTypes=[]string{
		"Margherita",
		"Pepperoni",
		"Vegetarian",
		"Hawaiian",
		"Bbq Chicken",
		"Onion",
		"Cheese nd Corn",
		"FarmHouse",
		"Meal (veg)",
		"Meal (Non Veg)",
	}

	PizzaSizes=[]string{
		"Small","Medium","Large","X-large",
	}
 )

 type OrderModel struct{
	DB *gorm.DB
 }

 type Order struct{
	ID string `gorm:"primaryKey;size:14" json:"id"`
	Status string `gorm:"not null" json:"status"`
	CustomerName string `gorm:"not null" json:"customer_name"`
	Phone string `gorm:"not null" json:"phone"`
	Address string `gorm:"not null" json:"Address"`
	
 }
