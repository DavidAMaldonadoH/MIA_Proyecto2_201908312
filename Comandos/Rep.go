package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Rep struct {
	Parameters []util.Parameter
	util.Command
}

func NewRep(parameters []util.Parameter) *Rep {
	return &Rep{Parameters: parameters}
}

func (rep *Rep) Execute() interface{} {
	fmt.Println("Comando Rep")
	return nil
}
