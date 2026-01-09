package structhandler

import (
	"gorm.io/gorm"
)

type Task struct{
	gorm.Model
	NameOfUser string
	NameOfTask string
	Done bool
}

