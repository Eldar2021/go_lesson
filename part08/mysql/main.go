package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/// MySQL
/*
* MySQL, açık kaynaklı bir ilişkisel veritabanı yönetim sistemidir. Verileri
* tablolar halinde depolar ve SQL (Structured Query Language) kullanarak
* veritabanına erişim ve yönetim sağlar. MySQL, geniş bir kullanıcı tabanına
* sahip olup, ölçeklenebilir, güvenilir ve hızlı bir veritabanı çözümü olarak
* tercih edilir.
*
* SQLite ise hafif bir ilişkisel veritabanı yönetim sistemidir. Bir sunucuya
* ihtiyaç duymadan yerel dosya tabanlı bir veritabanı oluşturmanıza olanak tanır.
* SQLite, küçük ölçekli uygulamalar ve yerel depolama gereksinimleri için idealdir.
 */
/*
* Go ile MySQL kullanarak temel CRUD (Create, Read, Update, Delete) işlemlerini
* gerçekleştiren bir örnek:
*
* Öncelikle, MySQL veritabanına bağlanmak için go-sql-driver/mysql paketini
* kullanmamız gerekiyor. Bu paketi yüklemek için terminalde şu komutu
* kullanabilirsiniz: `go get -u github.com/go-sql-driver/mysql`
 */

func main() {
	// Veritabanına bağlanma
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testdb")
	controlError(err)
	defer db.Close()

	// Veritabanı bağlantısının açık olduğundan emin olma
	err = db.Ping()
	controlError(err)

	// CREATE - Veri ekleme
	insertStatement := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err = db.Exec(insertStatement, "Eldiiar Almazbek", "eldiiaralmazbekov@gmail.com")
	controlError(err)

	// READ - Veri sorgulama
	query := "SELECT id, name, email FROM users"
	rows, err := db.Query(query)
	controlError(err)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var email string
		err = rows.Scan(&id, &name, &email)
		controlError(err)
		fmt.Printf("ID: %d, Ad: %s, E-posta: %s\n", id, name, email)
	}

	// UPDATE - Veri güncelleme
	updateStatement := "UPDATE users SET name = ? WHERE id = ?"
	_, err = db.Exec(updateStatement, "Eldar2021", 1)
	controlError(err)

	// DELETE - Veri silme
	deleteStatement := "DELETE FROM users WHERE id = ?"
	_, err = db.Exec(deleteStatement, 1)
	controlError(err)
}

func controlError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
