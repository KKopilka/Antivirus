package avs

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	searchtree "kkopilka/AV/internal/search-tree"
	"kkopilka/AV/internal/signature"

	"github.com/beevik/prefixtree"
)

type AVScanner struct {
	signatures map[string]signature.Signature
	VirusStats map[string][]string
}

var ErrSignatureFoundInFile = errors.New("Signature was found in file")

func FindInFile(filepath string, signTree *prefixtree.Tree) (*searchtree.SignTree, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
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
			return nil, err
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
			return nil, err
		}

		// чет найдено, пытаемся преобразовать
		if d, ok := s.(*searchtree.SignTree); ok {
			return d, ErrSignatureFoundInFile // ошиб очка
		}
		// не получилось преобразовать
		return nil, errors.New("signature data corrupted")
	}

	return nil, nil
}

var signSearchStats map[string][]string

func SearchResults() map[string][]string {
	return signSearchStats
}

func generateFilepathWalkFunction(tree *prefixtree.Tree, signSearchStats map[string][]string) func(path string, info fs.FileInfo, err error) error {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		log.Println("find in file: ", path)

		detectedSignature, err := FindInFile(path, tree)
		if err != nil {
			if err == ErrSignatureFoundInFile {
				// Сигнатура найдена в файле

				// тут тип в статистику отрицательную добавить надо
				if _, ok := signSearchStats[path]; !ok {
					signSearchStats[path] = make([]string, 0)
				}

				signSearchStats[path] = append(signSearchStats[path], detectedSignature.Name())

				return nil
			} else if err == io.EOF {
				signSearchStats[path] = nil

				return nil
			}

			return err
		}

		// файл чист
		signSearchStats[path] = nil

		return nil
	}
}

func FindSignatures(searchlocation string, tree *prefixtree.Tree) error {
	signSearchStats = make(map[string][]string)

	err := filepath.Walk(searchlocation, generateFilepathWalkFunction(tree, signSearchStats))

	if err != nil {
		return err
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

func NewAVScanner(signatures map[string]signature.Signature) *AVScanner {
	a := &AVScanner{
		signatures: signatures,
		VirusStats: make(map[string][]string),
	}

	return a
}
