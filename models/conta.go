package models

import (
	"math"
	"strconv"

	"github.com/FernandoCagale/golang-csv/lib"
)

type Conta struct {
	ID           int
	SaldoInicial int
}

const multa = -500

// Parse funcao que faz o parse do csv para a Struct
func (c *Conta) Parse(line []string) error {
	id, err := strconv.Atoi(line[0])
	if err != nil {
		return lib.ErrorInvalidFormat()
	}
	saldo, err := strconv.Atoi(line[1])
	if err != nil {
		return lib.ErrorInvalidFormat()
	}

	c.ID = id
	c.SaldoInicial = saldo
	return nil
}

// Verifica se saldo inicial da conta é negativo
func (c *Conta) aplicaMulta() bool {
	return math.Signbit(float64(c.SaldoInicial))
}

// Verifica se o saldo inicial somado ao valor da transação ficará com saldo negativo
func (c *Conta) aplicaMultaComBaseNoValorTransacao(valorTransacao int) bool {
	return math.Signbit(float64(c.SaldoInicial + valorTransacao))
}

// Adiciona valor da multa ao saldo inicial
func (c *Conta) adicionaMulta() {
	c.SaldoInicial += multa
}

// Transacao funcao responsavel por atualizar o saldo inicial da conta
func (c *Conta) Transacao(transacao Transacao) {
	if transacao.IsDebito() {
		if c.aplicaMulta() || c.aplicaMultaComBaseNoValorTransacao(transacao.Valor) {
			c.adicionaMulta()
		}
	} else if c.aplicaMultaComBaseNoValorTransacao(transacao.Valor) {
		c.adicionaMulta()
	}
	c.SaldoInicial += transacao.Valor
}
