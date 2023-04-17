package avs

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type AVScanner struct {
	signatures map[string]Signature
	VirusStats map[string][]string
}

type Signature struct {
	id          int64
	Sign        []byte
	sha         int64
	offsetBegin int64 // смещение в байтах от начала
	offsetEnd   int64
	dtype       string
}

var errFoundSign = errors.New("найдена сигнатура. Ты бара")

func main() {
	fmt.Println(ReadSignatureDatabase())
}
func (s *Signature) FindInFile(f *os.File) error {
	virSigLen := len(s.Sign)
	byteSlice := make([]byte, virSigLen)
	n, err := f.ReadAt(byteSlice, s.offsetBegin)
	if err != nil {
		return err
	}
	// fmt.Println("file read result:", byteSlice)

	if n < virSigLen {
		// это ошибка, потому что мы прочитали меньше в файле, чем длина сигнатуры, вируса там нет
		return errors.New("прочитанно меньше байтов, чем длина сигнатуры")
	}

	if bytes.Equal(byteSlice, s.Sign) {
		// fmt.Println("found sig", s.Sign)
		return errFoundSign
	}

	return nil
}

// поиск и загрузка сигнатур
func NewAVScanner() *AVScanner {
	a := &AVScanner{
		signatures: make(map[string]Signature),
		VirusStats: make(map[string][]string),
	}

	err := filepath.Walk("../database/signatures.db", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		s, err := ReadSignatureDatabase()
		if err != nil {
			fmt.Println("signature not loaded", path, "err", err.Error())
		} else {
			a.signatures[info.Name()] = s
		}
		return nil
	})
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	// fmt.Printf("Loaded signatures: %+v", a.signatures)

	return a
}

func (avs *AVScanner) ScanFile(filepath string) error {
	if _, ok := avs.VirusStats[filepath]; !ok {
		avs.VirusStats[filepath] = []string{}
	}

	// open file
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	for sigName, s := range avs.signatures {
		// fmt.Println("Check sign", sigName)
		if err := s.FindInFile(f); err != nil {

			// fmt.Println("FindInFile err:", err)

			// тут надо сделать тип ошибки отдельный для найденной сигнатуры и добавлять её в stats
			if err == errFoundSign {
				avs.VirusStats[filepath] = append(avs.VirusStats[filepath], sigName)
			} else if err == io.EOF {
				// fmt.Println("file is empty or EOF found", filepath)
			} else {
				return err
			}
		}
	}

	return nil
}

func ReadSignatureDatabase() (Signature, error) {
	db, err := sql.Open("sqlite3", "../database/signatures.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM signatures")
	if err != nil {
		panic(err)
	}
	defer result.Close()
	// signature := []Signature{}
	s := Signature{}
	for result.Next() {

		err := result.Scan(&s.id, &s.Sign, &s.sha, &s.offsetBegin, &s.offsetEnd, &s.dtype)

		if err != nil {
			fmt.Println(err)
			continue
		}
		// signature = append(signature, s)
	}

	return s, nil

	// sign, err := ioutil.ReadFile(filepath)

	// if err != nil {
	// 	return s, err
	// }

	// // s.sign = sign
	// // fmt.Println("readed bytes:", sign)

	// var b bytes.Buffer
	// b.Write(sign)
	// scanner := bufio.NewScanner(&b)

	// scanner.Scan()
	// offstr := scanner.Text()

	// if offstr == "" {
	// 	return s, errors.New("empty offset in signature file")
	// }

	// if v, err := strconv.Atoi(offstr); err != nil {
	// 	return s, err
	// } else {
	// 	s.Offset = int64(v)
	// }

	// scanner.Scan()
	// s.Sign = scanner.Bytes()

	// return s, nil
}
