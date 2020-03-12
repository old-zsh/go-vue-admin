package model

import "github.com/jinzhu/gorm"

// video 模型

type Video struct {
	gorm.Model
	Title           string
	Info            string
}