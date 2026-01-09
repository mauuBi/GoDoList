package main

import (
	"fmt"

	"github.com/mauuBi/ToDoListGolang/handleDatabase"
	"github.com/mauuBi/ToDoListGolang/handleInput"
)

func main(){
	db := handleDatabase.CreateAndMigrateDB()
	handleDatabase.PrintAllTask(db)
	fmt.Println("Welcome to your ToDoList !")
	var input string
	for {
		handleinput.WelcomeMessage(db)
		_, err := fmt.Scan(&input)
		if err != nil{
			fmt.Println(err)
			break
		}
		if input == "exit" {
            fmt.Println("Au revoir !")
            break // On sort de la boucle
        } 
		handleinput.HandleInput(db, input)
	}
}
