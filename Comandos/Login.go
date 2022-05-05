package comandos

import (
	util "Proyecto2/Util"
	"bytes"
	"os"
	"strings"
	"unsafe"
)

var Current_user User = User{Name: "", Password: ""}
var Is_logged_in bool = false

type User struct {
	Name string
	Password string
}

type Login struct {
	Parameters []util.Parameter
	util.Command
}

func NewLogin(parameters []util.Parameter) *Login {
	return &Login{Parameters: parameters}
}

func (login *Login) Execute() interface{} {
	// parametros
	usuario := ""
	password := ""
	id := ""
	err := false
	// verificar parametros
	for _, p := range login.Parameters {
		if p.Key == "usuario" {
			usuario = p.Value.(string)
		} else if p.Key == "password" {
			password = p.Value.(string)
		} else if p.Key == "id" {
			id = p.Value.(string)
		} else {
			err = true
		}
	}

	if usuario == "" || password == "" || id == "" {
		err = true
	}

	if !err {
		log(usuario, password, id)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Login")
	}

	return nil
}

func log(usuario, password, id string) {
	if !Is_logged_in {
		id = strings.ToLower(id)
		partitions := GetMountedPartitions()
		var partition util.MountedPartition
		for _, p := range partitions {
			if id == p.Id {
				partition = p
				break
			}
		}
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
		for _, v := range groups_users {
			info := strings.Split(v, ",")
			if info[1] == "U" {
				users =append(users, v)
			}
		}
		for _, u := range users {
			data := strings.Split(u, ",")
			if data[3] == usuario && data[4] == password {
				Current_user.Name = usuario
				Current_user.Password = password
				Is_logged_in = true
				if data[0] == "0" {
					Current_user.Name = ""
					Current_user.Password = ""
					Is_logged_in = false
					util.InfoMsg("El usuario " + usuario + " con la contraseña ingresada ya a sido eliminado!")
				}
			}
		}
		if Is_logged_in {
			util.InfoMsg("Bienvenido: " + usuario)
			util.SuccessMsg("Se ha iniciado sesión con éxito!")
		} else {
			util.ErrorMsg("No se ha podido iniciar sesión, credenciales incorrectas!")
		}
	} else {
		util.ErrorMsg("No se puede iniciar sesión ya hay una sesión abierta!")
	}
}
