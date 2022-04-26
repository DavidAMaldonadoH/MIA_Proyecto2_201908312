package util

import (
	"fmt"
	"math/rand"
	"time"
)

type MBR struct {
	Tamano     int64
	Fecha      [10]byte
	Signatue   int64
	Fit        byte
	Partitions []Partition
}

func NewMBR(tamano int64, fit byte) *MBR {
	// obtener fecha de creacion
	currentTime := time.Now()
	time := currentTime.Format("01-02-2006")
	var time_bytes [10]byte
	copy(time_bytes[:], time)
	// obtener firma
	random := rand.Intn(100)
	// particiones
	partitions := make([]Partition, 4)

	return &MBR{Tamano: tamano, Fecha: time_bytes, Signatue: int64(random), Fit: fit, Partitions: partitions}
}

func (m *MBR) PrintInfo() {
	fmt.Printf("Size: %d\n", m.Tamano)
	fmt.Printf("Time: %s\n", string(m.Fecha[:]))
	fmt.Printf("Signature: %d\n", m.Signatue)
	fmt.Printf("Fit: %s\n", string(m.Fit))
	for i, p := range m.Partitions {
		fmt.Printf("Partition %d:\n", i)
		p.PrintInfo()
	}
}
