package comandos

import (
	util "Proyecto2/Util"
	"math"
	"os"
	"strings"
	"time"
	"unsafe"
)

type Mkfs struct {
	Parameters []util.Parameter
	util.Command
}

func NewMkfs(parameters []util.Parameter) *Mkfs {
	return &Mkfs{Parameters: parameters}
}

func (mkfs *Mkfs) Execute() interface{} {
	// parametros
	id := ""
	tipo := "full"
	err := false
	// verificar parametros
	for _, p := range mkfs.Parameters {
		if p.Key == "id" {
			id = p.Value.(string)
		} else if p.Key == "type" {
			tipo = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if id == "" {
		err = true
	}

	if !err {
		formatPartition(id, tipo)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando Fdisk")
	}
	return nil
}

func formatPartition(id, tipo string) {
	id = strings.ToLower(id)
	partitions := GetMountedPartitions()
	var partition util.MountedPartition
	for _, p := range partitions {
		if id == p.Id {
			partition = p
			break
		}
	}
	n := util.GetN(partition.P.Size)
	numero_structs := math.Floor(n)

	if numero_structs >= 2 {
		disk, err := os.OpenFile(partition.Path, os.O_RDWR, 0777)
		if err != nil {
			util.ErrorMsg(err.Error())
		}

		if strings.ToLower(tipo) == "full" {
			formatFull(partition.P, disk)
		}

		createSuperBlock(partition.P, disk, numero_structs)
		createBitMaps(partition.P, disk, numero_structs)
		superb := util.ReadSuperBlock(disk, partition.P.Start)
		// fmt.Println(superb.Free_blocks_count)
		// fmt.Println(superb.Free_inodes_count)
		createRoot(partition.P, disk, n, &superb)
		pos, err := disk.Seek(partition.P.Start, 0)
		if err != nil {
			util.ErrorMsg(err.Error())
		}
		util.WriteSuperBlock(disk, superb, pos)
		// fmt.Println(superb.Free_blocks_count)
		// fmt.Println(superb.Free_inodes_count)
		disk.Close()
		util.InfoMsg("Formato de partición en " + partition.Path + " con nombre: " + partition.Name)
		util.SuccessMsg("mkfs realizado con éxito!")
	} else {
		util.ErrorMsg("No hay suficiente espacio para formatear la partición: " + partition.Name)
	}
}

func formatFull(partition util.Partition, disk *os.File) {
	_, err := disk.Seek(partition.Start, 0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}
	buffer := make([]byte, 1)
	buffer[0] = 0
	for i := 0; i < int(partition.Size); i++ {
		disk.Write(buffer)
		disk.Seek(partition.Start+int64(i), 0)
	}
}

func createSuperBlock(partition util.Partition, disk *os.File, n float64) {
	var super_block util.SuperBlock
	super_block.Filesystem_type = 0
	super_block.Inodes_count = int64(n)
	super_block.Blocks_count = 3 * int64(n)
	super_block.Free_blocks_count = 3 * int64(n)
	super_block.Free_inodes_count = int64(n)
	super_block.Mnt_count = 1
	// obtener fecha de creacion
	currentTime := time.Now()
	time := currentTime.Format("01-02-2006")
	var time_bytes [16]byte
	copy(time_bytes[:], time)
	super_block.M_time = time_bytes
	super_block.Magic = 0xEF53
	super_block.Inode_size = int64(unsafe.Sizeof(util.Inode{}))
	super_block.Block_size = int64(unsafe.Sizeof(util.File{}))
	super_block.First_ino = 2
	super_block.First_blo = 2
	super_block.Bm_inode_start = partition.Start + int64(unsafe.Sizeof(util.SuperBlock{}))
	super_block.Bm_block_start = super_block.Bm_inode_start + int64(n)
	super_block.Inode_start = super_block.Bm_block_start + 3*int64(n)
	super_block.Block_start = super_block.Bm_block_start + int64(n)*int64(unsafe.Sizeof(util.Inode{}))
	// posicionarse al inicio de la particion y escribir el super bloque
	pos, err := disk.Seek(partition.Start, 0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}

	util.WriteSuperBlock(disk, super_block, pos)
}

func createBitMaps(partition util.Partition, disk *os.File, n float64) {
	bitmap_inodes := make([]byte, int(n))
	bitmap_blocks := make([]byte, int(3*n))
	for i := 0; i < int(n); i++ {
		bitmap_inodes[i] = 0
	}
	for i := 0; i < 3*int(n); i++ {
		bitmap_blocks[i] = 0
	}
	bitmap_inodes[0] = 1 // inodo carpeta '/'
	bitmap_inodes[1] = 2 // inodo archivo 'users.txt'
	bitmap_blocks[0] = 1 // bloque carpeta '/' 1 para carpetas
	bitmap_blocks[1] = 2 // bloque archivo 'users.txt' 2 para archivos
	// escribir bitmap de inodos
	bm_inodes_start := partition.Start + int64(unsafe.Sizeof(util.SuperBlock{}))
	disk.WriteAt(bitmap_inodes, bm_inodes_start)
	// escribir bitmap de bloques
	bm_blocks_start := bm_inodes_start + int64(n)
	disk.WriteAt(bitmap_blocks, bm_blocks_start)
}

func createRoot(partition util.Partition, disk *os.File, n float64, super_block *util.SuperBlock) {
	// inodo carpeta '/'
	var carpeta_root util.Inode
	carpeta_root.Uid = 1
	carpeta_root.Gid = 1
	carpeta_root.Size = 64
	// obtener fecha de creacion
	currentTime := time.Now()
	time := currentTime.Format("01-02-2006")
	var time_bytes [16]byte
	copy(time_bytes[:], time)
	carpeta_root.Atime = time_bytes
	carpeta_root.Ctime = time_bytes
	carpeta_root.Mtime = time_bytes
	for i := 0; i < 16; i++ {
		carpeta_root.Block[i] = -1 // apuntadores en null
	}
	carpeta_root.Block[0] = 0 // apunta al bloque de carpeta
	carpeta_root.Tipo = '0'
	carpeta_root.Perm = 664
	// bloque carpeta
	var carpeta util.Folder
	var contenido util.Content
	// bloque carpeta '/'
	copy(contenido.Name[:], "/") // nombre de la carpeta
	contenido.Inodo = 0          // apuntador de inodo
	carpeta.B_content[0] = contenido
	// bloque archivo users.txt
	copy(contenido.Name[:], "users.txt") // nombre del archivo
	contenido.Inodo = 1                  // apuntador de inodo
	carpeta.B_content[1] = contenido
	// ultimos dos bloques vacios
	copy(contenido.Name[:], "")
	contenido.Inodo = -1
	carpeta.B_content[2] = contenido
	contenido.Inodo = -1
	copy(contenido.Name[:], "")
	carpeta.B_content[3] = contenido
	// escribir inodo carpeta '/'
	pos, err := disk.Seek(super_block.Inode_start, 0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}
	util.WriteInode(disk, carpeta_root, pos)
	super_block.Free_inodes_count -= 1 // disminuir la cantidad de indodos disponibles
	// escribir bloque de carpeta '/'
	pos2, err := disk.Seek(super_block.Block_start, 0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}
	util.WriteBlock(disk, carpeta, pos2)
	super_block.Free_blocks_count -= 1 // disminuimos la cantidad de bloques disponibles
	// archivo users.txt
	data_users := "1,G,root\n1,U,root,root,123\n"
	var bytes_data_users [64]byte
	copy(bytes_data_users[:], data_users)
	// inodo de archivo users.txt
	var users_file util.Inode
	users_file.Uid = 1
	users_file.Gid = 1
	users_file.Size = 27
	// obtener fecha de creacion
	users_file.Atime = time_bytes
	users_file.Ctime = time_bytes
	users_file.Mtime = time_bytes
	for i := 0; i < 16; i++ {
		users_file.Block[i] = -1 // apuntadores en null
	}
	users_file.Block[0] = 1 // inodo al que apunta
	users_file.Tipo = '1'
	users_file.Perm = 664
	// bloque de archivo 'users.txt'
	var file_block util.File
	file_block.B_content = bytes_data_users
	pos3, err := disk.Seek(super_block.Inode_start+int64(unsafe.Sizeof(util.Inode{})), 0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}
	util.WriteInode(disk, users_file, pos3)
	super_block.Free_inodes_count -= 1
	pos4, err := disk.Seek(super_block.Block_start+64, 0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}
	util.WriteFile(disk, file_block, pos4)
	super_block.Free_blocks_count -= 1
}
