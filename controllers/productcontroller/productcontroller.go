package productcontroller

import (
	"golang-rest-api-mysql/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}
func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.First(&product, id)
	c.JSON(http.StatusOK, gin.H{"data": product})
	
}
func Create(c *gin.Context) {
	var input models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	product := models.Product{Name: input.Name, Description: input.Description}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
	
}
func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": product})
	
}
func Delete(c *gin.Context) {

	var product models.Product
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
	
}
