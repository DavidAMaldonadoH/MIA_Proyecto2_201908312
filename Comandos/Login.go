package comandos

import (
	util "Proyecto2/Util"
	"fmt"
	"os"
	"strings"
)

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
	pos, err := disk.Seek(partition.P.Start, 0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}
	super_block := util.ReadSuperBlock(disk, pos)

	pos2, err := disk.Seek(super_block.Block_start+64, 0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}
	file_block := util.ReadFile(disk, pos2)
	content_file := strings.Split(string(file_block.B_content[:]), "\n")
	fmt.Println(len(content_file))
}
