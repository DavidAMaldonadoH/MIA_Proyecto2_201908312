package comandos

import (
	util "Proyecto2/Util"
	"bytes"
	"encoding/binary"
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
	// parametros
	path := ""
	name := ""
	id := ""
	ruta := ""
	err := false
	// verificar parametros
	for _, p := range rep.Parameters {
		if p.Key == "path" {
			path = p.Value.(string)
		} else if p.Key == "name" {
			name = p.Value.(string)
		} else if p.Key == "id" {
			id = p.Value.(string)
		} else if p.Key == "ruta" {
			ruta = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if path == "" || name == "" || id == "" || ruta == "" {
		err = true
	}

	if !err {
		reportar(path, name, id, ruta)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando rep")
	}

	return nil
}

func reportar(path, name, id, ruta string) {
	disk, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer disk.Close()

	// posicionarse al inicio del archivo
	disk.Seek(0, 0)
	// crear y obtener el tamano del mbr
	mbr := util.MBR{}
	size_mbr := unsafe.Sizeof(mbr)
	// leer del archivo y pasarlo al mbr
	data := util.ReadBytes(disk, int(size_mbr))
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &mbr)
	if err != nil {
		util.ErrorMsg(err.Error())
	}

	//mbr.PrintInfo()
}
