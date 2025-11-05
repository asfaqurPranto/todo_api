package models

import(
	"gorm.io/gorm"
)

type Todo struct{
	gorm.Model
	Title string	`json:"title"`
	Description string	`json:"description"`
	Completed bool		`json:"completed" gorm:"default:false"`

	UserID uint `json:"user_id"`
	User User `gorm:"foreignKey:UserID"`

}

type CreateTodo struct{
	Title string	`json:"title" binding:"required"`
	Description string	`json:"description" binding:"required"`

}
type EditTodo struct{
	Title string	`json:"title"`
	Description string	`json:"description"`
	Completed bool `json:"completed"`
}