package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Fdisk struct {
	Parameters []util.Parameter
	util.Command
}

func NewFdisk(parameters []util.Parameter) *Fdisk {
	return &Fdisk{Parameters: parameters}
}

func (fdisk *Fdisk) Execute() interface{} {
	fmt.Println("Comando Fdisk")
	return nil
}
