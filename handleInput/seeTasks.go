package handleinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)


func SeeTasksForOneUser(db *gorm.DB){
	scanner := bufio.NewScanner(os.Stdin)
	var tasks []Task
	var nameOfUser string
	for {
		fmt.Println("You want to see the tasks of who?")
		PrintAllClientsName(db)
		if scanner.Scan() {
				nameOfUser = strings.ToLower(scanner.Text())
				fmt.Println("Here are all the tasks remaining for ", nameOfUser)
				db.Where("name_of_user = ?", nameOfUser).Find(&tasks)
				if len(tasks) > 0{
					for i, r := range tasks{
						fmt.Printf("%d : %s\n", i, r.NameOfTask)
					}
					break
				} else{
					fmt.Println("This user has no task remaining !")
				}
				break
			}
	}
}

func DebugSeeTasksForOneUser(db *gorm.DB, nameOfClient string){
	var tasks []Task
	fmt.Println("Here are all the tasks remaining for ", nameOfClient)
	db.Where("name_of_user = ?", nameOfClient).Find(&tasks)
	if len(tasks) > 0{
		for _, r := range tasks{
			fmt.Println(r)
		}
	} else{
		fmt.Println("This user has no task remaining !")
	}
}

func PrintAllClientsName(db *gorm.DB){
	var tasks []Task
	db.Find(&tasks)
	beenWrited := make([]string,0)
	for _, r := range tasks{
		if (!checkIfAlreadyWrote(r.NameOfUser, beenWrited)){
			fmt.Printf("%s\n", r.NameOfUser)
			beenWrited = append(beenWrited, r.NameOfUser)
		}
	}
}

func checkIfAlreadyWrote(nameToWrite string, namesAlreadyWritten []string) bool{
	for _, r := range namesAlreadyWritten{
		if (nameToWrite == r){
			return true
		}
	}
	return false
}