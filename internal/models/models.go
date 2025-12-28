package models

import(
	"gorm.io/gorm"
"gorm.io/driver/sqlite"

)


type DBModel struct{
	OrderModel
}

func Connect()(*gorm.DB,error){
	db,err:=gorm.Open(sqlite.Open("test.db"),&gorm.Config{})
	if err!=nil{
		return nil,err
	}

	return db,nil
}