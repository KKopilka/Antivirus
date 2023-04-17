package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/beevik/prefixtree"
	_ "github.com/mattn/go-sqlite3"
)

// type AVScanner struct {
// 	signatures map[string]Signature
// 	VirusStats map[string][]string
// }

type Signature struct {
	id          int64
	Sign        []byte
	sha         string
	offsetBegin string // смещение в байтах от начала
	offsetEnd   string
	dtype       string
}

type SignTree struct {
	// s        map[byte]*SignTree
	offsetBegin string
	dtype       string
}

func main() {
	db, err := sql.Open("sqlite3", "C:/Users/yanas/AV/database/signatures.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// ReadSignatureDatabase()
	LoadSignatures(db)

}

// func (t *Trie) Insert(s string) {
//     node := t.root
//     for _, r := range s {
//         if _, ok := node.children[r]; !ok {
//             node.children[r] = &Node{children: make(map[rune]*Node)}
//         }
//         node = node.children[r]
//     }
//     node.isEnd = true
// }

// func InsertSignature(tree *Tree, signature string, offset uint64, fileType string) {
// 	node := tree
// 	for i := 0; i < len(signature); i++ {
// 		char := signature[i]
// 		if _, ok := node.s[char]; !ok {
// 			node.s[char] = &Tree{s: make(map[byte]*Tree)}
// 		}
// 		node = node.s[char]
// 	}
// 	node.offset = offset
// 	node.fileType = fileType
// }

// func Insert(tree *prefixtree.Tree, signature string, offset uint64, fileType string) {
// 	node := tree
// 	for i := 0; i < len(signature); i++ {
// 		r := signature[i]
// 		if _, ok := node.s[r]; !ok {
// 			node.s[r] = &SignTree{s: make(map[byte]*SignTree)}
// 		}
// 		node = node.s[r]
// 	}
// 	node.offset = offset
// 	node.fileType = fileType
// }

func LoadSignatures(db *sql.DB) (*prefixtree.Tree, error) {
	tree := prefixtree.New()
	// tree := &SignTree{s: make(map[byte]*SignTree)}
	rows, err := db.Query("SELECT byte, offsetBegin, dtype FROM signatures")
	if err != nil {
		return tree, err
	}
	defer rows.Close()

	for rows.Next() {
		var signature []byte
		var offsetBegin string
		var dtype string
		if err := rows.Scan(&signature, &offsetBegin, &dtype); err != nil {
			return tree, err
		}

		sigBytes := []byte(signature)
		tree.Add(string(sigBytes), SignTree{offsetBegin: offsetBegin, dtype: dtype})
		fmt.Println(tree)
	}
	return tree, err
}

func (signature *Signature) FindInFile(f *os.File) error {
	return nil
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

// func ReadSignatureDatabase() {
// 	result, err := database.Query("SELECT * FROM signatures")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer result.Close()
// 	signature := []Signature{}

// 	for result.Next() {
// 		s := Signature{}
// 		err := result.Scan(&s.id, &s.Sign, &s.sha, &s.offsetBegin, &s.offsetEnd, &s.dtype)

// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		signature = append(signature, s)
// 	}
// 	fmt.Println(signature)
// }
