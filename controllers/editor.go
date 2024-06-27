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
