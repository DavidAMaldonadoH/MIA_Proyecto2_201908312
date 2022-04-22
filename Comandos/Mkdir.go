package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkdir struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkdir(parameters []util.Parameter) *Mkdir {
	return &Mkdir{Parameters: parameters}
}

func (mkdir *Mkdir) Execute() interface{} {
	fmt.Println("Comando Mkdir")
	return nil
}
