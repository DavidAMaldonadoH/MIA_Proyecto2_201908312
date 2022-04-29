package util

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

func ErrorMsg(text string) {
	fmt.Printf("\033[31m>> Error: %s\033[0m\r\n", text)
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

func StructToBytes(s interface{}) []byte {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(s)
	if err != nil && err != io.EOF {
		ErrorMsg(err.Error())
	}
	return buffer.Bytes()
}

func BytesToMBR(b []byte) MBR {
	m := MBR{}
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&m)
	if err != nil && err != io.EOF {
		ErrorMsg(err.Error())
	}
	return m
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
