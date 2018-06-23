package lib

import (
	"errors"
	"fmt"
)

func ErrorInvalidFormat() error {
	return errors.New("Dados invalido no arquivo")
}

func ErrorFileNotFound(fileName string) error {
	return errors.New(fmt.Sprint("Arquivo não encontrado: ", fileName))
}

func ErrorArgsInvalid() error {
	return errors.New("Informe os arquivos de contas e transações")
}

func ErrorArgsInvalidParams() error {
	return errors.New("Informe apenas dois parametros, arquivo de contas e transações")
}
