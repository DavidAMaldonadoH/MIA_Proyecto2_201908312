package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"unsafe"
)

func ErrorMsg(text string) {
	fmt.Printf("\033[31m>> Error: %s\033[0m\r\n", text)
}

func InfoMsg(text string) {
	fmt.Printf("\033[36m>> %s\033[0m\r\n", text)
}

func SuccessMsg(text string) {
	fmt.Printf("\033[32m>> %s\033[0m\r\n", text)
}

func CalcSize(size int, unit string) int {
	if unit == "b" {
		return size
	} else if unit == "k" {
		return size * 1024
	} else {
		return size * 1024 * 1024
	}
}

func ReadBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		ErrorMsg(err.Error())
	}

	return bytes
}

func UsedName(name string, names []string) bool {
	var name_bytes [16]byte
	copy(name_bytes[:], name)
	for _, n := range names {
		if string(name_bytes[:]) == n {
			return true
		}
	}
	return false
}

func WriteMbr(file *os.File, mbr MBR, pos int64) {
	var buffer_bytes bytes.Buffer
	binary.Write(&buffer_bytes, binary.BigEndian, &mbr)

	_, err2 := file.WriteAt(buffer_bytes.Bytes(), pos)

	if err2 != nil {
		ErrorMsg(err2.Error())
	}
}

func WriteEbr(file *os.File, ebr EBR, pos int64) {
	var buffer_ebr bytes.Buffer
	binary.Write(&buffer_ebr, binary.BigEndian, &ebr)

	_, err3 := file.WriteAt(buffer_ebr.Bytes(), pos)

	if err3 != nil {
		ErrorMsg(err3.Error())
	}
}

func WriteSuperBlock(file *os.File, sb SuperBlock, pos int64) {
	var buffer_bytes bytes.Buffer
	binary.Write(&buffer_bytes, binary.BigEndian, &sb)

	_, err2 := file.WriteAt(buffer_bytes.Bytes(), pos)

	if err2 != nil {
		ErrorMsg(err2.Error())
	}
}

func WriteInode(file *os.File, inode Inode, pos int64) {
	var buffer_bytes bytes.Buffer
	binary.Write(&buffer_bytes, binary.BigEndian, &inode)

	_, err2 := file.WriteAt(buffer_bytes.Bytes(), pos)

	if err2 != nil {
		ErrorMsg(err2.Error())
	}
}

func WriteBlock(file *os.File, block Folder, pos int64) {
	var buffer_bytes bytes.Buffer
	binary.Write(&buffer_bytes, binary.BigEndian, &block)

	_, err2 := file.WriteAt(buffer_bytes.Bytes(), pos)

	if err2 != nil {
		ErrorMsg(err2.Error())
	}
}

func WriteFile(file *os.File, file_block File, pos int64) {
	var buffer_bytes bytes.Buffer
	binary.Write(&buffer_bytes, binary.BigEndian, &file_block)

	_, err2 := file.WriteAt(buffer_bytes.Bytes(), pos)

	if err2 != nil {
		ErrorMsg(err2.Error())
	}
}

func WriteBitmap(file *os.File, bitmap []byte, pos int64) {
	var buffer_bytes bytes.Buffer
	binary.Write(&buffer_bytes, binary.BigEndian, &bitmap)

	_, err2 := file.WriteAt(buffer_bytes.Bytes(), pos)

	if err2 != nil {
		ErrorMsg(err2.Error())
	}
}


func ReadMbr(file *os.File) MBR {
	// posicionarse al inicio del archivo
	file.Seek(0, 0)
	// crear y obtener el tamano del mbr
	mbr := MBR{}
	size_mbr := unsafe.Sizeof(mbr)
	// leer del archivo y pasarlo al mbr
	data := ReadBytes(file, int(size_mbr))
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &mbr)
	if err != nil {
		ErrorMsg(err.Error())
	}

	return mbr
}

func ReadEbr(file *os.File, start int64) EBR {
	// posicionarse en el incio de la particion extendida
	file.Seek(start, 0)
	// crear y obtener el tamano del ebr
	ebr := EBR{}
	size_ebr := unsafe.Sizeof(ebr)
	// leer del archivo y pasarl al ebr
	data := ReadBytes(file, int(size_ebr))
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &ebr)
	if err != nil {
		ErrorMsg(err.Error())
	}

	return ebr
}

func ReadSuperBlock(file *os.File, start int64) SuperBlock {
	// posicionarse en el incio de la particion 
	// crear y obtener el tamano del super_block
	super_block := SuperBlock{}
	size_sb := unsafe.Sizeof(super_block)
	// leer del archivo y pasarlo a super bloque
	data := ReadBytes(file, int(size_sb))
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &super_block)
	if err != nil {
		ErrorMsg(err.Error())
	}

	return super_block
}

func ReadInode(file *os.File, start int64) Inode {
	// posicionarse en el incio de la particion 
	// crear y obtener el tamano del inode
	inode := Inode{}
	size_inode := unsafe.Sizeof(inode)
	// leer del archivo y pasarlo a inodo
	data := ReadBytes(file, int(size_inode))
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &inode)
	if err != nil {
		ErrorMsg(err.Error())
	}

	return inode
}

func ReadBlock(file *os.File, start int64) Folder {
	// posicionarse en el incio de la particion 
	// crear y obtener el tamano del bloque de carpeta
	folder := Folder{}
	size_folder := unsafe.Sizeof(folder)
	// leer del archivo y pasarlo a bloque de carpeta
	data := ReadBytes(file, int(size_folder))
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &folder)
	if err != nil {
		ErrorMsg(err.Error())
	}

	return folder
}

func ReadFile(file *os.File, start int64) File {
	// posicionarse en el incio de la particion 
	// crear y obtener el tamano del bloque de archivo
	file_block := File{}
	size_file := unsafe.Sizeof(file_block)
	// leer del archivo y pasarlo a bloque de archivo
	data := ReadBytes(file, int(size_file))
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &file_block)
	if err != nil {
		ErrorMsg(err.Error())
	}

	return file_block
}

func AlreadyMount(name, path string, mounted_partitions []MountedPartition) bool {
	for _, mp := range mounted_partitions {
		if mp.Path == path && mp.Name == name {
			return true
		}
	}
	return false
}

func FileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func GetN(part_size int64) float64 {
	var n float64
	size_sb := int(unsafe.Sizeof(SuperBlock{}))
	size_inode := int(unsafe.Sizeof(Inode{}))
	size_block := int(unsafe.Sizeof(Folder{}))
	n = float64(int(part_size)-size_sb) / float64(4+size_inode+(3*size_block))
	return n
}