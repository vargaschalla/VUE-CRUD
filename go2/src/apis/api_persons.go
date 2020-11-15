package apis

import (
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for items table
func PersonsIndex(c *gin.Context) {
	var lis []models.Person

	db, _ := c.Get("db")

	conn := db.(gorm.DB)
	// Migrate the schema
	conn.AutoMigrate(&models.Person{})

	conn.Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}

func PersonsCreate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	//var d Person
	d := models.Person{Name: c.PostForm("name"), Age: c.PostForm("age")}
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}

func PersonsGet(c *gin.Context) {

	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Person
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//db.First(&d, id)
	//c.BindJSON(&d)
	c.JSON(http.StatusOK, &d)

}
func PersonsUpdate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Person
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	d.Name = c.PostForm("name")
	d.Age = c.PostForm("age")
	c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)
}

func PersonsDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Person

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
