package database

import (
	"database/sql"
	"fmt"
	"kkopilka/AV/internal/avs"
	"strconv"
)

var signatures []avs.Signature
var db *sql.DB

func Open() error {
	var err error
	// подключение к БД
	db, err = sql.Open("sqlite3", "database/signatures.db")
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	// Соединения с базой нет, закрывать нечего
	if db == nil {
		return nil
	}

	return db.Close()
}

func Signatures() []avs.Signature {
	return signatures
}

func LoadSignatures() error {
	// загружаем сигнатуры из БД
	result, err := db.Query("SELECT * FROM signatures")
	if err != nil {
		return err
	}

	loadedSignatures := []avs.Signature{}

	for result.Next() {

		signature := avs.Signature{}
		var offsetBegin, offsetEnd string

		if err := result.Scan(
			&signature.Id,
			&signature.Sign,
			&signature.Sha,
			&offsetBegin,
			&offsetEnd,
			&signature.Dtype,
		); err != nil {
			return err
		}

		if offBegin, err := parseOffsetString(offsetBegin); err != nil {
			return err
		} else {
			signature.OffsetBegin = offBegin
		}
		if offEnd, err := parseOffsetString(offsetEnd); err != nil {
			return err
		} else {
			signature.OffsetEnd = offEnd
		}

		loadedSignatures = append(loadedSignatures, signature)

		fmt.Printf("Сигнатура №%d\n", signature.Id)
		fmt.Printf("Байт: %s\n", signature.Sign)
		fmt.Printf("SHA-256: %s\n", signature.Sha)
		fmt.Printf("offsetBegin: %d\n", signature.OffsetBegin)
		fmt.Printf("offsetEnd: %d\n", signature.OffsetEnd)
		fmt.Printf("Тип файла: %s\n", signature.Dtype)
		fmt.Printf("--------------------------------------------\n")
	}

	signatures = loadedSignatures

	return nil
}

func parseOffsetString(offset string) (int64, error) {
	n, err := strconv.ParseUint(offset, 16, 64)
	if err != nil {
		return 0, err
	}

	return int64(n), nil
}
