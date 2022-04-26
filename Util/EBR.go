package util

type EBR struct {
	Status byte
	Fit    byte
	Start  int64
	Size   int64
	Next   int64
	Name   [25]byte
}
