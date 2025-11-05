package models
import(
	//"gorm.io/gorm"
)

type User struct{
	ID uint 	`json:"id"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string	`json:"password"`
	Todos []Todo `gorm:"foreignKey:UserID"`
}

type LoginInfo struct{
	Email string	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

type RegisterInfo struct{
	Name string		`json:"name" binding:"required" `
	Email string	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}