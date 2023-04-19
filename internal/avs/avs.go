package avs

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

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
	offsetBegin int64 // смещение в байтах от начала
	offsetEnd   int64
	dtype       string
}

type SignTree struct {
	name        string
	offsetBegin string
	dtype       string
}

// var errFoundSign = errors.New("найдена сигнатура. Ты бара")

func LoadSignatures(db *sql.DB) *prefixtree.Tree {
	tree := prefixtree.New()
	rows, err := db.Query("SELECT byte, offsetBegin, dtype FROM signatures")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var signature []byte
		var offsetBegin string
		var dtype string
		err := rows.Scan(&signature, &offsetBegin, &dtype)

		if err != nil {
			fmt.Println(err)
			continue
		}
		s := &SignTree{name: string(signature), offsetBegin: offsetBegin, dtype: dtype}
		tree.Add(string(signature), s)
	}
	// tree.Output()

	return tree
}

func FindInFile(filepath string, signTree *prefixtree.Tree) error {
	f, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	defer f.Close()

	// создаем буфер на 2 символа
	byteSlice := make([]byte, 2)
	// начинаем с начала файла
	curOffset := 0
	for {
		// читаем байты в буфер
		_, err := f.ReadAt(byteSlice, int64(curOffset))
		if err != nil {
			return nil
		}

		// ищем 2 символа из буфера в префиксном дереве
		s, err := signTree.Find(string(byteSlice))
		if err != nil {
			// ошибка конец файла -> выходим из цикла
			if err == io.EOF {
				break
			}
			// ошибка префикс не найден
			if err == prefixtree.ErrPrefixNotFound {
				// смещаем указатель на 1 байт
				curOffset++
				// и возвращаемся вверх цикла
				continue
			}
			// неизвестная ошибка выходим
			return nil
		}

		// чет найдено, пытаемся преобразовать
		if d, ok := s.(*SignTree); ok {
			return d // ошиб очка
		}
		// не получилось преобразовать
		return errors.New("signature data corrupted")
	}

	return nil
}

func (avs *AVScanner) FindSignatures(path string, tree *SignTree) error {
	if len(os.Args) > 1 {
		for index, ScanDir := range os.Args {
			if index == 0 {
				continue
			} else {
				path, err := os.Stat(ScanDir)
				if err != nil {
					fmt.Printf(err.Error())
					continue
				}
				if path.IsDir() {
					err := filepath.Walk(ScanDir, func(path string, info fs.FileInfo, err error) error {
						if err != nil {
							return err
						}
						err = FindInFile(path, tree)
						if err != nil {
							return err
						}
						return nil
					})
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
	return nil
}

// func (avs *AVScanner) ScanFile(filepath string) error {
// 	if _, ok := avs.VirusStats[filepath]; !ok {
// 		avs.VirusStats[filepath] = []string{}
// 	}

// 	db, err := sql.Open("sqlite3", "C:/Users/yanas/AV/database/signatures.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	tree := LoadSignatures(db)

// 	for sigName, s := range avs.signatures {
// 		// fmt.Println("Check sign", sigName)
// 		if err := findSignatures(tree, filepath); err != nil {

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

func NewAVScanner(signatures map[string]Signature) *AVScanner {
	a := &AVScanner{
		signatures: signatures,
		VirusStats: make(map[string][]string),
	}

	return a
}
