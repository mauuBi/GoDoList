package handleInput

import (
	"strings"

	donetask "github.com/mauuBi/ToDoListGolang/handleInput/doneTask"
	"github.com/mauuBi/ToDoListGolang/handleInput/newTask"
	seetasks "github.com/mauuBi/ToDoListGolang/handleInput/seeTasks"
	"gorm.io/gorm"
)


func HandleInput(db *gorm.DB, input string, hashMap map[string]uint){
	switch strings.ToLower(input) {
	case "new task", "nt":
		newtask.HandlingNewTask(db, hashMap)
	case "see tasks", "st":
		seetasks.SeeTasksForOneUser(db, hashMap)
	case "done tasks", "dt":
		donetask.DeleteTask(db, hashMap)
	}
}