package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mount struct {
	Parameters []util.Parameter
	util.Command
}

func NewMount(parameters []util.Parameter) *Mount {
	return &Mount{Parameters: parameters}
}

func (mount *Mount) Execute() interface{} {
	fmt.Println("Comando Mount")
	return nil
}
