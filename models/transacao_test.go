package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Parse OK
func TestParseTransacaoOk(t *testing.T) {
	line := []string{"1", "1250"}

	var transacao Transacao

	err := transacao.Parse(line)

	assert.Nil(t, err)
}

//Parse inválido, Id não é um valor valido
func TestParseTransacaoInvalidId(t *testing.T) {
	line := []string{"1a", "1250"}

	var transacao Transacao

	err := transacao.Parse(line)
	assert.Equal(t, err.Error(), "Dados invalido no arquivo")
}

//Parse inválido, Saldo não é um valor valido
func TestParseTransacaoInvalidValor(t *testing.T) {
	line := []string{"1", "1250a"}

	var transacao Transacao

	err := transacao.Parse(line)
	assert.Equal(t, err.Error(), "Dados invalido no arquivo")
}

//Quando valor for positivo e credito
func TestIsCredito(t *testing.T) {
	transacao := Transacao{
		ID:    1,
		Valor: 1000,
	}

	assert.Equal(t, transacao.IsDebito(), false)
}

//Quando valor for negativo e debito
func TestIsDebito(t *testing.T) {
	transacao := Transacao{
		ID:    1,
		Valor: -1000,
	}

	assert.Equal(t, transacao.IsDebito(), true)
}
