package donetask

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	seetasks "github.com/mauuBi/ToDoListGolang/handleInput/seeTasks"
	structhandler "github.com/mauuBi/ToDoListGolang/structHandler"
	"gorm.io/gorm"
)

type Task structhandler.Task

func (t Task) String() string{
	return fmt.Sprintf("Task title: %s, User id: %d, Task is done: %+v", t.NameOfTask, t.UserId, t.Done)
}

func DeleteTask(db *gorm.DB, hashMap map[string]uint){
	scanner := bufio.NewScanner(os.Stdin)
	var tasks []Task
	var taskToDone string
	var nameOfClient string
	for {
		fmt.Println("You want to done the tasks of who?")
		if scanner.Scan(){
			nameOfClient = strings.ToLower(scanner.Text())
		}
		if len(nameOfClient) > 0{
			for {
				seetasks.DebugSeeTasksForOneUser(db, nameOfClient, hashMap)
				fmt.Println("You want to done which task?")
				if scanner.Scan(){
					taskToDone = strings.ToLower(scanner.Text())
					if len(taskToDone) > 0{
						db.Where("name_of_task = ? AND user_id = ?", taskToDone, hashMap[nameOfClient]).Find(&tasks)
						if len(tasks) < 1{
							fmt.Println("You never added this task:", taskToDone)
							goto end
						}
						db.Delete(&tasks)
						fmt.Println("Well done ! This task has been removed from your ToDoList:", taskToDone)
						seetasks.DebugSeeTasksForOneUser(db, nameOfClient, hashMap)
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