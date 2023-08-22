package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

/// SQLite
/*
* SQLite, C dili ile yazılmış, küçük, kendi kendine yeten, yüksek güvenilirlik,
* tam özellikli bir SQL veritabanı motorudur. SQLite, 1990'ların başında
* D. Richard Hipp tarafından geliştirildi ve şimdi dünyanın en popüler
* veritabanı motorlarından biridir.
*
* SQLite, çok çeşitli uygulamalarda kullanılır, bunlara dahildir:
*  - Mobil uygulamalar
*  - Masaüstü uygulamaları
*  - Web uygulamaları
*  - Embedded sistemler
*  - Bilimsel hesaplamalar
*  - Finansal uygulamalar
*  - Sağlık kayıtları
*  - Oyunlar
*
* SQLite, Mac bilgisayarlarda yerleşik olarak gelir. MacOS'ta SQLite kullanmak için,
* sqlite3 komut satırı programını kullanabilirsiniz. sqlite3 komut satırı programı,
* SQLite veritabanları oluşturmak, verileri eklemek, verileri değiştirmek ve
* verileri silmek için kullanılabilir.
*
* SQLite'yi Go diline entegre etmek için, Go'nun SQLite sürücüsünü kullanabilirsiniz.
* Go'nun SQLite sürücüsü, Go programlarından SQLite veritabanlarına erişmek ve
* bunlarla etkileşim kurmak için kullanılabilir.
 */

const (
	DB_HOST     = "localhots"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "password"
	DB_NAME     = "go_crud"
)

func OpenDataBase(driverName string, dataSourceName string) (*sql.DB, error) {
	// Remove the todo database file if exists.
	// Comment out the below line if you don't want to remove the database.
	os.Remove("./todo.db")

	// Open database connection
	db, err := sql.Open(driverName, dataSourceName)
	return db, err
}

func main() {
	db, err := OpenDataBase("sqlite3", "./todo.db")

	// Check if database connection was opened successfully
	_controllError(err)

	// close database connection before exiting program.
	defer db.Close()

	{ // Create table Block
		sqlCreateTableCode := `
	CREATE TABLE task (id INTEGER NOT NULL PRIMARY KEY, task TEXT, owner TEXT, checked INTEGER);
	DELETE FROM task;
	`
		_, err := db.Exec(sqlCreateTableCode)

		_controllError(err)
	}

	{ // Create records Block
		// Begin transaction
		tx, err := db.Begin()
		_controllError(err)

		// Prepare prepared statement that can be reused.
		stmt, err := tx.Prepare("INSERT INTO task(id, task, owner, checked) VALUES(?, ?, ?, ?)")
		_controllError(err)

		// close statement before exiting program.
		defer stmt.Close()

		task1 := todo{id: 1, task: "Learn Go", owner: "Eldi", checked: 0}

		// Insert records
		_, err = stmt.Exec(task1.id, task1.task, task1.owner, task1.checked)
		_controllError(err)

		// Commit the transaction, so that inserts are permanent.
		tx.Commit()
	}

	{ // Read records Block
		// Start reading records
		// Hint: try changing query to meet your needs.
		stmt, err := db.Prepare("SELECT id, task, owner from task where checked = ?")
		_controllError(err)
		defer stmt.Close()

		rows, err := stmt.Query(0)
		_controllError(err)
		defer rows.Close()

		for rows.Next() {
			var id int
			var task string
			var owner string
			err = rows.Scan(&id, &task, &owner)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, task, owner)
		}
		// To just check if any error was occured during iteration.
		err = rows.Err()
		_controllError(err)
	}

	{ // Update records Block
		tx, err := db.Begin()
		_controllError(err)

		stmt, err := tx.Prepare("UPDATE task SET owner = ? where id = 1")
		_controllError(err)
		defer stmt.Close()

		_, err = stmt.Exec("GoLang Community")
		_controllError(err)

		tx.Commit()
	}

	{ // Read records Block
		// Start reading records
		// Hint: try changing query to meet your needs.
		stmt, err := db.Prepare("SELECT id, task, owner from task where checked = ?")
		_controllError(err)
		defer stmt.Close()

		rows, err := stmt.Query(0)
		_controllError(err)
		defer rows.Close()

		for rows.Next() {
			var id int
			var task string
			var owner string
			err = rows.Scan(&id, &task, &owner)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, task, owner)
		}
		// To just check if any error was occured during iteration.
		err = rows.Err()
		_controllError(err)
	}

	{ // Delete records block

	}
}

func _controllError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// todo struct, represent single task
type todo struct {
	id      int
	task    string
	owner   string
	checked int
}
