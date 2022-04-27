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
	// parametros
	size := 0
	unit := "k"
	path := ""
	tipo := "p"
	fit := "w"
	name := ""
	err := false
	// verificar parametros
	for _, p := range fdisk.Parameters {
		if p.Key == "size" {
			size = p.Value.(int)
		} else if p.Key == "unit" {
			unit = p.Value.(string)
		} else if p.Key == "path" {
			path = p.Value.(string)
		} else if p.Key == "tipo" {
			tipo = p.Value.(string)
		} else if p.Key == "fit" {
			if p.Value.(string) == "ff" {
				fit = "f"
			} else if p.Value.(string) == "bf" {
				fit = "b"
			} else {
				fit = "w"
			}
		} else if p.Key == "name" {
			name = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if size == 0 || path == "" || name == "" {
		err = true
	}

	if !err {
		formatDisk(size, unit, path, tipo, fit, name)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Fdisk")
	}
	return nil
}

func formatDisk(size int, unit, path, tipo, fit, name string) {
	fmt.Println(fit)
}
