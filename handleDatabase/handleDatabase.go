package handleDatabase

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/mauuBi/ToDoListGolang/structHandler"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Task structhandler.Task

func (t Task) String() string{
	return fmt.Sprintf("This user need to %s", t.NameOfTask)
}

func CreateAndMigrateDB() *gorm.DB{
	var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), 
	})
	if err != nil{
		panic(err)
	}
	db.AutoMigrate(&Task{})
	return db
}

func CreateTask(db *gorm.DB, nameOfUser string, newNameOfTask string) Task{
	newTask := Task{
		NameOfUser: nameOfUser,
		NameOfTask: newNameOfTask,
		Done: false,
	}
	result := db.Create(&newTask)
	if result.Error != nil{
		panic(result.Error)
	}
	return newTask
}

func GetTask(db *gorm.DB, IdToFind uint) []Task{
	var tasks []Task
	db.Model(&Task{}).Where("user_id = ?", IdToFind).Find(&tasks)
	return tasks
	// res := db.Where("user_id = ?", IdToFind).Find()
	// if len(tasks) < 1 {
	// 	log.Fatalln("Erreur lors de la récupération :", res.Error)
	// }
	// return tasks
}

func CheckEmptyDatabase(db *gorm.DB) bool{
	var tasks []Task
	res := db.First(&tasks, 1)
	if len(tasks) < 1 || res.Error != nil{
		return true
	}
	return false
}

func PrintAllTask(db *gorm.DB){
	var tasks []Task
	res := db.Find(&tasks)
	if res != nil{}
	for _, r := range tasks{
		fmt.Println(r)
	}
}

// Supprime toutes les tâches de l'utilisateur qui portent ce nom