package util

type Folder struct {
	B_content [4]Content
}

type Content struct {
	Name  [12]byte
	Inodo int32
}
