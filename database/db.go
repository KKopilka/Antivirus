package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "signatures.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM signatures")
	// defer result.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("--------------------------------------------\n")
	for result.Next() {
		var id int
		var bite string
		var sha string
		var offsetBegin string
		var offsetEnd string
		var dtype string
		result.Scan(&id, &bite, &sha, &offsetBegin, &offsetEnd, &dtype)
		fmt.Printf("Сигнатура №%d\n", id)
		fmt.Printf("Байт: %s\n", bite)
		fmt.Printf("SHA-256: %s\n", sha)
		fmt.Printf("offsetBegin: %s\n", offsetBegin)
		fmt.Printf("offsetEnd: %s\n", offsetEnd)
		fmt.Printf("Тип файла: %s\n", dtype)
		fmt.Printf("--------------------------------------------\n")
	}
}
