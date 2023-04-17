package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/beevik/prefixtree"
	_ "github.com/mattn/go-sqlite3"
)

type AVScanner struct {
	signatures map[string]Signature
	VirusStats map[string][]string
}

type Signature struct {
	id          int64
	Sign        []byte
	sha         string
	offsetBegin string // смещение в байтах от начала
	offsetEnd   string
	dtype       string
}

var errFoundSign = errors.New("найдена сигнатура. Ты бара")
var database *sql.DB

func main() {
	db, err := sql.Open("sqlite3", "C:/Users/yanas/AV/database/signatures.db")
	if err != nil {
		panic(err)
	}
	database = db
	defer db.Close()
	ReadSignatureDatabase()
}

func (signature *Signature) FindInFile(f *os.File) error {
	tree := prefixtree.New()

	for _, prefix := range []




	return nil
	// strconv.Atoi(signature.offsetBegin)

	// virSigLen := len(signature.Sign)
	// byteSlice := make([]byte, virSigLen)

	// n, err := f.ReadAt(byteSlice, signature.offsetBegin)
	// if err != nil {
	// 	return err
	// }
	// // fmt.Println("file read result:", byteSlice)

	// if n < virSigLen {
	// 	// это ошибка, потому что мы прочитали меньше в файле, чем длина сигнатуры, вируса там нет
	// 	return errors.New("прочитанно меньше байтов, чем длина сигнатуры")
	// }

	// if bytes.Equal(byteSlice, s.Sign) {
	// 	// fmt.Println("found sig", s.Sign)
	// 	return errFoundSign
	// }

	
}

// поиск и загрузка сигнатур
// func NewAVScanner() *AVScanner {
// 	a := &AVScanner{
// 		signatures: make(map[string]Signature),
// 		VirusStats: make(map[string][]string),
// 	}

// 	err := filepath.Walk("./database/signature.db", func(path string, info fs.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if info.IsDir() {
// 			return nil
// 		}

// 		s, err := ReadSignatureDatabase()
// 		if err != nil {
// 			fmt.Println("signature not loaded", path, "err", err.Error())
// 		} else {
// 			a.signatures[info.Name()] = s
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		log.Fatalf("%s", err.Error())
// 	}

// 	// fmt.Printf("Loaded signatures: %+v", a.signatures)

// 	return a
// }

// func (avs *AVScanner) ScanFile(filepath string) error {
// 	if _, ok := avs.VirusStats[filepath]; !ok {
// 		avs.VirusStats[filepath] = []string{}
// 	}

// 	// open file
// 	f, err := os.Open(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	for sigName, s := range avs.signatures {
// 		// fmt.Println("Check sign", sigName)
// 		if err := s.FindInFile(f); err != nil {

// 			// fmt.Println("FindInFile err:", err)

// 			// тут надо сделать тип ошибки отдельный для найденной сигнатуры и добавлять её в stats
// 			if err == errFoundSign {
// 				avs.VirusStats[filepath] = append(avs.VirusStats[filepath], sigName)
// 			} else if err == io.EOF {
// 				// fmt.Println("file is empty or EOF found", filepath)
// 			} else {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

func ReadSignatureDatabase() {
	result, err := database.Query("SELECT * FROM signatures")
	if err != nil {
		panic(err)
	}
	defer result.Close()
	signature := []Signature{}

	for result.Next() {
		s := Signature{}
		err := result.Scan(&s.id, &s.Sign, &s.sha, &s.offsetBegin, &s.offsetEnd, &s.dtype)

		if err != nil {
			fmt.Println(err)
			continue
		}
		signature = append(signature, s)
	}
	fmt.Println(signature)
}
