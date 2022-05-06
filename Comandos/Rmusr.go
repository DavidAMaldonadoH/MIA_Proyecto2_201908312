package comandos

import (
	util "Proyecto2/Util"
	"bytes"
	"os"
	"strings"
	"unsafe"
)

type Rmusr struct {
	Parameters []util.Parameter
	util.Command
}

func NewRmusr(parameters []util.Parameter) *Rmusr {
	return &Rmusr{Parameters: parameters}
}

func (rmusr *Rmusr) Execute() interface{} {
	if len(rmusr.Parameters) == 1 {
		if rmusr.Parameters[0].Key == "usuario" {
			removeUser(rmusr.Parameters[0].Value.(string))
		} else {
			util.ErrorMsg("El comando rmusr recibió un parámetro no permitido!")
		}
	} else {
		util.ErrorMsg("El comando rmusr recibió mas parámetros de los permitidos!")
	}
	return nil
}

func removeUser(usuario string) {
	if Is_logged_in {
		if Current_user.Name == "root" {
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

			pos2, err := disk.Seek(super_block.Inode_start+int64(unsafe.Sizeof(util.Inode{})), 0)
			if err != nil {
				util.ErrorMsg(err.Error())
			}
		
			content_file := make([]byte, 0)
			inode := util.ReadInode(disk, pos2)
			for _, ap := range inode.Block {
				if ap == -1 {
					break
				} else {
					pos3, err := disk.Seek(super_block.Block_start+ap*64, 0)
					if err != nil {
						util.ErrorMsg(err.Error())
					}
					content := util.ReadFile(disk, pos3)
					content_file = append(content_file, content.B_content[:]...)
				}
			}
			b := bytes.Trim(content_file[:], "\x00")
			content_file_clean := string(b)
			groups_users := strings.Split(content_file_clean, "\n")
			groups_users = groups_users[:len(groups_users)-1]
			found := false
			for i, v := range groups_users {
				info := strings.Split(v, ",")
				if len(info) == 5 {	
					if info[1] == "U" && info[3] == usuario {
						info[0] = "0"
						groups_users[i] = strings.Join(info, ",")
						found = true
						break
					}
				}
			}
			if found {
				content_file_clean = strings.Join(groups_users, "\n") + "\n"
				for index, ap := range inode.Block {
					if ap == -1 {
						break
					} else {
						if len(content_file_clean)%64 != 0 {
							for i := 0; i < len(content_file_clean)%64; i++ {
								content_file_clean += "\x00"
							}
						}
						sub_string := content_file_clean[index*64 : (index+1)*64]
						var bytes_data_users [64]byte
						copy(bytes_data_users[:], sub_string)
						var file_block util.File
						file_block.B_content = bytes_data_users
						pos3, err := disk.Seek(super_block.Block_start+ap*64, 0)
						if err != nil {
							util.ErrorMsg(err.Error())
						}
						util.WriteFile(disk, file_block, pos3)
					}
				}
				util.SuccessMsg("Usuario con nombre: " + usuario + " eliminado con éxito!")
			} else {
				util.ErrorMsg("No hay un usuario con ese nombre no se puede eliminar!")
			}
		} else {
			util.ErrorMsg("Solo el usuario root puede eliminar usuarios!")
		}
	} else {
		util.ErrorMsg("Se necesita que una sesión este iniciada para poder eliminar un usuario!")
	}
}
