package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name    string `json:"Name"`
	Age     string `json:"Age"`
	Paterno string `json:"Paterno"`
}
