package main

import (
	"fmt"
	"os"

	"github.com/FernandoCagale/golang-csv/lib"
	"github.com/FernandoCagale/golang-csv/load"
	"github.com/FernandoCagale/golang-csv/models"
)

func main() {
	if err := isValidArgs(os.Args[1:]); err != nil {
		fmt.Println(err.Error())
		return
	}

	fileConta := os.Args[1]

	linesConta, err := lib.New(fileConta)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	contas, err := load.GetContas(linesConta)
	if err != nil {
		fmt.Println(err.Error(), fileConta)
		return
	}

	fileTransacao := os.Args[2]

	linesTransacao, err := lib.New(fileTransacao)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	transacoes, err := load.GetTransacoes(linesTransacao)
	if err != nil {
		fmt.Println(err.Error(), fileTransacao)
		return
	}

	for _, t := range transacoes {
		for index := range contas {
			if t.ID == contas[index].ID {
				contas[index].Transacao(t)
				break
			}
		}
	}

	for _, conta := range contas {
		fmt.Println(format(conta))
	}
}

func format(saldo models.Conta) string {
	return fmt.Sprintf("%d,%d", saldo.ID, saldo.SaldoInicial)
}

func isValidArgs(args []string) error {
	if len(args) < 2 {
		return lib.ErrorArgsInvalid()
	} else if len(args) > 2 {
		return lib.ErrorArgsInvalidParams()
	}
	return nil
}
