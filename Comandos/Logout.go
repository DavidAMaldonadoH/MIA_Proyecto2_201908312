package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Logout struct {
	util.Command
}

func NewLogout() *Logout {
	return &Logout{}
}

func (Logout *Logout) Execute() interface{} {
	fmt.Println("Comando Logout")
	return nil
}
