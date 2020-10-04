package controller

import (
	"errors"
	"ilearn-backend/models"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddTeacher export
func AddTeacher(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var teacherCreateInput models.TeacherCreateInput
	err := c.ShouldBindJSON(&teacherCreateInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var teacher models.Teacher
	result := db.Where("email=?", teacherCreateInput.Email).First(&teacher)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db := c.MustGet("db").(*gorm.DB)
		teacher := models.Teacher{ID: uuid.New(), FullName: teacherCreateInput.FullName, Username: teacherCreateInput.Username, Mobile: teacherCreateInput.Mobile, Email: teacherCreateInput.Email}
		db.Create(&teacher)
		c.JSON(http.StatusCreated, gin.H{"message": "Teacher Successfully Signed up", "data": teacher})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The email Already Exists"})
	}
}

// GetAllTeachers export
func GetAllTeachers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var teachers []models.Teacher
	db.Find(&teachers)
	c.JSON(http.StatusOK, gin.H{"data": teachers})
}
