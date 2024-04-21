package controllers

import (
	"net/http"
	"strconv"

	"example.com/mod/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateMenSneaker
func CreateMenSneaker(c *gin.Context) {
	var menSneaker models.MenSneaker

	if err := c.ShouldBindJSON(&menSneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	err := models.CreateMenSneaker(db, &menSneaker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, menSneaker)
}

// GetMenSneakers
func GetMenSneakers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	menSneakers, err := models.GetMenSneakers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menSneakers)
}

// GetMenSneakerByID
func GetMenSneakerByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	menSneaker, err := models.GetMenSneakerByID(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Men's sneaker not found"})
		return
	}

	c.JSON(http.StatusOK, menSneaker)
}

// UpdateMenSneaker
func UpdateMenSneaker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var menSneaker models.MenSneaker
	if err := db.First(&menSneaker, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Men's sneaker not found"})
		return
	}

	if err := c.ShouldBindJSON(&menSneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateMenSneaker(db, &menSneaker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menSneaker)
}

// DeleteMenSneaker
func DeleteMenSneaker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = models.DeleteMenSneaker(db, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
