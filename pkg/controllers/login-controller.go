package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/hplxsarthak/go-sql-project/pkg/models"
	"github.com/hplxsarthak/go-sql-project/pkg/utils"
)

const SecretKey = "secret"

func Login(c *gin.Context){
	user := &models.User{}

	err := utils.ParseBody(c.Request,user)

	if err != nil {
		log.Printf("%v",err)
	}

	dbUser := models.SearchUser(user.Email)

	if dbUser.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "user not found", 
		})
		return
	}

	if dbUser.Password != user.Password {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect Password",
		})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: dbUser.Role,
		ExpiresAt: jwt.At(time.Now().Add(5 * time.Minute)) , // Expires at 5 minutes after login
	})

	token,err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Could not login",
		})
	}

	cookie := &http.Cookie{
		Name: "jwt",
		Value: token,
		Path: "/",
		Expires: time.Now().Add(2 * time.Minute),
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	c.IndentedJSON(http.StatusOK,gin.H{
		"message": "Successfully set cookie",
	})

}

func Logout(c *gin.Context) {
	cookie := &http.Cookie{
		Name: "jwt",
		Value: "",
		Path: "/",
		Expires: time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Successfully removed cookie",
	})

}