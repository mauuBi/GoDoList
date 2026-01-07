package handleInput

import (
	"fmt"

	"github.com/mauuBi/ToDoListGolang/handleDatabase"
	"gorm.io/gorm"
)

func WelcomeMessage(db *gorm.DB){
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	if (handleDatabase.CheckEmptyDatabase(db)){
		fmt.Println("There is no Task in the ToDoList, enter : 'New Task' or 'NT' to create a new task or 'Exit' to leave the program.")
	} else{
		fmt.Println("Enter : 'New Task' or 'NT' to create a new task.")
		fmt.Println("Enter : 'See Tasks' or 'ST' to see the tasks remaining.")
		fmt.Println("Enter : 'Done Task' or 'DT' to mark a task as done and erase it.")
		fmt.Println("Enter : 'Exit' to leave the program.")
	}
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
}