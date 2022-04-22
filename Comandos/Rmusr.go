package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Rmusr struct {
	Parameters []util.Parameter
	util.Command
}

func NewRmusr(parameters []util.Parameter) *Rmusr {
	return &Rmusr{Parameters: parameters}
}

func (rmusr *Rmusr) Execute() interface{} {
	fmt.Println("Comando Rmusr")
	return nil
}
