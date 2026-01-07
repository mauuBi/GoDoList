package seetasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mauuBi/ToDoListGolang/structHandler"
	"gorm.io/gorm"
)

type Task structhandler.Task

func (t Task) String() string{
	return fmt.Sprintf("This user need to %s", t.NameOfTask)
}

func SeeTasksForOneUser(db *gorm.DB, hashMap map[string]uint){
	scanner := bufio.NewScanner(os.Stdin)
	var tasks []Task
	var tmp string
	for {
		fmt.Println("You want to see the tasks of who?")
		if scanner.Scan() {
				tmp = scanner.Text()
				tmp = strings.ToLower(tmp)
				fmt.Println("Here are all the tasks remaining for ", tmp)
				db.Where("user_id = ?", hashMap[tmp]).Find(&tasks)
				if len(tasks) > 0{
					for _, r := range tasks{
						fmt.Println(r)
					}
					break
				} else{
					fmt.Println("This user has no task remaining !")
				}
				break
			}
	}
}

func DebugSeeTasksForOneUser(db *gorm.DB, nameOfClient string, hashMap map[string]uint){
	var tasks []Task
	fmt.Println("Here are all the tasks remaining for ", nameOfClient)
	db.Where("user_id = ?", hashMap[nameOfClient]).Find(&tasks)
	if len(tasks) > 0{
		for _, r := range tasks{
			fmt.Println(r)
		}
	} else{
		fmt.Println("This user has no task remaining !")
	}
}