package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Login struct {
	Parameters []util.Parameter
	util.Command
}

func NewLogin(parameters []util.Parameter) *Login {
	return &Login{Parameters: parameters}
}

func (login *Login) Execute() interface{} {
	// parametros
	usuario := ""
	password := ""
	id := ""
	err := false
	// verificar parametros
	for _, p := range login.Parameters {
		if p.Key == "usuario" {
			usuario = p.Value.(string)
		} else if p.Key == "password" {
			password = p.Value.(string)
		} else if p.Key == "id" {
			id = p.Value.(string)
		} else {
			err = true
		}
	}

	if usuario == "" || password == "" || id == "" {
		err = true
	}

	if !err {
		log(usuario, password, id)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Login")
	}

	return nil
}

func log(usuario, password, id string) {
	fmt.Println(usuario)
	fmt.Println(password)
	fmt.Println(id)
}
