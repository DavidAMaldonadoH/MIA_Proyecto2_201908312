package comandos

import (
	util "Proyecto2/Util"
	"fmt"
	"os"
)

type Rep struct {
	Parameters []util.Parameter
	util.Command
}

func NewRep(parameters []util.Parameter) *Rep {
	return &Rep{Parameters: parameters}
}

func (rep *Rep) Execute() interface{} {
	// parametros
	path := ""
	name := ""
	id := ""
	ruta := ""
	err := false
	// verificar parametros
	for _, p := range rep.Parameters {
		if p.Key == "path" {
			path = p.Value.(string)
		} else if p.Key == "name" {
			name = p.Value.(string)
		} else if p.Key == "id" {
			id = p.Value.(string)
		} else if p.Key == "ruta" {
			ruta = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if path == "" || name == "" || id == "" || ruta == "" {
		err = true
	}

	if !err {
		reportar(path, name, id, ruta)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando rep")
	}

	return nil
}

func reportar(path, name, id, ruta string) {
	disk, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer disk.Close()

	mbr := util.ReadMbr(disk)

	mbr.PrintInfo()
}
