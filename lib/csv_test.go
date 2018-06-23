package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Arquivo valido
func TestNewOk(t *testing.T) {
	contas := [][]string{
		[]string{"1", "1250"},
		[]string{"2", "500"},
		[]string{"3", "0"},
		[]string{"4", "2500"},
		[]string{"5", "-1000"},
		[]string{"6", "-1000"},
	}

	lines, err := New("../csv/contas.csv")

	assert.Nil(t, err)
	assert.NotNil(t, lines)
	assert.Equal(t, contas, lines)
}

func TestFileNotFound(t *testing.T) {
	err := existFile("../csv/contass.csv")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Arquivo não encontrado: ../csv/contass.csv")
}

func TestFileOk(t *testing.T) {
	err := existFile("../csv/contas.csv")

	assert.Nil(t, err)
}
