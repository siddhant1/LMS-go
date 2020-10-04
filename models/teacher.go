package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID       uuid.UUID `json:"ID" gorm:"primary_key"`
	FullName string    `json:"fullname"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Mobile   string    `json:"mobile"`
	Email    string    `json:"email"`
}

type TeacherCreateInput struct {
	FullName string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
