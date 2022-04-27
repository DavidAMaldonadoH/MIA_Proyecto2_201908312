package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkfs struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkfs(parameters []util.Parameter) *Mkfs {
	return &Mkfs{Parameters: parameters}
}

func (mkfs *Mkfs) Execute() interface{} {
	// parametros
	id := ""
	tipo := "full"
	err := false
	// verificar parametros
	for _, p := range mkfs.Parameters {
		if p.Key == "id" {
			id = p.Value.(string)
		} else if p.Key == "type" {
			tipo = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if id == "" {
		err = true
	}

	if !err {
		formatPartition(id, tipo)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Fdisk")
	}
	return nil
}

func formatPartition(id, tipo string) {
	fmt.Println(id)
}
