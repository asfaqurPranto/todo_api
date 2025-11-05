package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)
var DB *gorm.DB
func Connect_Mysql_Server(){

	dsn := "root:@tcp(127.0.0.1:3306)/TODO_API?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err!=nil{
		panic(err.Error())
		
	}
	DB=db

}