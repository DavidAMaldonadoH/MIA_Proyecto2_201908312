package comandos

import (
	util "Proyecto2/Util"
	"fmt"
	"os"
	"unsafe"
)

type Rep struct {
	Parameters []util.Parameter
	util.Command
}

func NewRep(parameters []util.Parameter) *Rep {
	return &Rep{Parameters: parameters}
}

func (rep *Rep) Execute() interface{} {
	if len(rep.Parameters) == 1 {
		if rep.Parameters[0].Key == "path" {
			reportar(rep.Parameters[0].Value.(string))
		} else {
			fmt.Println(">> Error el comando rep recibio un parametro no permitido!")
		}
	} else {
		fmt.Println(">> Error el comando rep recibio mas parametros de los permitidos!")
	}

	return nil
}

func reportar(path string) {
	disk, err := os.OpenFile(path, os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer disk.Close()

	mbr := util.MBR{}
	var size int64 = int64(unsafe.Sizeof(mbr))
	mbr_bytes := util.ReadBytes(disk, int(size))

	if err != nil {
		fmt.Println(err.Error())
	}

	mbr = util.BytesToMBR(mbr_bytes)
	mbr.PrintInfo()
}
