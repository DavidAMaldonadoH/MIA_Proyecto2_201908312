package comandos

import (
	util "Proyecto2/Util"
)

type Exec struct {
	Parameters []util.Parameter
	util.Command
}

func NewExec(parameters []util.Parameter) *Exec {
	return &Exec{Parameters: parameters}
}

func (exec *Exec) Execute() interface{} {
	if len(exec.Parameters) == 1 {
		if exec.Parameters[0].Key == "path" {
			return exec.Parameters[0].Value.(string)
		} else {
			util.ErrorMsg("El comando exec recibió un parámetro no permitido!")
		}
	} else {
		util.ErrorMsg("El comando exec recibió mas parámetros de los permitidos!")
	}
	return ""
}
