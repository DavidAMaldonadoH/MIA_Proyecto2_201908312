package comandos

import (
	util "Proyecto2/Util"
	"io"
	"os"
	"unsafe"
)

type Fdisk struct {
	Parameters []util.Parameter
	util.Command
}

func NewFdisk(parameters []util.Parameter) *Fdisk {
	return &Fdisk{Parameters: parameters}
}

func (fdisk *Fdisk) Execute() interface{} {
	// parametros
	size := 0
	unit := "k"
	path := ""
	tipo := "p"
	fit := "w"
	name := ""
	err := false
	// verificar parametros
	for _, p := range fdisk.Parameters {
		if p.Key == "size" {
			size = p.Value.(int)
		} else if p.Key == "unit" {
			unit = p.Value.(string)
		} else if p.Key == "path" {
			path = p.Value.(string)
		} else if p.Key == "type" {
			tipo = p.Value.(string)
		} else if p.Key == "fit" {
			if p.Value.(string) == "ff" {
				fit = "f"
			} else if p.Value.(string) == "bf" {
				fit = "b"
			} else {
				fit = "w"
			}
		} else if p.Key == "name" {
			name = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if size <= 0 || path == "" || name == "" {
		err = true
	}

	if !err {
		formatDisk(size, unit, path, tipo, fit, name)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Fdisk")
	}
	return nil
}

func formatDisk(size int, unit, path, tipo, fit, name string) {
	disk, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		util.ErrorMsg(err.Error())
	}

	defer disk.Close()

	mbr := util.ReadMbr(disk)
	size_mbr := unsafe.Sizeof(mbr)

	if unit == "b" || unit == "k" || unit == "m" {
		part_size := util.CalcSize(size, unit)
		used_space := int(size_mbr)
		exists_extended := false
		extended_index := -1
		index := -1
		var names []string
		for i, part := range mbr.Partitions {
			// verificar si ya existe una partición extendida
			if part.Type == 'e' {
				exists_extended = true
			}
			// obtener el indice de donde se puede crear la partición
			if part.Start == -1 {
				index = i
				break
			} else { // guardar los nombres y aumentar el espacio usado
				names = append(names, string(part.Name[:]))
				used_space += int(part.Size)
				if part.Type == 'e' {
					extended_index = i
				}
			}
		}

		if index != -1 && tipo == "p" {
			cretePrimary(disk, used_space, part_size, index, mbr, name, path, fit, names)
		} else if index != -1 && tipo == "e" {
			if !exists_extended {
				creteExtended(disk, used_space, part_size, index, mbr, name, path, fit, names)
			} else {
				util.ErrorMsg("Ya existe una partición extendida no se puede crear otra!")
			}
		} else if exists_extended && tipo == "l" {
			creteLogic(disk, part_size, index, extended_index, mbr, name, path, fit, names)
		} else if !exists_extended && tipo == "l" {
			util.ErrorMsg("Debe existir una partición extendida para poder crear una lógica")
		} else if index == -1 && (tipo == "p" || tipo == "e") {
			msg := "Ya no se puede crear otra partición primaria o extendida el disco: " + path
			util.ErrorMsg(msg)
		} else {
			msg := "Desconocido en el fdisk de: " + path
			util.ErrorMsg(msg)
		}

	} else {
		util.ErrorMsg("Unidad inválida para fdisk!")
	}
}

func cretePrimary(disk *os.File, used_space, part_size, index int, mbr util.MBR, name, path, fit string, names []string) {
	if used_space+part_size <= int(mbr.Tamano) {
		flag_name := util.UsedName(name, names)
		if !flag_name {
			mbr.Partitions[index].Fit = byte(fit[0])
			mbr.Partitions[index].Size = int64(part_size)
			mbr.Partitions[index].Start = int64(used_space + 1)
			mbr.Partitions[index].Type = 'p'
			mbr.Partitions[index].Status = '1'
			copy(mbr.Partitions[index].Name[:], name)
			// colocar el puntero del fichero al inicio
			pos, err := disk.Seek(0, io.SeekStart)

			if err != nil {
				util.ErrorMsg(err.Error())
			}
			util.WriteMbr(disk, mbr, pos)

			util.InfoMsg("Creación de partición primaria en " + path + " con nombre: " + name)
			util.SuccessMsg("fdisk realizado con éxito!")
		} else {
			util.ErrorMsg("Ya existe una partición con ese nombre!")
		}
	} else {
		util.ErrorMsg("No hay espacio suficiente para crear la partición!")
	}
}

func creteExtended(disk *os.File, used_space, part_size, index int, mbr util.MBR, name, path, fit string, names []string) {
	if used_space+part_size <= int(mbr.Tamano) {
		flag_name := util.UsedName(name, names)
		if !flag_name {
			mbr.Partitions[index].Fit = byte(fit[0])
			mbr.Partitions[index].Size = int64(part_size)
			mbr.Partitions[index].Start = int64(used_space + 1)
			mbr.Partitions[index].Type = 'e'
			mbr.Partitions[index].Status = '1'
			copy(mbr.Partitions[index].Name[:], name)
			// colocar el puntero del fichero al inicio
			pos, err := disk.Seek(0, io.SeekStart)

			if err != nil {
				util.ErrorMsg(err.Error())
			}

			util.WriteMbr(disk, mbr, pos)

			ebr := util.NewEBR(used_space+1, 0)
			pos2, err := disk.Seek(int64(used_space+1), io.SeekStart)
			if err != nil {
				util.ErrorMsg(err.Error())
			}
			util.WriteEbr(disk, ebr, pos2)

			util.InfoMsg("Creación de partición extendida en " + path + " con nombre: " + name)
			util.SuccessMsg("fdisk realizado con éxito!")
		} else {
			util.ErrorMsg("Ya existe una partición con ese nombre!")
		}
	} else {
		util.ErrorMsg("No hay espacio suficiente para crear la partición!")
	}
}

func creteLogic(disk *os.File, part_size, index, extended_index int, mbr util.MBR, name, path, fit string, names []string) {
	extended_part := mbr.Partitions[extended_index]
	ebr := util.ReadEbr(disk, extended_part.Start)
	size_ebr := unsafe.Sizeof(ebr)
	used_space := 0
	// recorrer los mbr hasta encontrar el que termine con -1
	for ebr.Next != -1 {
		used_space += int(ebr.Size) + int(size_ebr)
		names = append(names, string(ebr.Name[:]))
		pos, err := disk.Seek(ebr.Next, io.SeekStart)
		if err != nil {
			util.ErrorMsg(err.Error())
		}
		ebr = util.ReadEbr(disk, pos)
	}

	if ebr.Next == -1 {
		if used_space+part_size+int(size_ebr) <= int(extended_part.Size) {
			flag_name := util.UsedName(name, names)
			if !flag_name {
				ebr.Fit = byte(fit[0])
				ebr.Next = extended_part.Start + int64(used_space) + int64(part_size) + int64(size_ebr)
				ebr.Size = int64(part_size)
				ebr.Status = '1'
				ebr.Start = extended_part.Start + int64(used_space)
				copy(ebr.Name[:], name)
				pos, err := disk.Seek(ebr.Start, io.SeekStart)
				if err != nil {
					util.ErrorMsg(err.Error())
				}

				util.WriteEbr(disk, ebr, pos)
				ebr2 := util.NewEBR(int(ebr.Next), 0)
				pos2, err := disk.Seek(ebr.Next, io.SeekStart)
				if err != nil {
					util.ErrorMsg(err.Error())
				}

				util.WriteEbr(disk, ebr2, pos2)
				util.InfoMsg("Creación de partición lógica en " + path + " con nombre: " + name)
				util.SuccessMsg("fdisk realizado con éxito!")
			} else {
				util.ErrorMsg("Ya existe una partición con ese nombre!")
			}
		} else {
			util.ErrorMsg("No hay espacio suficiente para crear la partición!")
		}
	}
}
