package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkfile struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkfile(parameters []util.Parameter) *Mkfile {
	return &Mkfile{Parameters: parameters}
}

func (mkfile *Mkfile) Execute() interface{} {
	fmt.Println("Comando Mkfile")
	return nil
}
