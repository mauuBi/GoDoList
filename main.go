package main

import (
	"fmt"

	"github.com/mauuBi/ToDoListGolang/handleDatabase"
	"github.com/mauuBi/ToDoListGolang/handleinput"
)

func main(){
	db := handleDatabase.CreateAndMigrateDB()
	fmt.Println("Welcome to your ToDoList !")
	var input string
	hashMapName := map[string]uint{}
	for {
		handleInput.WelcomeMessage(db)
		_, err := fmt.Scan(&input)
		if err != nil{
			fmt.Println(err)
			break
		}
		if input == "exit" {
            fmt.Println("Au revoir !")
            break // On sort de la boucle
        } 
		handleInput.HandleInput(db, input, hashMapName)
	}
}
