package util

import "fmt"

type MountedPartition struct {
	Id   string
	Name string
	Path string
	P    Partition
}

func (m MountedPartition) ShowInfo() {
	fmt.Println("Id: " + m.Id)
	fmt.Println("Nombre: " + m.Name)
	fmt.Println("Path: " + m.Path)
}