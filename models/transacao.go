package models

import (
	"math"
	"strconv"

	"github.com/FernandoCagale/golang-csv/lib"
)

type Transacao struct {
	ID    int
	Valor int
}

// Parse funcao que faz o parse do csv para a Struct
func (t *Transacao) Parse(line []string) error {
	id, err := strconv.Atoi(line[0])
	if err != nil {
		return lib.ErrorInvalidFormat()
	}
	valor, err := strconv.Atoi(line[1])
	if err != nil {
		return lib.ErrorInvalidFormat()
	}

	t.ID = id
	t.Valor = valor
	return nil
}

// IsDebito funcao retorna se a transacao e um debito
func (t *Transacao) IsDebito() bool {
	return math.Signbit(float64(t.Valor))
}
