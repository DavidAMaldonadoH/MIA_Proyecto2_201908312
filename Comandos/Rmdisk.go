package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Rmdisk struct {
	Parameters []util.Parameter
	util.Command
}

func NewRmdisk(parameters []util.Parameter) *Rmdisk {
	return &Rmdisk{Parameters: parameters}
}

func (rmdisk *Rmdisk) Execute() interface{} {
	fmt.Println("Comando Rmdisk")
	return nil
}
