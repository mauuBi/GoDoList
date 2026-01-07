package handleDatabase

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/mauuBi/ToDoListGolang/structHandler"
	"gorm.io/gorm"
)

type Task structhandler.Task

func (t Task) String() string{
	return fmt.Sprintf("Task title: %s, User id: %d, Task is done: %+v", t.NameOfTask, t.UserId, t.Done)
}

func CreateAndMigrateDB() gorm.DB{
	var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	db.AutoMigrate(&Task{})
	return *db
}

func CreateTask(db *gorm.DB, newFd uint, newNameOfTask string) Task{
	newTask := Task{
		UserId: newFd,
		NameOfTask: newNameOfTask,
		Done: false,
	}
	result := db.Create(&newTask)
	if result.Error != nil{
		panic(result.Error)
	}
	fmt.Printf("Nombre de tâches ajoutées : %d\n", result.RowsAffected)
	return newTask
}

func GetTask(db *gorm.DB, IdToFind uint) Task{
	targetTask := Task{UserId: IdToFind}
	if res := db.First(&targetTask); res.Error != nil{
		fmt.Println("JE SUIS LA")

		panic(res.Error)
	}
	return targetTask
}