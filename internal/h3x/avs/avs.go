package avs

import (
	"bytes"
	"io"
	"os"
)

type AVScanner struct {
	VirusStats map[string][]string
}

func NewAVScanner() *AVScanner {
	return &AVScanner{
		VirusStats: make(map[string][]string),
	}
}

var VirSigN1 = []byte{
	0x11,
	0x22,
	0x33,
	0x44,
	0x55,
	0x66,
	0x77,
	0x88,
	0x99,
	0xAA,
	0x22,
	0x00,
	0x0B,
	0x02,
	0x0E,
	0x19,
	0x00,
	0xA2,
}

func (avs *AVScanner) ScanFile(filepath string, sigName string, virusSignature []byte) error {
	if _, ok := avs.VirusStats[sigName]; !ok {
		avs.VirusStats[sigName] = []string{}
	}

	// open file
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	virSigLen := len(virusSignature)
	var curPos int64
	var sigFound bool

	byteSlice := make([]byte, virSigLen)
	for !sigFound {
		n, err := f.ReadAt(byteSlice, curPos)
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}

		if n < virSigLen {
			// это ошибка потому что мы прочитали меньше в файле, чем длина сигнатуры, вируса там нет.
			return nil
		}

		if bytes.Equal(byteSlice, virusSignature) {
			sigFound = true
		}

		curPos++
	}

	if sigFound {
		avs.VirusStats[sigName] = append(avs.VirusStats[sigName], filepath)
	}

	return nil
}
