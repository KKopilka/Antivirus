package searchtree

import (
	"fmt"
	"kkopilka/AV/database"

	// "strconv"
	"strings"

	"github.com/beevik/prefixtree"
)

var actualTree *prefixtree.Tree

type SignTree struct {
	name        string
	offsetBegin string
	dtype       string
}

func (st *SignTree) Name() string {
	return strings.Join([]string{st.name, st.offsetBegin, st.dtype}, ":")
}

// func (st *SignTree) Offset() (int64, error) {
// 	return strconv.ParseInt(st.offsetBegin, 16, 64)
// }

func BuildSearchTree() error {
	tree := prefixtree.New()
	rows, err := database.GetConnection().Query("SELECT byte, offsetBegin, dtype FROM signatures")
	if err != nil {
		return err
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

	actualTree = tree

	return nil
}

func GetSearchTree() *prefixtree.Tree {
	return actualTree
}
