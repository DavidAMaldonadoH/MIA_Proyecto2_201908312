package comandos

import (
	util "Proyecto2/Util"
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

type Pause struct {
	util.Command
}

func NewPause() *Pause {
	return &Pause{}
}

func (pause *Pause) Execute() interface{} {
	fmt.Print("> Presiona una tecla para continuar: ")
	state, err := term.MakeRaw(0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}
	defer func() {
		if err := term.Restore(0, state); err != nil {
			util.ErrorMsg(err.Error())
		}
	}()

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err != nil {
			util.ErrorMsg(err.Error())
			break
		}
		fmt.Printf("%q\r\n", r)
		if r != '\t' {
			break
		}
	}
	return nil
}
