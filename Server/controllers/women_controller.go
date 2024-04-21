package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"example.com/mod/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateWomenSneaker
func CreateWomenSneaker(c *gin.Context) {
	var sneaker models.Sneaker
	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	err := models.CreateWomenSneaker(db, &sneaker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sneaker)
}

// Get all sneakers
func GetWomenSneakers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	sneakers, err := models.GetWomenSneakers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sneakers)
}

// GetSneakerByID
func GetWomenSneakerByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sneaker, err := models.GetWomenSneakerByID(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sneaker not found"})
		return
	}

	c.JSON(http.StatusOK, sneaker)
}

// UpdateWomenSneaker
func UpdateWomenSneaker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var sneaker models.Sneaker
	if err := db.First(&sneaker, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sneaker not found"})
		return
	}

	if file, err := c.FormFile("image"); err == nil {

		uploadPath := "./upload/Images"

		filePath := filepath.Join(uploadPath, file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}

		sneaker.ImageURL = filePath
	}

	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateWomenSneaker(db, &sneaker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sneaker)
}

// DeleteWomenSneaker
func DeleteWomenSneaker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = models.DeleteWomenSneaker(db, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// UploadImage handles the upload of an image for a specific sneaker ID
func UploadImage(c *gin.Context) {

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
		return
	}

	uploadPath := "./upload/Images"

	err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
		return
	}

	filePath := filepath.Join(uploadPath, file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Get the sneaker ID from the form and validate it
	id := c.PostForm("id")
	sneakerID, err := strconv.Atoi(id)
	if err != nil || sneakerID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sneaker ID"})
		return
	}

	// Update the sneaker's image URL in the database
	db := c.MustGet("db").(*gorm.DB)
	var sneaker models.Sneaker
	if err := db.First(&sneaker, sneakerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sneaker not found"})
		return
	}

	// Update the sneaker's image URL
	sneaker.ImageURL = filePath
	if err := db.Save(&sneaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sneaker image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "image_url": filePath})
}
