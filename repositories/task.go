package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type TaskStatus int

const (
	Done TaskStatus = iota
	Pending
	Cancel
)

type TodoTask struct {
	Id     int
	Title  string
	Status TaskStatus
	UpdatedAt time.Time
}

func NewDbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create a table
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
        "title" TEXT,
        "status" TEXT,
				"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				"updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(createTableSQL)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetListTodoTasks() []TodoTask {
	db := NewDbConnection()

	results, err := db.Query(`SELECT id, title, status, updated_at FROM tasks ORDER BY created_at`)

	if err != nil {
		log.Fatal(err)
	}

	defer results.Close()

	var tasks []TodoTask
	for results.Next() {
		
		task := TodoTask{}
		err := results.Scan(&task.Id, &task.Title, &task.Status, &task.UpdatedAt)

		tasks = append(tasks, task)

		if err != nil {
			log.Fatal(err)
		}
	}

	return tasks
}

func CreateTask(title string) {
	db := NewDbConnection()
	insertQuery := `INSERT INTO tasks (title, status) VALUES (?, ?)`

	_, err := db.Exec(insertQuery, title, Pending)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Create task successfully.")
}

func UpdateTask(id int, status TaskStatus) {
	db := NewDbConnection()
	updateQuery := `UPDATE tasks SET status = (?) where id = (?)`
 
	_, err := db.Exec(updateQuery, status, id)

	if err != nil {
		log.Fatal(err)
	}
}
