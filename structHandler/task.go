package structhandler

import (
	"gorm.io/gorm"
)

type Task struct{
	gorm.Model
	UserId uint
	NameOfTask string
	Done bool
}

