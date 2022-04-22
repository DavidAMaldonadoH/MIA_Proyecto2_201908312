package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkuser struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkuser(parameters []util.Parameter) *Mkuser {
	return &Mkuser{Parameters: parameters}
}

func (mkuser *Mkuser) Execute() interface{} {
	fmt.Println("Comando Mkuser")
	return nil
}
