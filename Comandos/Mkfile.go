package comandos

import (
	util "Proyecto2/Util"
	"bytes"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Mkfile struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkfile(parameters []util.Parameter) *Mkfile {
	return &Mkfile{Parameters: parameters}
}

func (mkfile *Mkfile) Execute() interface{} {
	// parametros
	path := ""
	r := false
	size := 0
	cont := ""
	err := false
	// verificar parametros
	for _, p := range mkfile.Parameters {
		if p.Key == "path" {
			path = p.Value.(string)
		} else if p.Key == "r" {
			r = true
		} else if p.Key == "size" {
			if p.Value.(int) < 0 {
				err = true
				break
			} else {
				size = p.Value.(int)
			}
		} else if p.Key == "cont" {
			cont = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if path == "" {
		err = true
	}

	if !err {
		makeFile(path, cont, r, size)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Mkfile")
	}

	return nil
}

func makeFile(path, cont string, r bool, size int) {
	if Is_logged_in {
		dir := strings.Split(path, "/")
		dir = dir[1:]
		partition := Current_user.Partition
		disk, err := os.OpenFile(partition.Path, os.O_RDWR, 0777)
		if err != nil {
			util.ErrorMsg(err.Error())
		}

		defer disk.Close()

		pos, err := disk.Seek(partition.P.Start, 0)
		if err != nil {
			util.ErrorMsg(err.Error())
		}

		super_block := util.ReadSuperBlock(disk, pos)
		pos2, err := disk.Seek(super_block.Inode_start, 0)
		if err != nil {
			util.ErrorMsg(err.Error())
		}

		inode0 := util.ReadInode(disk, pos2)
		exists_file := false
		index_ap := -1
		index_folder := -1
		ap_folder := -1
		for i, ap := range inode0.Block {
			if ap == -1 {
				index_ap = i
				break
			} else {
				pos3, err := disk.Seek(super_block.Block_start+(ap*64), 0)
				if err != nil {
					util.ErrorMsg(err.Error())
				}
				root_folder := util.ReadBlock(disk, pos3)
				if len(dir) == 1 {
					for j, cont := range root_folder.B_content {
						b := bytes.Trim(cont.Name[:], "\x00")
						nombre := string(b)
						if nombre == dir[0] {
							exists_file = true
							break
						}
						if cont.Inodo == -1 {
							ap_folder = int(ap)
							index_folder = j
							break
						}
					}
				} // else por si es una carpeta
			}
		}

		if !exists_file {
			nombre := dir[0]
			n := util.GetN(partition.P.Size)
			disk.Seek(super_block.Bm_inode_start, 0)
			bitmap_inodes := util.ReadBytes(disk, int(math.Floor(n)))
			index_inode := -1
			for i, bit := range bitmap_inodes {
				if bit == 0 {
					index_inode = i
					break
				}
			}
			currentTime := time.Now()
			time := currentTime.Format("01-02-2006")
			var time_bytes [16]byte
			copy(time_bytes[:], time)
			// inodo del archivo
			uid, _ := strconv.Atoi(Current_user.Uid)
			gid, _ := strconv.Atoi(Current_user.Gid)
			// si no hay espacio en la bloque de la carpeta crear otro
			var contenido util.Content
			if index_folder == -1 {
				// bloque carpeta
				var carpeta util.Folder
				copy(contenido.Name[:], nombre)      // nombre del archivo
				contenido.Inodo = int32(index_inode) // apuntador de inodo
				carpeta.B_content[0] = contenido
				// ultimos 3 bloques vacios
				contenido.Inodo = -1
				copy(contenido.Name[:], "")
				carpeta.B_content[1] = contenido
				contenido.Inodo = -1
				copy(contenido.Name[:], "")
				carpeta.B_content[2] = contenido
				contenido.Inodo = -1
				copy(contenido.Name[:], "")
				carpeta.B_content[3] = contenido
				pos_bloque := util.CreateFolderBlock(partition.P, &super_block, disk)
				pos3, err := disk.Seek(super_block.Block_start+pos_bloque*64, 0)
				inode0.Block[index_ap] = pos_bloque
				if err != nil {
					util.ErrorMsg(err.Error())
				}
				util.WriteBlock(disk, carpeta, pos3)
			} else {
				pos3, err := disk.Seek(super_block.Block_start+(int64(ap_folder)*64), 0)
				if err != nil {
					util.ErrorMsg(err.Error())
				}
				root_folder := util.ReadBlock(disk, pos3)
				copy(contenido.Name[:], nombre)      // nombre del archivo
				contenido.Inodo = int32(index_inode) // apuntador de inodo
				root_folder.B_content[index_folder] = contenido
				util.WriteBlock(disk, root_folder, pos3)
			}

			var file_inode util.Inode
			file_inode.Uid = int64(uid)
			file_inode.Gid = int64(gid)
			file_inode.Size = int64(size)
			// obtener fecha de creacion
			file_inode.Atime = time_bytes
			file_inode.Ctime = time_bytes
			file_inode.Mtime = time_bytes
			for i := 0; i < 16; i++ {
				file_inode.Block[i] = -1 // apuntadores en null
			}
			file_inode.Tipo = '1'
			file_inode.Perm = 664

			if cont != "" {
				contenido, err := os.Open(cont)
				if err != nil {
					util.ErrorMsg(err.Error())
				}
				data_bytes, _ := ioutil.ReadAll(contenido)
				file_inode.Size = int64(len(data_bytes))
				data := string(data_bytes[:])
				if len(data)%64 != 0 {
					for i := 0; i < len(data)%64; i++ {
						data += "\x00"
					}
				}
				for i := 0; i < len(data)/64; i++ {
					ap := util.CreateFileBlock(partition.P, &super_block, &file_inode, disk)
					sub_string := data[i*64 : (i+1)*64]
					var bytes_data_users [64]byte
					copy(bytes_data_users[:], sub_string)
					var file_block util.File
					file_block.B_content = bytes_data_users
					posx, err := disk.Seek(super_block.Block_start+ap*64, 0)
					if err != nil {
						util.ErrorMsg(err.Error())
					}
					util.WriteFile(disk, file_block, posx)
				}
    			contenido.Close()
			} else {
				var data string
				for i := 0; i < size; i++ {
					data += " "
				}
				if len(data)%64 != 0 {
					for i := 0; i < len(data)%64; i++ {
						data += "\x00"
					}
				}
				for i := 0; i < len(data)/64; i++ {
					ap := util.CreateFileBlock(partition.P, &super_block, &file_inode, disk)
					sub_string := data[i*64 : (i+1)*64]
					var bytes_data_users [64]byte
					copy(bytes_data_users[:], sub_string)
					var file_block util.File
					file_block.B_content = bytes_data_users
					posx, err := disk.Seek(super_block.Block_start+ap*64, 0)
					if err != nil {
						util.ErrorMsg(err.Error())
					}
					util.WriteFile(disk, file_block, posx)
				}
			}

			pos4, err := disk.Seek(super_block.Inode_start+(super_block.Inode_size*int64(index_inode)), 0)
			if err != nil {
				util.ErrorMsg(err.Error())
			}
			util.WriteInode(disk, file_inode, pos4)

			bitmap_inodes[index_inode] = 2
			pos5, _ := disk.Seek(super_block.Bm_inode_start, 0)
			util.WriteBitmap(disk, bitmap_inodes, pos5)
			super_block.Free_inodes_count -= 1

			pos6, err := disk.Seek(super_block.Inode_start, 0)
			if err != nil {
				util.ErrorMsg(err.Error())
			}
			util.WriteInode(disk, inode0, pos6)

			pos7, err := disk.Seek(partition.P.Start, 0)
			if err != nil {
				util.ErrorMsg(err.Error())
			}
			util.WriteSuperBlock(disk, super_block, pos7)
		} else {
			util.ErrorMsg("Ya existe un archivo con el mismo nombre!")
		}
	} else {
		util.ErrorMsg("Necesita haber una sesión iniciada para poder crear un archivo!")
	}
}
