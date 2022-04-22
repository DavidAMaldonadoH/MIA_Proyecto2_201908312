package main

import (
	comandos "Proyecto2/Comandos"
	parser "Proyecto2/Parser"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseGramaticaListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (*TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	result := ctx.GetList()
	for _, c := range result {
		exec, ok := c.(*comandos.Exec)
		if ok {
			path := exec.Execute().(string)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				Interpretar((string(data)))
			}
		} else {
			c.Execute()
		}
	}
}

func Interpretar(in string) {
	is := antlr.NewInputStream(in)

	lexer := parser.NewGramaticaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGramaticaParser(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	input := ""
	for input != "exit\n" {
		input = ""
		fmt.Print("davidmaldonado201908312:~/Proyecto 2$ ")
		input, _ := consoleReader.ReadString('\n')
		if input == "exit\n" {
			break
		} else {
			Interpretar(input)
		}
	}
}
