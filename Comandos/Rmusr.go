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
	if len(rmusr.Parameters) == 1 {
		if rmusr.Parameters[0].Key == "usuario" {
			removeUser(rmusr.Parameters[0].Value.(string))
		} else {
			util.ErrorMsg("El comando rmusr recibió un parámetro no permitido!")
		}
	} else {
		util.ErrorMsg("El comando rmusr recibió mas parámetros de los permitidos!")
	}
	return nil
}

func removeUser(usuario string) {
	fmt.Println(usuario)
}
