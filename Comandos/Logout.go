package comandos

import (
	util "Proyecto2/Util"
)

type Logout struct {
	util.Command
}

func NewLogout() *Logout {
	return &Logout{}
}

func (Logout *Logout) Execute() interface{} {
	if Is_logged_in {
		Current_user.Name = ""
		Current_user.Password = ""
		Is_logged_in = false
		util.SuccessMsg("Se ha cerrado sesión éxitosamente!")
	} else {
		util.ErrorMsg("No hay una sesión abierta, no se puede realizar logout!")
	}
	return nil
}
