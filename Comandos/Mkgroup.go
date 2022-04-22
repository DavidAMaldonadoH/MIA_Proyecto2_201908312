package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkgroup struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkgroup(parameters []util.Parameter) *Mkgroup {
	return &Mkgroup{Parameters: parameters}
}

func (mkgroup *Mkgroup) Execute() interface{} {
	fmt.Println("Comando Mkgroup")
	return nil
}
