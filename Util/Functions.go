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