package models

import (
	"fmt"

	"github.com/hplxsarthak/go-sql-project/pkg/config"
	"github.com/jinzhu/gorm"
)

// This file is to define schema of the user table

var dba *gorm.DB

type User struct {
	gorm.Model
	Email 	    string 	`gorm:"unique;not null" json: "email"`
	Name 	    string 	`json: "name"`
	Password 	string 	`gorm:"not null" json: "password"`
	Role 		string 	`gorm:"not null" json: "role"`
	
}

// Initialize the database
func init() {
	// Connect with the database
	config.Connect()
	// Get the connection
	dba = config.GetDB()
	// Auto migrate with empty User
	db.AutoMigrate(&User{})
}

// Function to create User
func (u *User) CreateUser () *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

// Function to get med
func GetUser() []User {
	var Users []User
	dba.Find(&Users)
	return Users
}

// Function to get User by id
func GetUserById (Id int64) (*User, *gorm.DB){
	var getUser User
	dba := dba.Where("ID=?", Id).Find(&getUser)
	return &getUser, dba
}

// Function to delete User by id
func DeleteUser (ID int64) User {
	var User User
	dba.Where("ID=?", ID).Unscoped().Delete(&User)
	return User
}

// Function to search with Usericine name or with company name
// func SearchUser(s string) []User {
// 	var Users []User 

// 	sql := "SELECT * FROM Users"

// 	if s != "" {
// 		sql = fmt.Sprintf("%s WHERE Name Like '%%%s%%'",sql,s)
// 	}

// 	db.Raw(sql).Scan(&Users)
// 	return (Users)
// }

// Function to do pagenation and search
func SearchPageUser(s string, page int, perPage int) ([]User, int64) {
	var Users []User

	sql := "SELECT * FROM Users"

	if s != "" {
		sql = fmt.Sprintf("%s WHERE Name Like '%%%s%%'",sql,s)
	}

	
	total := db.Raw(sql).Scan(&Users).RowsAffected

	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page - 1)*perPage)

	db.Raw(sql).Scan(&Users)

	return Users,total

}

// Function to check whether a email given user is present or not
func SearchUser (email string) (*User) {
	var dbUser User

	dba.Where("Email=?", email).Find(&dbUser)

	return &dbUser
}