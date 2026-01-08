package controllers

import (
	"gocrawler/app/db"
	"gocrawler/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	dbQuery := db.DB
	var lists []models.PageData

	result := dbQuery.Select("*").Find(&lists)

	if result.Error != nil {
		// handle error
	}

	c.JSON(http.StatusOK, lists)
}

func Create(c *gin.Context) {
	var pageData models.PageData

	if err := c.BindJSON(&pageData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&pageData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create page data"})
		return
	}

	c.JSON(http.StatusCreated, pageData)
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
