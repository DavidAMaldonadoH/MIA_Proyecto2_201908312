package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Mkuser struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkuser(parameters []util.Parameter) *Mkuser {
	return &Mkuser{Parameters: parameters}
}

func (mkuser *Mkuser) Execute() interface{} {
	// parametros
	usuario := ""
	pwd := ""
	grp := ""
	err := false
	// verificar parametros
	for _, p := range mkuser.Parameters {
		if p.Key == "usuario" {
			usuario = p.Value.(string)
		} else if p.Key == "pwd" {
			pwd = p.Value.(string)
		} else if p.Key == "grp" {
			grp = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if usuario == "" || pwd == "" || grp == "" {
		err = true
	}

	if !err {
		makeUser(usuario, pwd, grp)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Mkuser")
	}

	return nil
}

func makeUser(usuario, pwd, grp string) {
	fmt.Println(usuario)
}
