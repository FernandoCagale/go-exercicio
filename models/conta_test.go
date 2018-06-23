package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Parse OK
func TestParseOk(t *testing.T) {
	line := []string{"1", "1250"}

	var conta Conta

	err := conta.Parse(line)

	assert.Nil(t, err)
}

//Parse inválido, Id não é um valor valido
func TestParseInvalidId(t *testing.T) {
	line := []string{"1a", "1250"}

	var conta Conta

	err := conta.Parse(line)
	assert.Equal(t, err.Error(), "Dados invalido no arquivo")
}

//Parse inválido, Saldo não é um valor valido
func TestParseInvalidSaldo(t *testing.T) {
	line := []string{"1", "1250a"}

	var conta Conta

	err := conta.Parse(line)
	assert.Equal(t, err.Error(), "Dados invalido no arquivo")
}

//Quando saldo for positivo não deve aplicar multa
func TestNaoAplicaMulta(t *testing.T) {
	conta := Conta{
		ID:           1,
		SaldoInicial: 1000,
	}

	assert.Equal(t, conta.aplicaMulta(), false)
}

//Quando saldo for negativo devesse aplicar multa
func TestAplicaMultaValorNegativo(t *testing.T) {
	conta := Conta{
		ID:           1,
		SaldoInicial: -1000,
	}

	assert.Equal(t, conta.aplicaMulta(), true)
}

//Quando saldo ficar negativo com base no valor da transacao, devesse aplicar multa
func TestAplicaMultaComBaseNoValorTransacao(t *testing.T) {
	conta := Conta{
		ID:           1,
		SaldoInicial: 1000,
	}

	assert.Equal(t, conta.aplicaMultaComBaseNoValorTransacao(-2000), true)
}

//Quando saldo for negativo e houver uma transacao de credito menor que o saldo, deverá aplicar multa, porque saldo ainda ficará negativo
func TestAplicaMultaComBaseNoValorTransacaoSaldoNegativo(t *testing.T) {
	conta := Conta{
		ID:           1,
		SaldoInicial: -2000,
	}

	assert.Equal(t, conta.aplicaMultaComBaseNoValorTransacao(1000), true)
}

//Quando saldo for negativo e houver uma transacao de credito maior que o saldo, não deverá aplicar multa, porque saldo ficará positivo
func TestNaoAplicaMultaComBaseNoValorTransacaoSaldoNegativoTransacaoMaior(t *testing.T) {
	conta := Conta{
		ID:           1,
		SaldoInicial: -2000,
	}

	assert.Equal(t, conta.aplicaMultaComBaseNoValorTransacao(5000), false)
}

//Saldo conta menor que valor da transacao, adiciona transacao e multa ao saldo
func TestTransacaoAddMulta(t *testing.T) {
	conta := Conta{
		ID:           1,
		SaldoInicial: 1250,
	}

	transacao := Transacao{
		ID:    1,
		Valor: -1300,
	}

	conta.Transacao(transacao)

	assert.Equal(t, conta.SaldoInicial, -550)
}

//Saldo conta positivo, adiciona valor da transacao
func TestTransacaoAddTransacaoSemMulta(t *testing.T) {
	conta := Conta{
		ID:           1,
		SaldoInicial: 500,
	}

	transacao := Transacao{
		ID:    1,
		Valor: 1000,
	}

	conta.Transacao(transacao)

	assert.Equal(t, conta.SaldoInicial, 1500)
}

//Saldo conta zero, adiciona valor da transacao
func TestTransacaoSaldoZeroAddTransacaoSemMulta(t *testing.T) {
	conta := Conta{
		ID:           1,
		SaldoInicial: 0,
	}

	transacao := Transacao{
		ID:    1,
		Valor: 5550,
	}

	conta.Transacao(transacao)

	assert.Equal(t, conta.SaldoInicial, 5550)
}

//Saldo conta negativo, valor da transacao de credito menor que o saldo, deverá aplicar multa, porque saldo ainda ficará negativo
func TestTransacaoSaldoNegativoAddTransacaoAddMulta(t *testing.T) {
	conta := Conta{
		ID:           1,
		SaldoInicial: -1000,
	}

	transacao := Transacao{
		ID:    1,
		Valor: 500,
	}

	conta.Transacao(transacao)

	assert.Equal(t, conta.SaldoInicial, -1000)
}

//Saldo conta negativo, valor da transacao de credito menor que o saldo, deverá aplicar multa, porque saldo ainda ficará negativo
func TestTransacaoSaldoNegativoAddTransacaoNaoAddMulta(t *testing.T) {
	conta := Conta{
		ID:           1,
		SaldoInicial: -1000,
	}

	transacao := Transacao{
		ID:    1,
		Valor: 1500,
	}

	conta.Transacao(transacao)

	assert.Equal(t, conta.SaldoInicial, 500)
}
