package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mount struct {
	Parameters []util.Parameter
	util.Command
}

func NewMount(parameters []util.Parameter) *Mount {
	return &Mount{Parameters: parameters}
}

func (mount *Mount) Execute() interface{} {
	// parametros
	path := ""
	name := ""
	err := false
	// verificar parametros
	for _, p := range mount.Parameters {
		if p.Key == "path" {
			path = p.Value.(string)
		} else if p.Key == "name" {
			name = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if path == "" || name == "" {
		err = true
	}

	if !err {
		mountPartition(path, name)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Fdisk")
	}
	return nil
}

func mountPartition(path, name string) {
	fmt.Println(path)
	fmt.Println(name)
}
