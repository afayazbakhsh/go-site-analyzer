package controllers

import (
	"errors"
	"gocrawler/app/db"
	"gocrawler/app/httpserver/requests"
	"gocrawler/app/httpserver/response"
	"gocrawler/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	request := requests.PageDataIndexRequest{}

 	if err := c.ShouldBind(&request); err != nil {
		c.JSON(422, gin.H{"errors": err.Error()})

		return
	}

	dbQuery := db.DB
	baseQuery := dbQuery.Select("*").Table("page_data")

	if request.Title != nil {
		baseQuery.Where("title like ?", "%"+*request.Title+"%")
	}

	if request.URL != nil {
		baseQuery.Where("url", request.URL)
	}

	var lists []response.PageDataIndexResponse

	result := baseQuery.Find(&lists)

	if result.Error != nil {

	}

	c.JSON(http.StatusOK, lists)
}

func Show(c *gin.Context) {

	id := c.Param("id")

	var pageData models.PageData

	if err := db.DB.First(&pageData, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "DB error"})
		}

		return
	}

	c.JSON(http.StatusOK, pageData)
}

func Create(c *gin.Context) {

	var request requests.PageDataCreateRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create page data"})
		return
	}

	c.JSON(http.StatusCreated, request)
}

func Update(c *gin.Context) {
	var request requests.PageDataUpdateRequest
	id := c.Param("id")

	// Parse request body
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	var page models.PageData

	if err := db.DB.Where("id = ?", id).First(&page).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	if request.URL != nil {
		page.URL = *request.URL
	}

	if request.Title != nil {
		page.Title = *request.Title
	}

	if err := db.DB.Save(&page).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update page data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page updated successfully", "page": page})
}

func Delete(c *gin.Context) {

	id := c.Param("id")
	var pageData models.PageData

	if err := db.DB.Delete(&pageData, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
