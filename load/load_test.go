package load

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Arquivo conta valido
func TestGetContasOk(t *testing.T) {
	args := [][]string{
		[]string{"1", "250"},
		[]string{"2", "-500"},
	}

	contas, err := GetContas(args)

	assert.NotNil(t, contas)
	assert.Nil(t, err)
}

//Arquivo conta com dados invalidos
func TestGetContasDadosInvalidos(t *testing.T) {
	args := [][]string{
		[]string{"1", "a250"},
		[]string{"2", "-500"},
	}

	contas, err := GetContas(args)

	assert.Nil(t, contas)
	assert.NotNil(t, err)
}

//Arquivo valido
func TestGetTransacoesOk(t *testing.T) {
	args := [][]string{
		[]string{"1", "250"},
		[]string{"2", "-500"},
	}

	transacoes, err := GetTransacoes(args)

	assert.NotNil(t, transacoes)
	assert.Nil(t, err)
}

//Arquivo com dados invalidos
func TestGetTransacoesDadosInvalidos(t *testing.T) {
	args := [][]string{
		[]string{"1", "a250"},
		[]string{"2", "-500"},
	}

	transacoes, err := GetTransacoes(args)

	assert.Nil(t, transacoes)
	assert.NotNil(t, err)
}
