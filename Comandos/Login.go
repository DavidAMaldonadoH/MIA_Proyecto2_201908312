package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Login struct {
	Parameters []util.Parameter
	util.Command
}

func NewLogin(parameters []util.Parameter) *Login {
	return &Login{Parameters: parameters}
}

func (login *Login) Execute() interface{} {
	fmt.Println("Comando Login")
	return nil
}
