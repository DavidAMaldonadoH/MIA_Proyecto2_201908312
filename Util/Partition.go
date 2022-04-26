package util

import "fmt"

type Partition struct {
	Status byte
	Type   byte
	Fit    byte
	Start  int64
	Size   int64
	Name   [25]byte
}

func NewPartition(tipo, fit byte, start, size int64, name string) *Partition {
	var name_bytes [25]byte
	copy(name_bytes[:], name)
	return &Partition{Status: byte("0"[0]), Type: tipo, Fit: fit, Start: start, Size: size, Name: name_bytes}
}

func (p *Partition) PrintInfo() {
	fmt.Printf("Status: %s\n", string(p.Status))
	fmt.Printf("Type: %s\n", string(p.Type))
	fmt.Printf("Fit: %s\n", string(p.Fit))
	fmt.Printf("Start: %d\n", p.Start)
	fmt.Printf("Size: %d\n", p.Size)
	fmt.Printf("Name: %s\n", p.Name)
}
