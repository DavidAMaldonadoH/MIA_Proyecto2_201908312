package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkfile struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkfile(parameters []util.Parameter) *Mkfile {
	return &Mkfile{Parameters: parameters}
}

func (mkfile *Mkfile) Execute() interface{} {
	// parametros
	path := ""
	r := false
	size := 0
	cont := ""
	err := false
	// verificar parametros
	for _, p := range mkfile.Parameters {
		if p.Key == "path" {
			path = p.Value.(string)
		} else if p.Key == "r" {
			r = true
		} else if p.Key == "size" {
			if p.Value.(int) < 0 {
				err = true
				break
			} else {
				size = p.Value.(int)
			}
		} else if p.Key == "cont" {
			cont = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if path == "" {
		err = true
	}

	if !err {
		makeFile(path, cont, r, size)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Mkfile")
	}

	return nil
}

func makeFile(path, cont string, r bool, size int) {
	fmt.Println(path)
}
