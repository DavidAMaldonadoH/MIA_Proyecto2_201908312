package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkdir struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkdir(parameters []util.Parameter) *Mkdir {
	return &Mkdir{Parameters: parameters}
}

func (mkdir *Mkdir) Execute() interface{} {
	// parametros
	path := ""
	p := false
	err := false
	// verificar parametros
	for _, param := range mkdir.Parameters {
		if param.Key == "path" {
			path = param.Value.(string)
		} else if param.Key == "p" {
			p = true
		} else {
			err = true
			break
		}
	}

	if path == "" {
		err = true
	}

	if !err {
		makeDir(path, p)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Mkdir")
	}

	return nil
}

func makeDir(path string, p bool) {
	fmt.Println(path)
	fmt.Println(p)
}
