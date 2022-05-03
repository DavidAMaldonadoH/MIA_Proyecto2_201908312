package comandos

import (
	util "Proyecto2/Util"
	"fmt"
	"os"
	"strconv"
)

var counter int = 1
var mountedDisks map[string]int = make(map[string]int)
var mountedPartitions []util.MountedPartition

type Mount struct {
	Parameters []util.Parameter
	util.Command
}

func NewMount(parameters []util.Parameter) *Mount {
	return &Mount{Parameters: parameters}
}

func (mount *Mount) Execute() interface{} {
	// parametros
	path := ""
	name := ""
	err := false
	// verificar parametros
	for _, p := range mount.Parameters {
		if p.Key == "path" {
			path = p.Value.(string)
		} else if p.Key == "name" {
			name = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if path == "" || name == "" {
		err = true
	}

	if !err {
		mountPartition(path, name)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Fdisk")
	}
	return nil
}

func mountPartition(path, name string) {
	if _, err := os.Stat(path); os.IsNotExist(err) { // si el disco no existe
		util.ErrorMsg("No existe el disco ingresado!")
	} else {
		var identificador string
		var patition_counter int = 0
		letters := [26]string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
			"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
			"w", "x", "y", "z"}
		prefix := "12"
		_, found := mountedDisks[path]
		if !found {
			mountedDisks[path] = counter
			counter++
		}

		for _, p := range mountedPartitions {
			if path == p.Path {
				if p.Name != name {
					patition_counter += 1
				}
			}
		}
		number, found := mountedDisks[path]

		if found {
			identificador = prefix + strconv.Itoa(number) + letters[patition_counter]
		}

		fmt.Println(identificador)
		var mounted_part util.MountedPartition
		mounted_part.Id = identificador
		mounted_part.Name = name
		mounted_part.Path = path
		exists_part := false
		disk, err := os.OpenFile(path, os.O_RDWR, 0777)
		if err != nil {
			util.ErrorMsg(err.Error())
		}

		mbr := util.ReadMbr(disk)

		for _, p := range mbr.Partitions {
			var name_bytes [16]byte
			copy(name_bytes[:], name)
			if string(name_bytes[:]) == string(p.Name[:]) {
				exists_part = true
				mounted_part.P = p
				break
			}
			if p.Type == 'e' {
				ebr := util.ReadEbr(disk, p.Start)
				for ebr.Next != -1 {
					if string(name_bytes[:]) == string(ebr.Name[:]) {
						part := util.Partition{Name: ebr.Name, Status: ebr.Status, Type: 'l', Fit: ebr.Fit, Start: ebr.Start, Size: ebr.Size}
						mounted_part.P = part
						exists_part = true
						break
					}
					pos, err := disk.Seek(ebr.Next, 0)
					if err != nil {
						util.ErrorMsg(err.Error())
					}
					ebr = util.ReadEbr(disk, pos)
				}
			}
		}

		disk.Close()

		if exists_part {
			mountedPartitions = append(mountedPartitions, mounted_part)
			util.SuccessMsg("Partición montada con éxito con id: " + identificador)
		} else {
			util.InfoMsg("No existe una partición con ese nombre en el disco!")
			util.ErrorMsg("No se pudo montar la partición")
		}

	}
}
