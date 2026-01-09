package handleinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mauuBi/ToDoListGolang/handleDatabase"
	"gorm.io/gorm"
)

func HandlingNewTask(db *gorm.DB){
	var nameOfUser string
	var nameOfTask string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Which task you want to add?")
		if scanner.Scan() {
				nameOfTask = scanner.Text()
			}
		if len(nameOfTask) > 0 {
			for {
				fmt.Println("For who you want to add it?")
				PrintAllClientsName(db);
				if scanner.Scan() {
						nameOfUser = strings.ToLower(scanner.Text())
						handleDatabase.CreateTask(db, nameOfUser, nameOfTask)
						fmt.Printf("New task added : %s to : %+v\n", nameOfTask, nameOfUser)
						goto end
				}
			}
		}
	}
	end:
}