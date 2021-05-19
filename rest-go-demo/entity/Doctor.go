package entity

import "github.com/jinzhu/gorm"

//Person object for REST(CRUD)
type Doctor struct {
	gorm.Model
	
	Name string `json:"Name"`
	Age       int    `json:"age"`
	Department string `json:"department"`
	Mail string `json:"mail"`
	Phone string `json:"phone"`
}
