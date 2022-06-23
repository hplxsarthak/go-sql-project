package models

import (
	"github.com/jinzhu/gorm"
	"github.com/hplxsarthak/go-sql-project/pkg/config"
)

// This file is to define schema of the database

var db *gorm.DB

type Med struct {
	gorm.Model
	Med_Name 	string 	`gorm:""json: "med_name"`
	Comp_Name 	string 	`json: "comp_name"`
	Brand 		string 	`json: "brand"`
	Strength 	string 	`json: "strength"`
	Med_Type 	string 	`json: "med_type"`
}

// Initialize the database
func init() {
	// Connect with the database
	config.Connect()
	// Get the connection
	db = config.GetDB()
	// Auto migrate with empty Med
	db.AutoMigrate(&Med{})
}