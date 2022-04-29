package util

import "fmt"

type Partition struct {
	Name   [16]byte
	Status byte
	Type   byte
	Fit    byte
	_      [5]byte
	Start  int64
	Size   int64
}

func NewPartition() Partition {
	var name_bytes [16]byte
	copy(name_bytes[:], "")
	return Partition{Status: '0', Type: 'p', Fit: 'w', Start: -1, Size: 0, Name: name_bytes}
}

func (p Partition) PrintInfo() {
	fmt.Printf("Status: %s\n", string(p.Status))
	fmt.Printf("Type: %s\n", string(p.Type))
	fmt.Printf("Fit: %s\n", string(p.Fit))
	fmt.Printf("Start: %d\n", p.Start)
	fmt.Printf("Size: %d\n", p.Size)
	fmt.Printf("Name: %s\n", p.Name)
}
