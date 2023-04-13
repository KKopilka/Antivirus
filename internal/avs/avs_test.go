package avs_test

import (
	"bytes"
	"fmt"
	"kkopilka/AV/internal/avs"
	"testing"
)

func Test_ReadSignatureFile(t *testing.T) {
	if s, err := avs.ReadSignatureFile(`C:\Users\yanas\AV\signatures\signature_112233445566778899AA22000B020E1900A2`); err != nil {
		t.Fatal(err)
	} else {
		t.Log(fmt.Sprintf("%+v", s))
		if s.Offset != 276 {
			t.Fatalf("expected offset 274 got %d", s.Offset)
		}

		expectedSign := []byte{
			17,
			34,
			51,
			68,
			85,
			102,
			119,
			239,
			191,
			189,
			239,
			191,
			189,
			239,
			191,
			189,
			34,
			0,
			11,
			2,
			14,
			25,
			0,
			239,
			191,
			189,
		}

		if !bytes.Equal(s.Sign, expectedSign) {
			t.Fatalf("sign check failed. got: %+v", s.Sign)
		}
	}
}
