package util

import "fmt"

type EBR struct {
	Status byte
	Fit    byte
	_      [6]byte
	Start  int64
	Size   int64
	Next   int64
	Name   [16]byte
}

func NewEBR(start, size int) EBR {
	var name_bytes [16]byte
	copy(name_bytes[:], "")
	return EBR{Status: '0', Fit: 'f', Start: int64(start), Size: int64(size), Next: -1, Name: name_bytes}
}

func (e EBR) PrintInfo() {
	fmt.Printf("Status: %s\n", string(e.Status))
	fmt.Printf("Fit: %s\n", string(e.Fit))
	fmt.Printf("Start: %d\n", e.Start)
	fmt.Printf("Size: %d\n", e.Size)
	fmt.Printf("Next: %d\n", e.Next)
	fmt.Printf("Name: %s\n", string(e.Name[:]))
}
