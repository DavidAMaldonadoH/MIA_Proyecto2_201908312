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
	if len(rmgroup.Parameters) == 1 {
		if rmgroup.Parameters[0].Key == "name" {
			removeGroup(rmgroup.Parameters[0].Value.(string))
		} else {
			util.ErrorMsg("El comando rmgroup recibió un parámetro no permitido!")
		}
	} else {
		util.ErrorMsg("El comando rmgroup recibió mas parámetros de los permitidos!")
	}
	return nil
}

func removeGroup(name string) {
	fmt.Println(name)
}
