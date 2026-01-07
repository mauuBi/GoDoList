package main

import (
	"fmt"
	"github.com/mauuBi/ToDoListGolang/handleDatabase"
)

func main(){
	db := handleDatabase.CreateAndMigrateDB()
	freshTask := handleDatabase.CreateTask(&db, 32, "Acheter du pain")
	fmt.Println(freshTask)
	oldTask := handleDatabase.GetTask(&db, 1)
	fmt.Println(oldTask)
}

