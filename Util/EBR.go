package util

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
