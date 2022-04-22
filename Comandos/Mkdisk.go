package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkdisk struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkdisk(parameters []util.Parameter) *Mkdisk {
	return &Mkdisk{Parameters: parameters}
}

func (mkdisk *Mkdisk) Execute() interface{} {
	fmt.Println("Comando Mkdisk")
	return nil
}
