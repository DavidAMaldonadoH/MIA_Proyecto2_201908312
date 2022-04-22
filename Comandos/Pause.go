package comandos

import (
	util "Proyecto2/Util"
	"fmt"
)

type Pause struct {
	util.Command
}

func NewPause() *Pause {
	return &Pause{}
}

func (pause *Pause) Execute() interface{} {
	fmt.Println("Comando Pause")
	return nil
}
