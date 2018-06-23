package lib

import (
	"encoding/csv"
	"os"
)

// New funcao responsavel por ler o arquivo CSV
func New(fileName string) (lines [][]string, err error) {
	if err = existFile(fileName); err != nil {
		return nil, err
	}

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','

	lines, err = r.ReadAll()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

// Funcao que verifica se arquivo existe
func existFile(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return ErrorFileNotFound(fileName)
	}
	return nil
}
