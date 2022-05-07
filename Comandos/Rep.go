package comandos

import (
	util "Proyecto2/Util"
	"bytes"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"unsafe"
)

type Rep struct {
	Parameters []util.Parameter
	util.Command
}

func NewRep(parameters []util.Parameter) *Rep {
	return &Rep{Parameters: parameters}
}

func (rep *Rep) Execute() interface{} {
	// parametros
	path := ""
	name := ""
	id := ""
	ruta := ""
	err := false
	// verificar parametros
	for _, p := range rep.Parameters {
		if p.Key == "path" {
			path = p.Value.(string)
		} else if p.Key == "name" {
			name = p.Value.(string)
		} else if p.Key == "id" {
			id = p.Value.(string)
		} else if p.Key == "ruta" {
			ruta = p.Value.(string)
		} else {
			err = true
			break
		}
	}

	if path == "" || name == "" || id == "" {
		err = true
	}

	if !err {
		reportar(path, name, id, ruta)
	} else {
		util.ErrorMsg("Parámetros incorrectos en el comando rep")
	}

	return nil
}

func reportar(path, name, id, ruta string) {
	partitions := GetMountedPartitions()
	var partition util.MountedPartition
	for _, p := range partitions {
		if id == p.Id {
			partition = p
			break
		}
	}
	disk, err := os.OpenFile(partition.Path, os.O_RDWR, 0777)
	if err != nil {
		util.ErrorMsg(err.Error())
	}

	defer disk.Close()

	mbr := util.ReadMbr(disk)

	// verificar si existe el archivo .dot
	if !util.FileExists("reportes.dot") {
		archivo_dot, e := os.Create("reportes.dot")
		if e != nil {
			util.ErrorMsg(err.Error())
		}
		archivo_dot.Close()
	} else {
		if err := os.Truncate("reportes.dot", 0); err != nil {
			util.ErrorMsg(err.Error())
		}
	}

	// crear direccion de salida si no existe
	t := strings.Split(path, "/")
	dir_arr := t[:len(t)-1]
	dir := strings.Join(dir_arr[:], "/")
	if len(dir) > 0 {
		if _, err2 := os.Stat(path); os.IsNotExist(err2) {
			if err3 := os.MkdirAll(dir, os.ModePerm); err3 != nil {
				util.ErrorMsg(err3.Error())
			}
		}
	}

	if strings.ToLower(name) == "disk" {
		diskRep(mbr, disk)
		ext1 := filepath.Ext(path)
		app := "dot"
		arg0 := "-Nfontname=Arial"
		arg1 := "-T" + strings.Trim(ext1, ".")
		arg2 := "reportes.dot"
		arg3 := "-o"
		cmd := exec.Command(app, arg0, arg1, arg2, arg3, path)
		stdout, err := cmd.Output()
		if err != nil {
			util.ErrorMsg(err.Error())
			return
		}
		fmt.Print(string(stdout))
		util.SuccessMsg("Reporte disk creado con éxito!")
	} else if strings.ToLower(name) == "tree" {
		treeRep(disk)
		ext1 := filepath.Ext(path)
		app := "dot"
		arg0 := "-Nfontname=Arial"
		arg1 := "-T" + strings.Trim(ext1, ".")
		arg2 := "reportes.dot"
		arg3 := "-o"
		cmd := exec.Command(app, arg0, arg1, arg2, arg3, path)
		stdout, err := cmd.Output()
		if err != nil {
			util.ErrorMsg(err.Error())
			return
		}
		fmt.Print(string(stdout))
		util.SuccessMsg("Reporte tree creado con éxito!")
	} else if strings.ToLower(name) == "file" {
		if Is_logged_in {
			fileRep(disk, path, ruta)
			util.SuccessMsg("Reporte file creado con éxito!")
		} else {
			util.ErrorMsg("Necesita que una sesión este iniciada para crear el reporte!")
		}
	} else {
		util.ErrorMsg(name + " no es un tipo de reporte válido!")
	}
}

func diskRep(mbr util.MBR, disk *os.File) {
	dot_file, err := os.OpenFile("reportes.dot", os.O_RDWR, 0777)
	if err != nil {
		util.ErrorMsg(err.Error())
	}
	// cabecera del archivo dot
	dot_file.WriteString("digraph G {\n\trankdir=LR\n\tbgcolor=\"#fbd1a2\"\n\tfontname=Arial")
	dot_file.WriteString("\n\tlabelloc=t\n\tlabel=\"Reporte Disk " + strconv.Itoa(int(mbr.Signature)) + "\"\n\t")
	dot_file.WriteString("node [shape=plaintext fontname=Arial]")
	dot_file.WriteString("\n\tstrcut [label=<\n\t<TABLE>\n\t\t<TR>")
	// cuerpo del reporte
	valor := float64(unsafe.Sizeof(mbr)) / float64(mbr.Tamano)
	valor_str := fmt.Sprintf("%.4f", valor)
	dot_file.WriteString("\n\t\t\t<TD bgcolor=\"#f79256\">MBR<BR/>" + valor_str + "%</TD>")
	for i, part := range mbr.Partitions {
		// particion primaria
		if part.Type == 'p' && part.Start != -1 {
			valor = float64(part.Size) / float64(mbr.Tamano)
			valor_str = fmt.Sprintf("%.4f", valor)
			b := bytes.Trim(part.Name[:], "\x00")
			nombre := string(b)
			dot_file.WriteString("\n\t\t\t<TD bgcolor=\"#7dcfb6\">" + nombre + "<BR/>" + valor_str + "%</TD>")
		}
		// espacio libre
		if part.Start == -1 {
			if i > 0 {
				valor = (float64(mbr.Tamano) - float64(unsafe.Sizeof(mbr)) - float64(mbr.Partitions[i-1].Size)) / float64(mbr.Tamano)
				valor_str = fmt.Sprintf("%.4f", valor)
			} else {
				valor = (float64(mbr.Tamano) - float64(unsafe.Sizeof(mbr))) / float64(mbr.Tamano)
				valor_str = fmt.Sprintf("%.4f", valor)
			}
			dot_file.WriteString("\n\t\t\t<TD bgcolor=\"#00b2ca\"> Espacio libre<BR/>" + valor_str + "%</TD>")
			break
		}
		// extendida
		if part.Type == 'e' && part.Start != -1 {
			var fin int64
			dot_file.WriteString("\n\t\t\t<TD>")
			dot_file.WriteString("\n\t\t\t\t<TABLE bgcolor=\"#48cae4\" BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"4\">")
			dot_file.WriteString("\n\t\t\t\t\t<TR><TD COLSPAN=\"100%\">Extendida</TD></TR>")
			dot_file.WriteString("\n\t\t\t\t\t<TR>")
			ebr := util.ReadEbr(disk, part.Start)
			for ebr.Next != -1 {
				valor := float64(unsafe.Sizeof(ebr)) / float64(mbr.Tamano)
				valor_str := fmt.Sprintf("%.4f", valor)
				dot_file.WriteString("\n\t\t\t\t\t\t<TD bgcolor=\"#f79256\">EBR<BR/>" + valor_str + "%</TD>")
				valor = float64(ebr.Size) / float64(mbr.Tamano)
				valor_str = fmt.Sprintf("%.4f", valor)
				b := bytes.Trim(ebr.Name[:], "\x00")
				nombre := string(b)
				dot_file.WriteString("\n\t\t\t\t\t\t<TD bgcolor=\"#7dcfb6\">" + nombre + "<BR/>" + valor_str + "%</TD>")
				pos, err := disk.Seek(ebr.Next, 0)
				if err != nil {
					util.ErrorMsg(err.Error())
				}
				fin += int64(unsafe.Sizeof(ebr)) + ebr.Size
				ebr = util.ReadEbr(disk, pos)
			}
			valor := float64(unsafe.Sizeof(ebr)) / float64(mbr.Tamano)
			valor_str := fmt.Sprintf("%.4f", valor)
			dot_file.WriteString("\n\t\t\t\t\t\t<TD bgcolor=\"#f79256\">EBR<BR/>" + valor_str + "%</TD>")
			fin += int64(unsafe.Sizeof(ebr)) + part.Start
			valor = float64(fin) / float64(mbr.Tamano)
			valor_str = fmt.Sprintf("%.4f", valor)
			dot_file.WriteString("\n\t\t\t\t\t\t<TD bgcolor=\"#00b2ca\"> Espacio libre<BR/>" + valor_str + "%</TD>")
			dot_file.WriteString("\n\t\t\t\t\t</TR>")
			dot_file.WriteString("\n\t\t\t\t</TABLE>")
			dot_file.WriteString("\n\t\t\t</TD>")
		}
	}
	// final del archivo dot
	dot_file.WriteString("\n\t\t</TR>\n\t</TABLE>>\n\t]\n}")
	dot_file.Close()
}

func treeRep(disk *os.File) {
	if Is_logged_in {
		dot_file, err := os.OpenFile("reportes.dot", os.O_RDWR, 0777)
		if err != nil {
			util.ErrorMsg(err.Error())
		}
		// cuerpo del archivo
		partition := Current_user.Partition
		pos, err := disk.Seek(partition.P.Start, 0)
		if err != nil {
			util.ErrorMsg(err.Error())
		}
		super_block := util.ReadSuperBlock(disk, pos)
		n := util.GetN(partition.P.Size)
		disk.Seek(super_block.Bm_block_start, 0)
		bitmap_blocks := util.ReadBytes(disk, 3*int(math.Floor(n)))
		// cabecera del archivo dot
		dot_file.WriteString("digraph G {\n\tlabel=\"Reporte Tree " + Current_user.Partition.Id + "\"\n")
		dot_file.WriteString("\tfontname=Arial\n\tlabelloc=t\n\trankdir=LR\n\tbgcolor=\"#edf2f4\"\n\tfontsize=24")
		dot_file.WriteString("\n\tnode [ shape=none fontname=Arial]")
		for i, bit := range bitmap_blocks {
			node1 := fmt.Sprintf("bloque%d", i)
			pos2, err := disk.Seek(super_block.Block_start+(int64(i)*64), 0)
			if err != nil {
				util.ErrorMsg(err.Error())
			}
			if bit == 1 {
				folder_block := util.ReadBlock(disk, pos2)
				dot_file.WriteString("\n\t" + node1 + "[label=\"Bloque Carpeta\" shape=rectangle fillcolor=\"#ff9f1c\" style=filled]")
				for j, content := range folder_block.B_content {
					if content.Inodo == -1 {
						break
					}
					node2 := fmt.Sprintf("bloque%d%d", i, j)
					dot_file.WriteString("\n\t" + node2 + "[ label = <\n\t\t<table>\n\t\t\t<tr>")
					dot_file.WriteString("\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#ffbf69\">Nombre</td>")
					b := bytes.Trim(content.Name[:], "\x00")
					nombre := string(b)
					dot_file.WriteString("\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#ffbf69\">" + nombre + "</td>")
					dot_file.WriteString("\n\t\t\t</tr>\n\t\t\t<tr>")
					dot_file.WriteString("\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#ffbf69\">Apuntador</td>")
					ap := strconv.Itoa(int(content.Inodo))
					dot_file.WriteString("\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#ffbf69\">" + ap + "</td>")
					dot_file.WriteString("\n\t\t\t</tr>\n\t\t</table>\n\t> ]")
					dot_file.WriteString("\n\t" + node1 + "->" + node2)
					if node2 != "bloque00" {
						if node2 != "bloque01" {
							dot_file.WriteString("\n\t" + node2 + "->" + "inode" + ap)
						}
					}
				}
			} else if bit == 2 {
				file_block := util.ReadFile(disk, pos2)
				b := bytes.Trim(file_block.B_content[:], "\x00")
				contenido := string(b)
				contenido2 := strings.ReplaceAll(contenido, "\n", "\\n")
				dot_file.WriteString("\n\t" + node1 + "[label=\"" + contenido2 + "\" shape=rectangle fillcolor=\"#ff9f1c\" style=filled]")
			} else {
				break
			}
		}
		disk.Seek(super_block.Bm_inode_start, 0)
		bitmap_inodes := util.ReadBytes(disk, int(math.Floor(n)))
		for i, bit := range bitmap_inodes {
			if bit == 0 {
				break
			} else {
				node1 := fmt.Sprintf("inode%d", i)
				pos3, err := disk.Seek(super_block.Inode_start+(int64(i)*super_block.Inode_size), 0)
				if err != nil {
					util.ErrorMsg(err.Error())
				}
				inode := util.ReadInode(disk, pos3)
				gid := strconv.Itoa(int(inode.Gid))
				uid := strconv.Itoa(int(inode.Uid))
				size := strconv.Itoa(int(inode.Size))
				perm := strconv.Itoa(int(inode.Perm))
				b := bytes.Trim(inode.Ctime[:], "\x00")
				ctime := string(b)
				b = bytes.Trim(inode.Atime[:], "\x00")
				atime := string(b)
				b = bytes.Trim(inode.Mtime[:], "\x00")
				mtime := string(b)
				dot_file.WriteString("\n\t" + node1 + " [label = <\n\t\t<table>\n\t\t\t<tr><td colspan=\"2\" bgcolor=\"#2ec4b6\">Inodo</td>\n\t\t\t</tr>")
				dot_file.WriteString("\n\t\t\t<tr>\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">UID</td>")
				dot_file.WriteString("\n\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">" + uid + "</td>\n\t\t\t</tr>")
				dot_file.WriteString("\n\t\t\t<tr>\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">GID</td>")
				dot_file.WriteString("\n\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">" + gid + "</td>\n\t\t\t</tr>")
				dot_file.WriteString("\n\t\t\t<tr>\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">Size</td>")
				dot_file.WriteString("\n\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">" + size + "</td>\n\t\t\t</tr>")
				dot_file.WriteString("\n\t\t\t<tr>\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">atime</td>")
				dot_file.WriteString("\n\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">" + atime + "</td>\n\t\t\t</tr>")
				dot_file.WriteString("\n\t\t\t<tr>\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">ctime</td>")
				dot_file.WriteString("\n\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">" + ctime + "</td>\n\t\t\t</tr>")
				dot_file.WriteString("\n\t\t\t<tr>\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">mtime</td>")
				dot_file.WriteString("\n\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">" + mtime + "</td>\n\t\t\t</tr>")
				for j, ap := range inode.Block {
					if ap == -1 {
						break
					} else {
						index_ap := strconv.Itoa(j)
						valor_ap := strconv.Itoa(int(ap))
						dot_file.WriteString("\n\t\t\t<tr>\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">ap" + index_ap + "</td>")
						dot_file.WriteString("\n\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">" + valor_ap + "</td>\n\t\t\t</tr>")
					}
				}

				if bit == 1 {
					dot_file.WriteString("\n\t\t\t<tr>\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">Type</td>")
					dot_file.WriteString("\n\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">0</td>\n\t\t\t</tr>")
				} else {
					dot_file.WriteString("\n\t\t\t<tr>\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">Type</td>")
					dot_file.WriteString("\n\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">1</td>\n\t\t\t</tr>")
				}
				dot_file.WriteString("\n\t\t\t<tr>\n\t\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">Perm</td>")
				dot_file.WriteString("\n\t\t\t<td colspan=\"1\" bgcolor=\"#cbf3f0\">" + perm + "</td>\n\t\t\t</tr>")
				dot_file.WriteString("\n\t\t</table>\n\t> ]")
				for _, ap := range inode.Block {
					if ap == -1 {
						break
					} else {
						valor := strconv.Itoa(int(ap))
						dot_file.WriteString("\n" + node1 + "->" + "bloque" + valor)
					}
				}
			}
		}
		dot_file.WriteString("\n}")
		dot_file.Close()
	} else {
		util.ErrorMsg("Necesita que una sesión este iniciada para crear el reporte!")
	}
}

func fileRep(disk *os.File, path, ruta string) {
	archivo, err := os.Create(path)
	if err != nil {
		util.ErrorMsg(err.Error())
	}
	// cuerpo del archivo
	dir := strings.Split(ruta, "/")
	dir = dir[1:]
	partition := Current_user.Partition

	pos, err := disk.Seek(partition.P.Start, 0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}

	super_block := util.ReadSuperBlock(disk, pos)
	pos2, err := disk.Seek(super_block.Inode_start, 0)
	if err != nil {
		util.ErrorMsg(err.Error())
	}

	inode0 := util.ReadInode(disk, pos2)
	for _, ap := range inode0.Block {
		if ap == -1 {
			break
		} else {
			pos3, err := disk.Seek(super_block.Block_start+(ap*64), 0)
			if err != nil {
				util.ErrorMsg(err.Error())
			}
			root_folder := util.ReadBlock(disk, pos3)
			if len(dir) == 1 {
				for _, cont := range root_folder.B_content {
					if cont.Inodo == -1 {
						break
					}
					b := bytes.Trim(cont.Name[:], "\x00")
					nombre := string(b)
					if nombre == dir[0] {
						pos4, err := disk.Seek(super_block.Inode_start+(int64(cont.Inodo)*super_block.Inode_size), 0)
						if err != nil {
							util.ErrorMsg(err.Error())
						}
						content_file := make([]byte, 0)
						inode_aux := util.ReadInode(disk, pos4)
						for _, ap2 := range inode_aux.Block {
							if ap2 == -1 {
								break
							} else {
								pos5, err := disk.Seek(super_block.Block_start+ap2*64, 0)
								if err != nil {
									util.ErrorMsg(err.Error())
								}
								content := util.ReadFile(disk, pos5)
								content_file = append(content_file, content.B_content[:]...)
							}
						}
						b2 := bytes.Trim(content_file[:], "\x00")
						content_file_clean := string(b2)
						archivo.WriteString(content_file_clean)
					}
				}
			} // else por si es una carpeta
		}
	}
	archivo.Close()
}
