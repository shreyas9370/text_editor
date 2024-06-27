package controllers

import (
	"net/http"
	"text-editor/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddText(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Text string `json: "text"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var currentText models.Text
		db.Last(&currentText)

		newContent := currentText.Content + input.Text
		db.Create(&models.Text{Content: newContent})
		db.Create(&models.History{Content: input.Text, Action: "add"})

		c.JSON(http.StatusOK, gin.H{"content": newContent})
	}
}

func DeleteText(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteRequest models.DeleteRequest
		if err := c.ShouldBindJSON(&deleteRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var text models.Text
		if err := db.First(&text, deleteRequest.ID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Text not found"})
			return
		}

		if err := db.Delete(&text).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete text"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Text deleted successfully"})
	}
}
