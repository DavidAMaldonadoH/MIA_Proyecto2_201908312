package util

type Inode struct {
	Uid   int64
	Gid   int64
	Size  int64
	Perm  int64
	Atime [16]byte
	Mtime [16]byte
	Ctime [16]byte
	Block [16]int64
	Tipo  byte
}
