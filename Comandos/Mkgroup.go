package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkgroup struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkgroup(parameters []util.Parameter) *Mkgroup {
	return &Mkgroup{Parameters: parameters}
}

func (mkgroup *Mkgroup) Execute() interface{} {
	if len(mkgroup.Parameters) == 1 {
		if mkgroup.Parameters[0].Key == "name" {
			makeGroup(mkgroup.Parameters[0].Value.(string))
		} else {
			util.ErrorMsg("El comando mkgroup recibió un parámetro no permitido!")
		}
	} else {
		util.ErrorMsg("El comando mkgroup recibió mas parámetros de los permitidos!")
	}
	return nil
}

func makeGroup(name string) {
	fmt.Println(name)
}
