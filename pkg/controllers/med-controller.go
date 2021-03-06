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

var NewMed models.Med // NewMed is of type Med

func GetMed (c *gin.Context) {
	newMeds := models.GetMed()
	c.IndentedJSON(http.StatusOK, newMeds)
}

func GetMedById (c *gin.Context) {
	id := c.Param("id")
	ID,err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing ")
	}
	medDetails,_ := models.GetMedById(ID)
	c.IndentedJSON(http.StatusOK, medDetails)
}

func CreateMed (c *gin.Context) {
	Medicine := &models.Med{}
	// We have made a function in utils to parse the body into our desired schema format from json
	err := utils.ParseBody(c.Request, Medicine)

	if err != nil {
		log.Printf("%v",err)
	}
	// Medicine type is med so we can call the function of create med of models using medicine
	b := Medicine.CreateMed()
	// res,_ := json.Marshal(b)
	c.IndentedJSON(http.StatusCreated, b)
}

func DeleteMed (c *gin.Context) {
	id := c.Param("id")
	ID,err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing ")
	}
	med := models.DeleteMed(ID)
	c.IndentedJSON(http.StatusOK, med)
}

func UpdateMed (c *gin.Context) {
	var med = &models.Med{}
	er := utils.ParseBody(c.Request, med)

	if er != nil {
		log.Printf("%v",er)
	}


	id := c.Param("id")
	ID,err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error in parsing ")
	}

	medDetails,db := models.GetMedById(ID)
	if med.Med_Name != ""{
		medDetails.Med_Name = med.Med_Name
	}
	if med.Comp_Name != ""{
		medDetails.Comp_Name = med.Comp_Name
	}
	if med.Brand != ""{
		medDetails.Brand = med.Brand
	}
	if med.Strength != ""{
		medDetails.Strength = med.Strength
	}
	if med.Med_Type != ""{
		medDetails.Med_Type = med.Med_Type
	}

	db.Save(&medDetails)
	c.IndentedJSON(http.StatusOK, medDetails)

}

// func SearchMed (c *gin.Context) {
// 	// Getting the search text from the url and send it to db for search
// 	s := c.Query("s")
// 	searchMeds := models.SearchMed(s)

// 	c.IndentedJSON(http.StatusOK, searchMeds)
// }

func SearchPageMed (c *gin.Context) {
	s := c.Query("s")
	page,_ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage,_ := strconv.Atoi(c.DefaultQuery("limit", "4"))

	pageMeds, total := models.SearchPageMed(s,page,perPage)
	total_page := float64(total) / float64(perPage)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": pageMeds,
		"total": total,
		"page": page,
		"total_page": math.Ceil(total_page),
	})
}