package handleinput

import (
	"strings"

	"gorm.io/gorm"
)

func HandleInput(db *gorm.DB, input string){
	switch strings.ToLower(input) {
	case "new task", "nt":
		HandlingNewTask(db)
	case "see tasks", "st":
		SeeTasksForOneUser(db)
	case "done tasks", "dt":
		DeleteTask(db)
	}
}