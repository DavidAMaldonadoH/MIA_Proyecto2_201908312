package comandos

import (
	util "Proyecto2/Util"
	"bytes"
	"fmt"
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

	} else if strings.ToLower(name) == "file" {

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

