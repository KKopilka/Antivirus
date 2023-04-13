package test

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "signatures.db")
	defer db.Close()
	result, _ := db.Query("SELECT * FROM signatures")
	defer result.Close()
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
