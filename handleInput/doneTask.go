package handleinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mauuBi/ToDoListGolang/structHandler"
	"gorm.io/gorm"
)

type Task structhandler.Task

func DeleteTask(db *gorm.DB){ 
	scanner := bufio.NewScanner(os.Stdin)
	var taskToDone string
	var nameOfClient string
	for {
		fmt.Println("You want to done the tasks of who?")
		if scanner.Scan(){
			nameOfClient = strings.ToLower(scanner.Text())
		}
		if len(nameOfClient) > 0{
			for {
				DebugSeeTasksForOneUser(db, nameOfClient)
				fmt.Println("You want to done which task?")
				if scanner.Scan(){
					taskToDone = strings.ToLower(scanner.Text())
					if len(taskToDone) > 0{
						db.Model(&Task{}).Where("name_of_task = ? AND name_of_user = ?", taskToDone, nameOfClient).Delete(&Task{})
						fmt.Printf("Well done ! This task : %s has been removed from the ToDoList of : %s\n", taskToDone, nameOfClient)
						DebugSeeTasksForOneUser(db, nameOfClient)
						goto end
					}
				}
			}
		}
	}
	end:
}

func EraseTask(db *gorm.DB, IdToDel uint, taskToDel string) error {

	return nil
}