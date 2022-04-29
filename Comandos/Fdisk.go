package comandos

import (
	util "Proyecto2/Util"
	"bytes"
	"encoding/binary"
	"fmt"
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

	if unit == "b" || unit == "k" || unit == "m" {
		part_size := util.CalcSize(size, unit)
		used_space := int(size_mbr)
		exists_extended := false
		index := -1
		var names []string
		for i, part := range mbr.Partitions {
			// verificar si ya existe una particion extendida
			if part.Type == 'e' {
				exists_extended = true
			}
			// obtener el indice de donde se puede crear la particion
			if part.Start == -1 {
				index = i
				break
			} else { // guardar los nombres y aumentar el espacio usado
				names = append(names, string(part.Name[:]))
				used_space += int(part.Size)
			}

		}

		if index != -1 && tipo == "p" {
			if used_space+part_size <= int(mbr.Tamano) {
				flag_name := util.UsedName(name, names)
				if !flag_name {
					mbr.Partitions[index].Fit = byte(fit[0])
					mbr.Partitions[index].Size = int64(size)
					mbr.Partitions[index].Start = int64(used_space + 1)
					mbr.Partitions[index].Type = 'p'
					mbr.Partitions[index].Status = '1'
					copy(mbr.Partitions[index].Name[:], name)
					// colocar el puntero del fichero al inicio
					pos, err := disk.Seek(0, io.SeekStart)

					if err != nil {
						util.ErrorMsg(err.Error())
					}

					// ecribir mbr
					var buffer_bytes bytes.Buffer
					binary.Write(&buffer_bytes, binary.BigEndian, &mbr)

					_, err2 := disk.WriteAt(buffer_bytes.Bytes(), pos)

					if err2 != nil {
						util.ErrorMsg(err.Error())
					}

					fmt.Println("> fdisk realizado con exito!")
				} else {
					util.ErrorMsg("Ya existe una particion con ese nombre!")
				}
			} else {
				util.ErrorMsg("No hay espacio suficiente para crear la partición!")
			}
		} else if index != -1 && tipo == "e" {
			if !exists_extended {
				if used_space+part_size <= int(mbr.Tamano) {
					flag_name := util.UsedName(name, names)
					if !flag_name {
						mbr.Partitions[index].Fit = byte(fit[0])
						mbr.Partitions[index].Size = int64(size)
						mbr.Partitions[index].Start = int64(used_space + 1)
						mbr.Partitions[index].Type = 'e'
						mbr.Partitions[index].Status = '1'
						copy(mbr.Partitions[index].Name[:], name)
						// colocar el puntero del fichero al inicio
						pos, err := disk.Seek(0, io.SeekStart)

						if err != nil {
							util.ErrorMsg(err.Error())
						}

						// ecribir mbr
						var buffer_bytes bytes.Buffer
						binary.Write(&buffer_bytes, binary.BigEndian, &mbr)

						_, err2 := disk.WriteAt(buffer_bytes.Bytes(), pos)

						if err2 != nil {
							util.ErrorMsg(err.Error())
						}

						pos2, err := disk.Seek(int64(used_space+1), io.SeekStart)
						if err != nil {
							util.ErrorMsg(err.Error())
						}

						var buffer_ebr bytes.Buffer
						ebr := util.NewEBR(used_space+1, size)

						binary.Write(&buffer_ebr, binary.BigEndian, &ebr)

						_, err3 := disk.WriteAt(buffer_ebr.Bytes(), pos2)

						if err3 != nil {
							util.ErrorMsg(err3.Error())
						}
						fmt.Println("> fdisk realizado con exito!")
					} else {
						util.ErrorMsg("Ya existe una particion con ese nombre!")
					}
				} else {
					util.ErrorMsg("No hay espacio suficiente para crear la partición!")
				}
			} else {
				util.ErrorMsg("Ya existe una partición extendida no se puede crear otra!")
			}
		} else if exists_extended && tipo == "l" {

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
