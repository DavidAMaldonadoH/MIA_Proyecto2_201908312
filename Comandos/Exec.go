package comandos

import (
	util "Proyecto2/Util"
	"fmt"
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
			fmt.Println(">> Error el comando exec recibio un parametro no permitido!")
		}
	} else {
		fmt.Println(">> Error el comando exec recibio mas parametros de los permitidos!")
	}
	return ""
}
