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

// Function to create Med
func (m *Med) CreateMed () *Med {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

// Function to get med
func GetMed() []Med {
	var Medicines []Med
	db.Find(&Medicines)
	return Medicines
}

// Function to get med by id
func GetMedById (Id int64) (*Med, *gorm.DB){
	var getMed Med
	db := db.Where("ID=?", Id).Find(&getMed)
	return &getMed, db
}

// Function to delete med by id
func DeleteMed (ID int64) Med {
	var med Med
	db.Where("ID=?", ID).Delete(med)
	return med
}