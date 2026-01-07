package newtask

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mauuBi/ToDoListGolang/handleDatabase"
	handlehashmap "github.com/mauuBi/ToDoListGolang/handleHashMap"
	"gorm.io/gorm"
)

var AllId uint

func HandlingNewTask(db *gorm.DB, hashMap map[string]uint){
	var tmp string
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
				if scanner.Scan() {
						tmp = scanner.Text()
						tmp = strings.ToLower(tmp)
					}
				handlehashmap.PrintFullHashMap(hashMap)
				if (hashMap[tmp] != 0){
					handleDatabase.CreateTask(db, hashMap[tmp], nameOfTask)
					fmt.Printf("New task added : %s to : %+v\n", nameOfTask, tmp)
					goto end
				}else{
					if AllId == 0{
						AllId++
					}
					hashMap[tmp] = AllId
					AllId++
					handleDatabase.CreateTask(db, hashMap[tmp], nameOfTask)
					fmt.Printf("New task added : %s to : %+v\n", nameOfTask, tmp)
					goto end
				}
			}
		}
	}
	end:
}