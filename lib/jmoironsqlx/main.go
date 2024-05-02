package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

var schema = `
CREATE TABLE JmoironUsers (
    id INT AUTO_INCREMENT,
    name VARCHAR(255),
    email VARCHAR(255),
    PRIMARY KEY (id)
);`

func main() {

	// 데이터베이스 연결 설정
	db, err := sqlx.Connect("mysql", "root:Elvls0414!!@tcp(3.35.37.105:3306)/_test")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 테이블 생성
	// Result := db.MustExec(schema)
	// fmt.Printf("%+v\n", Result)

	// Create (C)
	res := createUser(db, "John Doe", "john@example.com")
	n1, err1 := res.LastInsertId()
	n2, err2 := res.RowsAffected()
	fmt.Println("createUser : ", n1, n2, err1, err2)

	// Read (R)
	users, err := getAllUsers(db)
	fmt.Println("Users:", users)
	if err != nil {
		fmt.Println("Users:", err)
	}

	// Update (U)
	updateUserEmail(db, 1, "newjohn@example.com")
	updatedUser, err := getUser(db, 1)
	fmt.Println("Updated JmoironUsers:", updatedUser)
	if err != nil {
		fmt.Println("Updated JmoironUsers:", err)
	}

	// Delete (D)
	deleteUser(db, 1)
	remainingUsers, err := getAllUsers(db)
	fmt.Println("Remaining JmoironUsers:", remainingUsers)
	if err != nil {
		fmt.Println("Remaining JmoironUsers:", err)
	}
}

func createUser(db *sqlx.DB, name string, email string) sql.Result {
	query := "INSERT INTO JmoironUsers (name, email) VALUES (?, ?)"
	result := db.MustExec(query, name, email)
	return result
}

func getAllUsers(db *sqlx.DB) ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT * FROM JmoironUsers")
	return users, err
}

func getUser(db *sqlx.DB, id int) (User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM JmoironUsers WHERE id = ?", id)
	return user, err
}

func updateUserEmail(db *sqlx.DB, id int, newEmail string) sql.Result {
	query := "UPDATE JmoironUsers SET email = ? WHERE id = ?"
	result := db.MustExec(query, newEmail, id)
	return result
}

func deleteUser(db *sqlx.DB, id int) sql.Result {
	query := "DELETE FROM JmoironUsers WHERE id = ?"
	result := db.MustExec(query, id)
	return result
}
