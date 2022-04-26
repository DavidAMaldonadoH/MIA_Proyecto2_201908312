package comandos

import (
	util "Proyecto2/Util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Rmdisk struct {
	Parameters []util.Parameter
	util.Command
}

func NewRmdisk(parameters []util.Parameter) *Rmdisk {
	return &Rmdisk{Parameters: parameters}
}

func (rmdisk *Rmdisk) Execute() interface{} {
	// verificar parametros
	if len(rmdisk.Parameters) == 1 {
		if rmdisk.Parameters[0].Key == "path" {
			removeDisk(rmdisk.Parameters[0].Value.(string))
		} else {
			util.ErrorMsg("El comando rmdisk recibió un parámetro no permitido!")
		}
	} else {
		util.ErrorMsg("El comando rmdisk recibió mas parámetros de los permitidos!")
	}
	return nil
}

func removeDisk(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) { // si el disco no existe
		util.ErrorMsg("No existe el disco ingresado!")
	} else {
		// preguntar al usuario si esta seguro de eliminar el disco
		fmt.Printf("Esta seguro de eliminar el disco: %s \n", path)
		fmt.Println("1. Si\n2. No")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> Ingrese una opción: ")
		scanner.Scan()
		input, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			util.ErrorMsg("La opción ingresada es incorrecta!")
		} else {
			if input == 1 {
				e := os.Remove(path)
				if e != nil {
					util.ErrorMsg("No se ha podido eliminar el disco!")
				} else {
					fmt.Printf("> Se ha eliminado el disco: %q con éxito!\n", path)
				}
			} else if input == 2 {
				fmt.Println("> No se eliminara el disco seleccionado.")
			} else {
				util.ErrorMsg("La opción ingresada es incorrecta!")
			}
		}
	}
}
