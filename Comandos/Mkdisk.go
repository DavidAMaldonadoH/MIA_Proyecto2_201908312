package comandos

import (
	util "Proyecto2/Util"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

type Mkdisk struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkdisk(parameters []util.Parameter) *Mkdisk {
	return &Mkdisk{Parameters: parameters}
}

func (mkdisk *Mkdisk) Execute() interface{} {
	// parametros
	size := 0
	fit := "f"
	unit := "m"
	path := ""
	err := false
	// verificar parametros
	for _, p := range mkdisk.Parameters {
		if p.Key == "size" {
			size = p.Value.(int)
		} else if p.Key == "fit" {
			if p.Value.(string) == "ff" {
				fit = "f"
			} else if p.Value.(string) == "bf" {
				fit = "b"
			} else {
				fit = "w"
			}
		} else if p.Key == "unit" {
			unit = p.Value.(string)
		} else if p.Key == "path" {
			path = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if size == 0 {
		err = true
	}

	if path == "" {
		err = true
	}

	// si no hay parametros incorrectos crear disco
	if !err {
		t := strings.Split(path, "/")
		dir_arr := t[:len(t)-1]
		dir := strings.Join(dir_arr[:], "/")
		if len(dir) > 0 {
			if _, err2 := os.Stat(path); os.IsNotExist(err2) {
				if err3 := os.MkdirAll(dir, os.ModePerm); err3 != nil {
					util.ErrorMsg(err3.Error())
				}
			}
		}
		createDisk(size, fit, unit, path)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Mkdisk")
	}

	return nil
}

func createDisk(size int, fit string, unit string, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) { // si el disco no existe
		if (unit == "m") || (unit == "k") { // si las unidades son mega o kilobytes
			totalSize := util.CalcSize(size, unit)
			buffer := make([]byte, 1024)
			korm := 0

			if unit == "k" {
				korm = size
			} else {
				korm = size * 1024
			}

			for i := 0; i < 1024; i++ {
				buffer[i] = 0
			}

			disk, err := os.Create(path)

			if err != nil {
				util.ErrorMsg(err.Error())
			}

			for i := 0; i < korm; i++ {
				_, err := disk.Write(buffer)
				if err != nil {
					util.ErrorMsg(err.Error())
				}
			}

			pos, err := disk.Seek(0, io.SeekStart)

			if err != nil {
				util.ErrorMsg(err.Error())
			}

			// ecribir mbr
			var buffer_bytes bytes.Buffer
			mbr := util.NewMBR(int64(totalSize), byte(fit[0]))
			binary.Write(&buffer_bytes, binary.BigEndian, &mbr)

			_, err2 := disk.WriteAt(buffer_bytes.Bytes(), pos)

			if err2 != nil {
				util.ErrorMsg(err2.Error())
			}

			disk.Close()
			fmt.Printf("> mkdisk en: %s realizado con exito\n", path)

		} else {
			util.ErrorMsg("La unidad escogida no es permitida para crear un disco!")
		}
	} else {
		util.ErrorMsg("Ya existe un disco en este directorio!")
	}
}
