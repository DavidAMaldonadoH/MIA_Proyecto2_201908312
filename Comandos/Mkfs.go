package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkfs struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkfs(parameters []util.Parameter) *Mkfs {
	return &Mkfs{Parameters: parameters}
}

func (mkfs *Mkfs) Execute() interface{} {
	fmt.Println("Comando Mkfs")
	return nil
}
