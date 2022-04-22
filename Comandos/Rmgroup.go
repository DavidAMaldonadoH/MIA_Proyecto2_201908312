package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Rmgroup struct {
	Parameters []util.Parameter
	util.Command
}

func NewRmgroup(parameters []util.Parameter) *Rmgroup {
	return &Rmgroup{Parameters: parameters}
}

func (rmgroup *Rmgroup) Execute() interface{} {
	fmt.Println("Comando Rmgroup")
	return nil
}
