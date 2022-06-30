package controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hplxsarthak/go-sql-project/pkg/models"
	"github.com/hplxsarthak/go-sql-project/pkg/utils"
)

var NewUser models.User // NewUser is of type User

func GetUser (c *gin.Context) {
	newUsers := models.GetUser()
	c.IndentedJSON(http.StatusOK, newUsers)
}

func GetUserById (c *gin.Context) {
	id := c.Param("id")
	ID,err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing ")
	}
	UserDetails,_ := models.GetUserById(ID)
	c.IndentedJSON(http.StatusOK, UserDetails)
}

func CreateUser (c *gin.Context) {
	User := &models.User{}
	// We have made a function in utils to parse the body into our desired schema format from json
	err := utils.ParseBody(c.Request, User)

	if err != nil {
		log.Printf("%v",err)
	}
	// User type is User so we can call the function of create User of models using User
	b := User.CreateUser()
	// res,_ := json.Marshal(b)
	c.IndentedJSON(http.StatusCreated, b)
}

func DeleteUser (c *gin.Context) {
	id := c.Param("id")
	ID,err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing ")
	}
	User := models.DeleteUser(ID)
	c.IndentedJSON(http.StatusOK, User)
}

func UpdateUser (c *gin.Context) {
	var User = &models.User{}
	er := utils.ParseBody(c.Request, User)

	if er != nil {
		log.Printf("%v",er)
	}

	id := c.Param("id")
	ID,err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing ")
	}

	UserDetails,db := models.GetUserById(ID)
	if User.Email != ""{
		UserDetails.Email = User.Email
	}
	if User.Name != ""{
		UserDetails.Name = User.Name
	}
	if User.Password != ""{
		UserDetails.Password = User.Password
	}
	if User.Role != ""{
		UserDetails.Role = User.Role
	}

	db.Save(&UserDetails)
	c.IndentedJSON(http.StatusOK, UserDetails)

}

// func SearchUser (c *gin.Context) {
// 	// Getting the search text from the url and send it to db for search
// 	s := c.Query("s")
// 	searchUsers := models.SearchUser(s)

// 	c.IndentedJSON(http.StatusOK, searchUsers)
// }

func SearchPageUser (c *gin.Context) {
	s := c.Query("s")
	page,_ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage,_ := strconv.Atoi(c.DefaultQuery("limit", "4"))

	pageUsers, total := models.SearchPageUser(s,page,perPage)
	total_page := float64(total) / float64(perPage)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": pageUsers,
		"total": total,
		"page": page,
		"total_page": math.Ceil(total_page),
	})
}