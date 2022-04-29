package util

import (
	"fmt"
	"math/rand"
	"time"
)

type MBR struct {
	Tamano     int64
	Fit        byte
	_          [3]byte
	Signature  int32
	Fecha      [16]byte
	Partitions [4]Partition
}

func NewMBR(tamano int64, fit byte) MBR {
	// obtener fecha de creacion
	currentTime := time.Now()
	time := currentTime.Format("01-02-2006")
	var time_bytes [16]byte
	copy(time_bytes[:], time)
	// obtener firma
	random := rand.Intn(501)
	// particiones
	var partitions [4]Partition

	for i := 0; i < 4; i++ {
		partitions[i] = NewPartition()
	}

	return MBR{Tamano: tamano, Fecha: time_bytes, Signature: int32(random), Fit: fit, Partitions: partitions}
}

func (m MBR) PrintInfo() {
	fmt.Printf("Size: %d\n", m.Tamano)
	fmt.Printf("Time: %s\n", string(m.Fecha[:]))
	fmt.Printf("Signature: %d\n", m.Signature)
	fmt.Printf("Fit: %s\n", string(m.Fit))
	for i, p := range m.Partitions {
		fmt.Printf("Partition %d:\n", i+1)
		p.PrintInfo()
	}
}

func (m MBR) PrintInfo2() {
	fmt.Printf("Size: %d\n", m.Tamano)
	fmt.Printf("Time: %s\n", string(m.Fecha[:]))
	fmt.Printf("Signature: %d\n", m.Signature)
	fmt.Printf("Fit: %s\n", string(m.Fit))
}
