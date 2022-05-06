package comandos

import (
	util "Proyecto2/Util"
	"bytes"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

type Mkuser struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkuser(parameters []util.Parameter) *Mkuser {
	return &Mkuser{Parameters: parameters}
}

func (mkuser *Mkuser) Execute() interface{} {
	// parametros
	usuario := ""
	pwd := ""
	grp := ""
	err := false
	// verificar parametros
	for _, p := range mkuser.Parameters {
		if p.Key == "usuario" {
			usuario = p.Value.(string)
		} else if p.Key == "pwd" {
			pwd = p.Value.(string)
		} else if p.Key == "grp" {
			grp = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if usuario == "" || pwd == "" || grp == "" {
		err = true
	}

	if !err {
		makeUser(usuario, pwd, grp)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Mkuser")
	}

	return nil
}

func makeUser(usuario, pwd, grp string) {
	if len(usuario) <= 10 && len(pwd) <= 10 && len(grp) <= 10 {
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
				users := make([]string, 0)
				exists_group := false
				exists_user := false
				for _, v := range groups_users {
					info := strings.Split(v, ",")
					if info[1] == "G" {
						if info[2] == grp {
							exists_group = true
						}
					}
					if info[1] == "U" {
						users = append(users, v)
						if info[3] == usuario {
							exists_user = true
							break
						}
					}
				}
				if !exists_user {
					if exists_group {
						new_user := strconv.Itoa(len(users) + 1)
						new_user += ",U," + grp + "," + usuario + "," + pwd + "\n"
						content_file_clean += new_user
						index_aux := -1
						for index, ap := range inode.Block {
							if ap == -1 {
								index_aux = index
								break
							} else {
								inode.Size = int64(len(content_file_clean))
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
						if index_aux != -1 {
							if index_aux*64 < len(content_file_clean) {
								ap := util.CreateFileBlock(partition.P, &super_block, &inode, disk)
								sub_string := content_file_clean[index_aux*64:]
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

						if len(content_file_clean) > 1024 {
							util.ErrorMsg("El archivo a excedido el máximo del tamaño permitido!")
						}

						pos4, err := disk.Seek(super_block.Inode_start+int64(unsafe.Sizeof(util.Inode{})), 0)
						if err != nil {
							util.ErrorMsg(err.Error())
						}

						util.WriteInode(disk, inode, pos4)

						pos5, err := disk.Seek(partition.P.Start, 0)
						if err != nil {
							util.ErrorMsg(err.Error())
						}
						util.WriteSuperBlock(disk, super_block, pos5)
						util.SuccessMsg("Usuario con nombre: " + usuario + " creado con éxito!")
					} else {
						util.ErrorMsg("No hay un grupo con el mismo nombre!")
					}
				} else {
					util.ErrorMsg("Ya existe un usuario con el mismo nombre!")
				}
			} else {
				util.ErrorMsg("Solo el usuario root puede crear grupos!")
			}
		} else {
			util.ErrorMsg("Se necesita que una sesión este iniciada para poder eliminar un grupo!")
		}
	} else {
		util.ErrorMsg("Longitud de los parámetros incorrecta!")
	}
}
