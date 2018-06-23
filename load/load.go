package load

import "github.com/FernandoCagale/golang-csv/models"

// GetContas funcao que executa o parse para as contas
func GetContas(lines [][]string) (contas []models.Conta, err error) {
	for _, line := range lines {
		var conta models.Conta
		if err := conta.Parse(line); err != nil {
			return nil, err
		}
		contas = append(contas, conta)
	}
	return contas, nil
}

// GetTransacoes funcao que executa o parse para as transacoes
func GetTransacoes(lines [][]string) (transacoes []models.Transacao, err error) {
	for _, line := range lines {
		var transacao models.Transacao
		if err := transacao.Parse(line); err != nil {
			return nil, err
		}
		transacoes = append(transacoes, transacao)
	}
	return transacoes, nil
}
