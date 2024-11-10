package repositories

import (
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TaskStatus int

const (
	Done TaskStatus = iota
	Pending
	Cancel
)

type TodoTask struct {
	gorm.Model
	Title  string
	Status TaskStatus
	UpdatedAt time.Time
}

func NewDbConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&TodoTask{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetListTodoTasks() []TodoTask {
	db := NewDbConnection()

	// results, err := db.Get(`SELECT id, title, status, updated_at FROM tasks ORDER BY created_at`)
	var tasks []TodoTask
	db.Find(&tasks)

	return tasks
}

func CreateTask(title string) {
	db := NewDbConnection()

	db.Create(&TodoTask{
		Title: title,
		Status: Pending,
	})

}

func UpdateTask(id int, status TaskStatus) {
	db := NewDbConnection()
	var task TodoTask
	db.First(&task, id)

	task.Status = status

	db.Save(&task)
}
