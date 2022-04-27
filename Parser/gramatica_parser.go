// Code generated from Gramatica.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Proyecto2/Util"
import "Proyecto2/Comandos"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GramaticaParser struct {
	*antlr.BaseParser
}

var gramaticaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gramaticaParserInit() {
	staticData := &gramaticaParserStaticData
	staticData.literalNames = []string{
		"", "'.'", "'-'", "'='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "EXEC", "MKDISK", "RMDISK", "FDISK", "MOUNT", "MKFS",
		"LOGIN", "LOGOUT", "MKGRP", "RMGRP", "MKUSER", "RMUSR", "MKFILE", "MKDIR",
		"PAUSE", "REP", "SIZE", "FIT", "UNIT", "PATH", "TYPE", "NAME", "IDENT",
		"USUARIO", "PASSWORD", "PWD", "GRP", "RP", "PP", "CONT", "RUTA", "BF",
		"FF", "WF", "KILO", "MEGA", "BYTE", "PRIMARY", "EXTENDED", "LOGIC",
		"FAST", "FULL", "ENTERO", "STRING", "ID", "DIR", "COMS", "COMD", "WS",
	}
	staticData.ruleNames = []string{
		"start", "commands", "comando", "fit_type", "unit_type", "part_type",
		"format_type", "texto", "parameters", "param",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 52, 241, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 30, 8, 1, 10, 1, 12,
		1, 33, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 95, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 103,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 111, 8, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 119, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6,
		125, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 139, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 146, 8,
		8, 10, 8, 12, 8, 149, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 239, 8, 9, 1,
		9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 0, 273, 0, 20, 1, 0,
		0, 0, 2, 24, 1, 0, 0, 0, 4, 94, 1, 0, 0, 0, 6, 102, 1, 0, 0, 0, 8, 110,
		1, 0, 0, 0, 10, 118, 1, 0, 0, 0, 12, 124, 1, 0, 0, 0, 14, 138, 1, 0, 0,
		0, 16, 140, 1, 0, 0, 0, 18, 238, 1, 0, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22,
		5, 0, 0, 1, 22, 23, 6, 0, -1, 0, 23, 1, 1, 0, 0, 0, 24, 25, 3, 4, 2, 0,
		25, 31, 6, 1, -1, 0, 26, 27, 3, 4, 2, 0, 27, 28, 6, 1, -1, 0, 28, 30, 1,
		0, 0, 0, 29, 26, 1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31,
		32, 1, 0, 0, 0, 32, 3, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 35, 5, 4, 0,
		0, 35, 36, 3, 16, 8, 0, 36, 37, 6, 2, -1, 0, 37, 95, 1, 0, 0, 0, 38, 39,
		5, 5, 0, 0, 39, 40, 3, 16, 8, 0, 40, 41, 6, 2, -1, 0, 41, 95, 1, 0, 0,
		0, 42, 43, 5, 6, 0, 0, 43, 44, 3, 16, 8, 0, 44, 45, 6, 2, -1, 0, 45, 95,
		1, 0, 0, 0, 46, 47, 5, 7, 0, 0, 47, 48, 3, 16, 8, 0, 48, 49, 6, 2, -1,
		0, 49, 95, 1, 0, 0, 0, 50, 51, 5, 8, 0, 0, 51, 52, 3, 16, 8, 0, 52, 53,
		6, 2, -1, 0, 53, 95, 1, 0, 0, 0, 54, 55, 5, 9, 0, 0, 55, 56, 3, 16, 8,
		0, 56, 57, 6, 2, -1, 0, 57, 95, 1, 0, 0, 0, 58, 59, 5, 10, 0, 0, 59, 60,
		3, 16, 8, 0, 60, 61, 6, 2, -1, 0, 61, 95, 1, 0, 0, 0, 62, 63, 5, 11, 0,
		0, 63, 95, 6, 2, -1, 0, 64, 65, 5, 12, 0, 0, 65, 66, 3, 16, 8, 0, 66, 67,
		6, 2, -1, 0, 67, 95, 1, 0, 0, 0, 68, 69, 5, 13, 0, 0, 69, 70, 3, 16, 8,
		0, 70, 71, 6, 2, -1, 0, 71, 95, 1, 0, 0, 0, 72, 73, 5, 14, 0, 0, 73, 74,
		3, 16, 8, 0, 74, 75, 6, 2, -1, 0, 75, 95, 1, 0, 0, 0, 76, 77, 5, 15, 0,
		0, 77, 78, 3, 16, 8, 0, 78, 79, 6, 2, -1, 0, 79, 95, 1, 0, 0, 0, 80, 81,
		5, 16, 0, 0, 81, 82, 3, 16, 8, 0, 82, 83, 6, 2, -1, 0, 83, 95, 1, 0, 0,
		0, 84, 85, 5, 17, 0, 0, 85, 86, 3, 16, 8, 0, 86, 87, 6, 2, -1, 0, 87, 95,
		1, 0, 0, 0, 88, 89, 5, 18, 0, 0, 89, 95, 6, 2, -1, 0, 90, 91, 5, 19, 0,
		0, 91, 92, 3, 16, 8, 0, 92, 93, 6, 2, -1, 0, 93, 95, 1, 0, 0, 0, 94, 34,
		1, 0, 0, 0, 94, 38, 1, 0, 0, 0, 94, 42, 1, 0, 0, 0, 94, 46, 1, 0, 0, 0,
		94, 50, 1, 0, 0, 0, 94, 54, 1, 0, 0, 0, 94, 58, 1, 0, 0, 0, 94, 62, 1,
		0, 0, 0, 94, 64, 1, 0, 0, 0, 94, 68, 1, 0, 0, 0, 94, 72, 1, 0, 0, 0, 94,
		76, 1, 0, 0, 0, 94, 80, 1, 0, 0, 0, 94, 84, 1, 0, 0, 0, 94, 88, 1, 0, 0,
		0, 94, 90, 1, 0, 0, 0, 95, 5, 1, 0, 0, 0, 96, 97, 5, 35, 0, 0, 97, 103,
		6, 3, -1, 0, 98, 99, 5, 36, 0, 0, 99, 103, 6, 3, -1, 0, 100, 101, 5, 37,
		0, 0, 101, 103, 6, 3, -1, 0, 102, 96, 1, 0, 0, 0, 102, 98, 1, 0, 0, 0,
		102, 100, 1, 0, 0, 0, 103, 7, 1, 0, 0, 0, 104, 105, 5, 38, 0, 0, 105, 111,
		6, 4, -1, 0, 106, 107, 5, 39, 0, 0, 107, 111, 6, 4, -1, 0, 108, 109, 5,
		40, 0, 0, 109, 111, 6, 4, -1, 0, 110, 104, 1, 0, 0, 0, 110, 106, 1, 0,
		0, 0, 110, 108, 1, 0, 0, 0, 111, 9, 1, 0, 0, 0, 112, 113, 5, 41, 0, 0,
		113, 119, 6, 5, -1, 0, 114, 115, 5, 42, 0, 0, 115, 119, 6, 5, -1, 0, 116,
		117, 5, 43, 0, 0, 117, 119, 6, 5, -1, 0, 118, 112, 1, 0, 0, 0, 118, 114,
		1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 11, 1, 0, 0, 0, 120, 121, 5, 44,
		0, 0, 121, 125, 6, 6, -1, 0, 122, 123, 5, 45, 0, 0, 123, 125, 6, 6, -1,
		0, 124, 120, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 13, 1, 0, 0, 0, 126,
		127, 5, 47, 0, 0, 127, 139, 6, 7, -1, 0, 128, 129, 5, 49, 0, 0, 129, 139,
		6, 7, -1, 0, 130, 131, 5, 48, 0, 0, 131, 139, 6, 7, -1, 0, 132, 133, 5,
		48, 0, 0, 133, 134, 5, 1, 0, 0, 134, 135, 5, 48, 0, 0, 135, 139, 6, 7,
		-1, 0, 136, 137, 5, 46, 0, 0, 137, 139, 6, 7, -1, 0, 138, 126, 1, 0, 0,
		0, 138, 128, 1, 0, 0, 0, 138, 130, 1, 0, 0, 0, 138, 132, 1, 0, 0, 0, 138,
		136, 1, 0, 0, 0, 139, 15, 1, 0, 0, 0, 140, 141, 3, 18, 9, 0, 141, 147,
		6, 8, -1, 0, 142, 143, 3, 18, 9, 0, 143, 144, 6, 8, -1, 0, 144, 146, 1,
		0, 0, 0, 145, 142, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0,
		0, 147, 148, 1, 0, 0, 0, 148, 17, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150,
		151, 5, 2, 0, 0, 151, 152, 5, 20, 0, 0, 152, 153, 5, 3, 0, 0, 153, 154,
		5, 46, 0, 0, 154, 239, 6, 9, -1, 0, 155, 156, 5, 2, 0, 0, 156, 157, 5,
		21, 0, 0, 157, 158, 5, 3, 0, 0, 158, 159, 3, 6, 3, 0, 159, 160, 6, 9, -1,
		0, 160, 239, 1, 0, 0, 0, 161, 162, 5, 2, 0, 0, 162, 163, 5, 22, 0, 0, 163,
		164, 5, 3, 0, 0, 164, 165, 3, 8, 4, 0, 165, 166, 6, 9, -1, 0, 166, 239,
		1, 0, 0, 0, 167, 168, 5, 2, 0, 0, 168, 169, 5, 23, 0, 0, 169, 170, 5, 3,
		0, 0, 170, 171, 3, 14, 7, 0, 171, 172, 6, 9, -1, 0, 172, 239, 1, 0, 0,
		0, 173, 174, 5, 2, 0, 0, 174, 175, 5, 24, 0, 0, 175, 176, 5, 3, 0, 0, 176,
		177, 3, 10, 5, 0, 177, 178, 6, 9, -1, 0, 178, 239, 1, 0, 0, 0, 179, 180,
		5, 2, 0, 0, 180, 181, 5, 24, 0, 0, 181, 182, 5, 3, 0, 0, 182, 183, 3, 12,
		6, 0, 183, 184, 6, 9, -1, 0, 184, 239, 1, 0, 0, 0, 185, 186, 5, 2, 0, 0,
		186, 187, 5, 25, 0, 0, 187, 188, 5, 3, 0, 0, 188, 189, 3, 14, 7, 0, 189,
		190, 6, 9, -1, 0, 190, 239, 1, 0, 0, 0, 191, 192, 5, 2, 0, 0, 192, 193,
		5, 26, 0, 0, 193, 194, 5, 3, 0, 0, 194, 195, 5, 48, 0, 0, 195, 239, 6,
		9, -1, 0, 196, 197, 5, 2, 0, 0, 197, 198, 5, 27, 0, 0, 198, 199, 5, 3,
		0, 0, 199, 200, 3, 14, 7, 0, 200, 201, 6, 9, -1, 0, 201, 239, 1, 0, 0,
		0, 202, 203, 5, 2, 0, 0, 203, 204, 5, 28, 0, 0, 204, 205, 5, 3, 0, 0, 205,
		206, 3, 14, 7, 0, 206, 207, 6, 9, -1, 0, 207, 239, 1, 0, 0, 0, 208, 209,
		5, 2, 0, 0, 209, 210, 5, 29, 0, 0, 210, 211, 5, 3, 0, 0, 211, 212, 3, 14,
		7, 0, 212, 213, 6, 9, -1, 0, 213, 239, 1, 0, 0, 0, 214, 215, 5, 2, 0, 0,
		215, 216, 5, 30, 0, 0, 216, 217, 5, 3, 0, 0, 217, 218, 3, 14, 7, 0, 218,
		219, 6, 9, -1, 0, 219, 239, 1, 0, 0, 0, 220, 221, 5, 2, 0, 0, 221, 222,
		5, 31, 0, 0, 222, 239, 6, 9, -1, 0, 223, 224, 5, 2, 0, 0, 224, 225, 5,
		32, 0, 0, 225, 239, 6, 9, -1, 0, 226, 227, 5, 2, 0, 0, 227, 228, 5, 33,
		0, 0, 228, 229, 5, 3, 0, 0, 229, 230, 3, 14, 7, 0, 230, 231, 6, 9, -1,
		0, 231, 239, 1, 0, 0, 0, 232, 233, 5, 2, 0, 0, 233, 234, 5, 34, 0, 0, 234,
		235, 5, 3, 0, 0, 235, 236, 3, 14, 7, 0, 236, 237, 6, 9, -1, 0, 237, 239,
		1, 0, 0, 0, 238, 150, 1, 0, 0, 0, 238, 155, 1, 0, 0, 0, 238, 161, 1, 0,
		0, 0, 238, 167, 1, 0, 0, 0, 238, 173, 1, 0, 0, 0, 238, 179, 1, 0, 0, 0,
		238, 185, 1, 0, 0, 0, 238, 191, 1, 0, 0, 0, 238, 196, 1, 0, 0, 0, 238,
		202, 1, 0, 0, 0, 238, 208, 1, 0, 0, 0, 238, 214, 1, 0, 0, 0, 238, 220,
		1, 0, 0, 0, 238, 223, 1, 0, 0, 0, 238, 226, 1, 0, 0, 0, 238, 232, 1, 0,
		0, 0, 239, 19, 1, 0, 0, 0, 9, 31, 94, 102, 110, 118, 124, 138, 147, 238,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GramaticaParserInit initializes any static state used to implement GramaticaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGramaticaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GramaticaParserInit() {
	staticData := &gramaticaParserStaticData
	staticData.once.Do(gramaticaParserInit)
}

// NewGramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	GramaticaParserInit()
	this := new(GramaticaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &gramaticaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Gramatica.g4"

	return this
}

// GramaticaParser tokens.
const (
	GramaticaParserEOF      = antlr.TokenEOF
	GramaticaParserT__0     = 1
	GramaticaParserT__1     = 2
	GramaticaParserT__2     = 3
	GramaticaParserEXEC     = 4
	GramaticaParserMKDISK   = 5
	GramaticaParserRMDISK   = 6
	GramaticaParserFDISK    = 7
	GramaticaParserMOUNT    = 8
	GramaticaParserMKFS     = 9
	GramaticaParserLOGIN    = 10
	GramaticaParserLOGOUT   = 11
	GramaticaParserMKGRP    = 12
	GramaticaParserRMGRP    = 13
	GramaticaParserMKUSER   = 14
	GramaticaParserRMUSR    = 15
	GramaticaParserMKFILE   = 16
	GramaticaParserMKDIR    = 17
	GramaticaParserPAUSE    = 18
	GramaticaParserREP      = 19
	GramaticaParserSIZE     = 20
	GramaticaParserFIT      = 21
	GramaticaParserUNIT     = 22
	GramaticaParserPATH     = 23
	GramaticaParserTYPE     = 24
	GramaticaParserNAME     = 25
	GramaticaParserIDENT    = 26
	GramaticaParserUSUARIO  = 27
	GramaticaParserPASSWORD = 28
	GramaticaParserPWD      = 29
	GramaticaParserGRP      = 30
	GramaticaParserRP       = 31
	GramaticaParserPP       = 32
	GramaticaParserCONT     = 33
	GramaticaParserRUTA     = 34
	GramaticaParserBF       = 35
	GramaticaParserFF       = 36
	GramaticaParserWF       = 37
	GramaticaParserKILO     = 38
	GramaticaParserMEGA     = 39
	GramaticaParserBYTE     = 40
	GramaticaParserPRIMARY  = 41
	GramaticaParserEXTENDED = 42
	GramaticaParserLOGIC    = 43
	GramaticaParserFAST     = 44
	GramaticaParserFULL     = 45
	GramaticaParserENTERO   = 46
	GramaticaParserSTRING   = 47
	GramaticaParserID       = 48
	GramaticaParserDIR      = 49
	GramaticaParserCOMS     = 50
	GramaticaParserCOMD     = 51
	GramaticaParserWS       = 52
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start       = 0
	GramaticaParserRULE_commands    = 1
	GramaticaParserRULE_comando     = 2
	GramaticaParserRULE_fit_type    = 3
	GramaticaParserRULE_unit_type   = 4
	GramaticaParserRULE_part_type   = 5
	GramaticaParserRULE_format_type = 6
	GramaticaParserRULE_texto       = 7
	GramaticaParserRULE_parameters  = 8
	GramaticaParserRULE_param       = 9
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_commands returns the _commands rule contexts.
	Get_commands() ICommandsContext

	// Set_commands sets the _commands rule contexts.
	Set_commands(ICommandsContext)

	// GetList returns the list attribute.
	GetList() []util.Command

	// SetList sets the list attribute.
	SetList([]util.Command)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	list      []util.Command
	_commands ICommandsContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_commands() ICommandsContext { return s._commands }

func (s *StartContext) Set_commands(v ICommandsContext) { s._commands = v }

func (s *StartContext) GetList() []util.Command { return s.list }

func (s *StartContext) SetList(v []util.Command) { s.list = v }

func (s *StartContext) Commands() ICommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandsContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *GramaticaParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramaticaParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)

		var _x = p.Commands()

		localctx.(*StartContext)._commands = _x
	}
	{
		p.SetState(21)
		p.Match(GramaticaParserEOF)
	}
	localctx.(*StartContext).list = localctx.(*StartContext).Get_commands().GetList()

	return localctx
}

// ICommandsContext is an interface to support dynamic dispatch.
type ICommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC1 returns the c1 rule contexts.
	GetC1() IComandoContext

	// GetC2 returns the c2 rule contexts.
	GetC2() IComandoContext

	// SetC1 sets the c1 rule contexts.
	SetC1(IComandoContext)

	// SetC2 sets the c2 rule contexts.
	SetC2(IComandoContext)

	// GetList returns the list attribute.
	GetList() []util.Command

	// SetList sets the list attribute.
	SetList([]util.Command)

	// IsCommandsContext differentiates from other interfaces.
	IsCommandsContext()
}

type CommandsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	list   []util.Command
	c1     IComandoContext
	c2     IComandoContext
}

func NewEmptyCommandsContext() *CommandsContext {
	var p = new(CommandsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_commands
	return p
}

func (*CommandsContext) IsCommandsContext() {}

func NewCommandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandsContext {
	var p = new(CommandsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_commands

	return p
}

func (s *CommandsContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandsContext) GetC1() IComandoContext { return s.c1 }

func (s *CommandsContext) GetC2() IComandoContext { return s.c2 }

func (s *CommandsContext) SetC1(v IComandoContext) { s.c1 = v }

func (s *CommandsContext) SetC2(v IComandoContext) { s.c2 = v }

func (s *CommandsContext) GetList() []util.Command { return s.list }

func (s *CommandsContext) SetList(v []util.Command) { s.list = v }

func (s *CommandsContext) AllComando() []IComandoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComandoContext); ok {
			len++
		}
	}

	tst := make([]IComandoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComandoContext); ok {
			tst[i] = t.(IComandoContext)
			i++
		}
	}

	return tst
}

func (s *CommandsContext) Comando(i int) IComandoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoContext)
}

func (s *CommandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterCommands(s)
	}
}

func (s *CommandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitCommands(s)
	}
}

func (p *GramaticaParser) Commands() (localctx ICommandsContext) {
	this := p
	_ = this

	localctx = NewCommandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GramaticaParserRULE_commands)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)

		var _x = p.Comando()

		localctx.(*CommandsContext).c1 = _x
	}
	localctx.(*CommandsContext).list = append(localctx.(*CommandsContext).list, localctx.(*CommandsContext).GetC1().GetC())
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GramaticaParserEXEC)|(1<<GramaticaParserMKDISK)|(1<<GramaticaParserRMDISK)|(1<<GramaticaParserFDISK)|(1<<GramaticaParserMOUNT)|(1<<GramaticaParserMKFS)|(1<<GramaticaParserLOGIN)|(1<<GramaticaParserLOGOUT)|(1<<GramaticaParserMKGRP)|(1<<GramaticaParserRMGRP)|(1<<GramaticaParserMKUSER)|(1<<GramaticaParserRMUSR)|(1<<GramaticaParserMKFILE)|(1<<GramaticaParserMKDIR)|(1<<GramaticaParserPAUSE)|(1<<GramaticaParserREP))) != 0 {
		{
			p.SetState(26)

			var _x = p.Comando()

			localctx.(*CommandsContext).c2 = _x
		}
		localctx.(*CommandsContext).list = append(localctx.(*CommandsContext).list, localctx.(*CommandsContext).GetC2().GetC())

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IComandoContext is an interface to support dynamic dispatch.
type IComandoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_parameters returns the _parameters rule contexts.
	Get_parameters() IParametersContext

	// Set_parameters sets the _parameters rule contexts.
	Set_parameters(IParametersContext)

	// GetC returns the c attribute.
	GetC() util.Command

	// SetC sets the c attribute.
	SetC(util.Command)

	// IsComandoContext differentiates from other interfaces.
	IsComandoContext()
}

type ComandoContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	c           util.Command
	_parameters IParametersContext
}

func NewEmptyComandoContext() *ComandoContext {
	var p = new(ComandoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_comando
	return p
}

func (*ComandoContext) IsComandoContext() {}

func NewComandoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandoContext {
	var p = new(ComandoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_comando

	return p
}

func (s *ComandoContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandoContext) Get_parameters() IParametersContext { return s._parameters }

func (s *ComandoContext) Set_parameters(v IParametersContext) { s._parameters = v }

func (s *ComandoContext) GetC() util.Command { return s.c }

func (s *ComandoContext) SetC(v util.Command) { s.c = v }

func (s *ComandoContext) EXEC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEXEC, 0)
}

func (s *ComandoContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *ComandoContext) MKDISK() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMKDISK, 0)
}

func (s *ComandoContext) RMDISK() antlr.TerminalNode {
	return s.GetToken(GramaticaParserRMDISK, 0)
}

func (s *ComandoContext) FDISK() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFDISK, 0)
}

func (s *ComandoContext) MOUNT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMOUNT, 0)
}

func (s *ComandoContext) MKFS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMKFS, 0)
}

func (s *ComandoContext) LOGIN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLOGIN, 0)
}

func (s *ComandoContext) LOGOUT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLOGOUT, 0)
}

func (s *ComandoContext) MKGRP() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMKGRP, 0)
}

func (s *ComandoContext) RMGRP() antlr.TerminalNode {
	return s.GetToken(GramaticaParserRMGRP, 0)
}

func (s *ComandoContext) MKUSER() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMKUSER, 0)
}

func (s *ComandoContext) RMUSR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserRMUSR, 0)
}

func (s *ComandoContext) MKFILE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMKFILE, 0)
}

func (s *ComandoContext) MKDIR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMKDIR, 0)
}

func (s *ComandoContext) PAUSE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPAUSE, 0)
}

func (s *ComandoContext) REP() antlr.TerminalNode {
	return s.GetToken(GramaticaParserREP, 0)
}

func (s *ComandoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterComando(s)
	}
}

func (s *ComandoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitComando(s)
	}
}

func (p *GramaticaParser) Comando() (localctx IComandoContext) {
	this := p
	_ = this

	localctx = NewComandoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GramaticaParserRULE_comando)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserEXEC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(GramaticaParserEXEC)
		}
		{
			p.SetState(35)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewExec(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserMKDISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Match(GramaticaParserMKDISK)
		}
		{
			p.SetState(39)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewMkdisk(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserRMDISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.Match(GramaticaParserRMDISK)
		}
		{
			p.SetState(43)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewRmdisk(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserFDISK:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(46)
			p.Match(GramaticaParserFDISK)
		}
		{
			p.SetState(47)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewFdisk(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserMOUNT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(50)
			p.Match(GramaticaParserMOUNT)
		}
		{
			p.SetState(51)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewMount(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserMKFS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(54)
			p.Match(GramaticaParserMKFS)
		}
		{
			p.SetState(55)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewMkfs(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserLOGIN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(58)
			p.Match(GramaticaParserLOGIN)
		}
		{
			p.SetState(59)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewLogin(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserLOGOUT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(62)
			p.Match(GramaticaParserLOGOUT)
		}
		localctx.(*ComandoContext).c = comandos.NewLogout()

	case GramaticaParserMKGRP:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(64)
			p.Match(GramaticaParserMKGRP)
		}
		{
			p.SetState(65)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewMkgroup(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserRMGRP:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(68)
			p.Match(GramaticaParserRMGRP)
		}
		{
			p.SetState(69)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewRmgroup(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserMKUSER:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(72)
			p.Match(GramaticaParserMKUSER)
		}
		{
			p.SetState(73)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewMkuser(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserRMUSR:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(76)
			p.Match(GramaticaParserRMUSR)
		}
		{
			p.SetState(77)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewRmusr(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserMKFILE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(80)
			p.Match(GramaticaParserMKFILE)
		}
		{
			p.SetState(81)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewMkfile(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserMKDIR:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(84)
			p.Match(GramaticaParserMKDIR)
		}
		{
			p.SetState(85)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewMkdir(localctx.(*ComandoContext).Get_parameters().GetParams())

	case GramaticaParserPAUSE:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(88)
			p.Match(GramaticaParserPAUSE)
		}
		localctx.(*ComandoContext).c = comandos.NewPause()

	case GramaticaParserREP:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(90)
			p.Match(GramaticaParserREP)
		}
		{
			p.SetState(91)

			var _x = p.Parameters()

			localctx.(*ComandoContext)._parameters = _x
		}
		localctx.(*ComandoContext).c = comandos.NewRep(localctx.(*ComandoContext).Get_parameters().GetParams())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFit_typeContext is an interface to support dynamic dispatch.
type IFit_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFit returns the fit attribute.
	GetFit() string

	// SetFit sets the fit attribute.
	SetFit(string)

	// IsFit_typeContext differentiates from other interfaces.
	IsFit_typeContext()
}

type Fit_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	fit    string
}

func NewEmptyFit_typeContext() *Fit_typeContext {
	var p = new(Fit_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_fit_type
	return p
}

func (*Fit_typeContext) IsFit_typeContext() {}

func NewFit_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fit_typeContext {
	var p = new(Fit_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_fit_type

	return p
}

func (s *Fit_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Fit_typeContext) GetFit() string { return s.fit }

func (s *Fit_typeContext) SetFit(v string) { s.fit = v }

func (s *Fit_typeContext) BF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserBF, 0)
}

func (s *Fit_typeContext) FF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFF, 0)
}

func (s *Fit_typeContext) WF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserWF, 0)
}

func (s *Fit_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fit_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fit_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterFit_type(s)
	}
}

func (s *Fit_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitFit_type(s)
	}
}

func (p *GramaticaParser) Fit_type() (localctx IFit_typeContext) {
	this := p
	_ = this

	localctx = NewFit_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramaticaParserRULE_fit_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(102)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserBF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Match(GramaticaParserBF)
		}
		localctx.(*Fit_typeContext).fit = "bf"

	case GramaticaParserFF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Match(GramaticaParserFF)
		}
		localctx.(*Fit_typeContext).fit = "ff"

	case GramaticaParserWF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.Match(GramaticaParserWF)
		}
		localctx.(*Fit_typeContext).fit = "wf"

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnit_typeContext is an interface to support dynamic dispatch.
type IUnit_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetU returns the u attribute.
	GetU() string

	// SetU sets the u attribute.
	SetU(string)

	// IsUnit_typeContext differentiates from other interfaces.
	IsUnit_typeContext()
}

type Unit_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	u      string
}

func NewEmptyUnit_typeContext() *Unit_typeContext {
	var p = new(Unit_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_unit_type
	return p
}

func (*Unit_typeContext) IsUnit_typeContext() {}

func NewUnit_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unit_typeContext {
	var p = new(Unit_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_unit_type

	return p
}

func (s *Unit_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Unit_typeContext) GetU() string { return s.u }

func (s *Unit_typeContext) SetU(v string) { s.u = v }

func (s *Unit_typeContext) KILO() antlr.TerminalNode {
	return s.GetToken(GramaticaParserKILO, 0)
}

func (s *Unit_typeContext) MEGA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMEGA, 0)
}

func (s *Unit_typeContext) BYTE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserBYTE, 0)
}

func (s *Unit_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unit_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unit_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterUnit_type(s)
	}
}

func (s *Unit_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitUnit_type(s)
	}
}

func (p *GramaticaParser) Unit_type() (localctx IUnit_typeContext) {
	this := p
	_ = this

	localctx = NewUnit_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GramaticaParserRULE_unit_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserKILO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(GramaticaParserKILO)
		}
		localctx.(*Unit_typeContext).u = "k"

	case GramaticaParserMEGA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.Match(GramaticaParserMEGA)
		}
		localctx.(*Unit_typeContext).u = "m"

	case GramaticaParserBYTE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Match(GramaticaParserBYTE)
		}
		localctx.(*Unit_typeContext).u = "b"

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPart_typeContext is an interface to support dynamic dispatch.
type IPart_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t attribute.
	GetT() string

	// SetT sets the t attribute.
	SetT(string)

	// IsPart_typeContext differentiates from other interfaces.
	IsPart_typeContext()
}

type Part_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	t      string
}

func NewEmptyPart_typeContext() *Part_typeContext {
	var p = new(Part_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_part_type
	return p
}

func (*Part_typeContext) IsPart_typeContext() {}

func NewPart_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Part_typeContext {
	var p = new(Part_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_part_type

	return p
}

func (s *Part_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Part_typeContext) GetT() string { return s.t }

func (s *Part_typeContext) SetT(v string) { s.t = v }

func (s *Part_typeContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPRIMARY, 0)
}

func (s *Part_typeContext) EXTENDED() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEXTENDED, 0)
}

func (s *Part_typeContext) LOGIC() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLOGIC, 0)
}

func (s *Part_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Part_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Part_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterPart_type(s)
	}
}

func (s *Part_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitPart_type(s)
	}
}

func (p *GramaticaParser) Part_type() (localctx IPart_typeContext) {
	this := p
	_ = this

	localctx = NewPart_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GramaticaParserRULE_part_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserPRIMARY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(GramaticaParserPRIMARY)
		}
		localctx.(*Part_typeContext).t = "p"

	case GramaticaParserEXTENDED:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Match(GramaticaParserEXTENDED)
		}
		localctx.(*Part_typeContext).t = "e"

	case GramaticaParserLOGIC:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)
			p.Match(GramaticaParserLOGIC)
		}
		localctx.(*Part_typeContext).t = "l"

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFormat_typeContext is an interface to support dynamic dispatch.
type IFormat_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t attribute.
	GetT() string

	// SetT sets the t attribute.
	SetT(string)

	// IsFormat_typeContext differentiates from other interfaces.
	IsFormat_typeContext()
}

type Format_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	t      string
}

func NewEmptyFormat_typeContext() *Format_typeContext {
	var p = new(Format_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_format_type
	return p
}

func (*Format_typeContext) IsFormat_typeContext() {}

func NewFormat_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Format_typeContext {
	var p = new(Format_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_format_type

	return p
}

func (s *Format_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Format_typeContext) GetT() string { return s.t }

func (s *Format_typeContext) SetT(v string) { s.t = v }

func (s *Format_typeContext) FAST() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFAST, 0)
}

func (s *Format_typeContext) FULL() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFULL, 0)
}

func (s *Format_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Format_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Format_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterFormat_type(s)
	}
}

func (s *Format_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitFormat_type(s)
	}
}

func (p *GramaticaParser) Format_type() (localctx IFormat_typeContext) {
	this := p
	_ = this

	localctx = NewFormat_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GramaticaParserRULE_format_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserFAST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(GramaticaParserFAST)
		}
		localctx.(*Format_typeContext).t = "fast"

	case GramaticaParserFULL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(GramaticaParserFULL)
		}
		localctx.(*Format_typeContext).t = "full"

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITextoContext is an interface to support dynamic dispatch.
type ITextoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_DIR returns the _DIR token.
	Get_DIR() antlr.Token

	// GetId1 returns the id1 token.
	GetId1() antlr.Token

	// GetId2 returns the id2 token.
	GetId2() antlr.Token

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_DIR sets the _DIR token.
	Set_DIR(antlr.Token)

	// SetId1 sets the id1 token.
	SetId1(antlr.Token)

	// SetId2 sets the id2 token.
	SetId2(antlr.Token)

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// GetVal returns the val attribute.
	GetVal() string

	// SetVal sets the val attribute.
	SetVal(string)

	// IsTextoContext differentiates from other interfaces.
	IsTextoContext()
}

type TextoContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	val     string
	_STRING antlr.Token
	_DIR    antlr.Token
	id1     antlr.Token
	id2     antlr.Token
	_ENTERO antlr.Token
}

func NewEmptyTextoContext() *TextoContext {
	var p = new(TextoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_texto
	return p
}

func (*TextoContext) IsTextoContext() {}

func NewTextoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextoContext {
	var p = new(TextoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_texto

	return p
}

func (s *TextoContext) GetParser() antlr.Parser { return s.parser }

func (s *TextoContext) Get_STRING() antlr.Token { return s._STRING }

func (s *TextoContext) Get_DIR() antlr.Token { return s._DIR }

func (s *TextoContext) GetId1() antlr.Token { return s.id1 }

func (s *TextoContext) GetId2() antlr.Token { return s.id2 }

func (s *TextoContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *TextoContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *TextoContext) Set_DIR(v antlr.Token) { s._DIR = v }

func (s *TextoContext) SetId1(v antlr.Token) { s.id1 = v }

func (s *TextoContext) SetId2(v antlr.Token) { s.id2 = v }

func (s *TextoContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *TextoContext) GetVal() string { return s.val }

func (s *TextoContext) SetVal(v string) { s.val = v }

func (s *TextoContext) STRING() antlr.TerminalNode {
	return s.GetToken(GramaticaParserSTRING, 0)
}

func (s *TextoContext) DIR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDIR, 0)
}

func (s *TextoContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserID)
}

func (s *TextoContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, i)
}

func (s *TextoContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GramaticaParserENTERO, 0)
}

func (s *TextoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterTexto(s)
	}
}

func (s *TextoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitTexto(s)
	}
}

func (p *GramaticaParser) Texto() (localctx ITextoContext) {
	this := p
	_ = this

	localctx = NewTextoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GramaticaParserRULE_texto)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)

			var _m = p.Match(GramaticaParserSTRING)

			localctx.(*TextoContext)._STRING = _m
		}
		localctx.(*TextoContext).val = strings.ReplaceAll((func() string {
			if localctx.(*TextoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*TextoContext).Get_STRING().GetText()
			}
		}()), "\"", "")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)

			var _m = p.Match(GramaticaParserDIR)

			localctx.(*TextoContext)._DIR = _m
		}
		localctx.(*TextoContext).val = (func() string {
			if localctx.(*TextoContext).Get_DIR() == nil {
				return ""
			} else {
				return localctx.(*TextoContext).Get_DIR().GetText()
			}
		}())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)

			var _m = p.Match(GramaticaParserID)

			localctx.(*TextoContext).id1 = _m
		}
		localctx.(*TextoContext).val = (func() string {
			if localctx.(*TextoContext).GetId1() == nil {
				return ""
			} else {
				return localctx.(*TextoContext).GetId1().GetText()
			}
		}())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(132)

			var _m = p.Match(GramaticaParserID)

			localctx.(*TextoContext).id1 = _m
		}
		{
			p.SetState(133)
			p.Match(GramaticaParserT__0)
		}
		{
			p.SetState(134)

			var _m = p.Match(GramaticaParserID)

			localctx.(*TextoContext).id2 = _m
		}
		localctx.(*TextoContext).val = (func() string {
			if localctx.(*TextoContext).GetId1() == nil {
				return ""
			} else {
				return localctx.(*TextoContext).GetId1().GetText()
			}
		}()) + "." + (func() string {
			if localctx.(*TextoContext).GetId2() == nil {
				return ""
			} else {
				return localctx.(*TextoContext).GetId2().GetText()
			}
		}())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(136)

			var _m = p.Match(GramaticaParserENTERO)

			localctx.(*TextoContext)._ENTERO = _m
		}
		localctx.(*TextoContext).val = (func() string {
			if localctx.(*TextoContext).Get_ENTERO() == nil {
				return ""
			} else {
				return localctx.(*TextoContext).Get_ENTERO().GetText()
			}
		}())

	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 rule contexts.
	GetP1() IParamContext

	// GetP2 returns the p2 rule contexts.
	GetP2() IParamContext

	// SetP1 sets the p1 rule contexts.
	SetP1(IParamContext)

	// SetP2 sets the p2 rule contexts.
	SetP2(IParamContext)

	// GetParams returns the params attribute.
	GetParams() []util.Parameter

	// SetParams sets the params attribute.
	SetParams([]util.Parameter)

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	params []util.Parameter
	p1     IParamContext
	p2     IParamContext
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) GetP1() IParamContext { return s.p1 }

func (s *ParametersContext) GetP2() IParamContext { return s.p2 }

func (s *ParametersContext) SetP1(v IParamContext) { s.p1 = v }

func (s *ParametersContext) SetP2(v IParamContext) { s.p2 = v }

func (s *ParametersContext) GetParams() []util.Parameter { return s.params }

func (s *ParametersContext) SetParams(v []util.Parameter) { s.params = v }

func (s *ParametersContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *GramaticaParser) Parameters() (localctx IParametersContext) {
	this := p
	_ = this

	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GramaticaParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)

		var _x = p.Param()

		localctx.(*ParametersContext).p1 = _x
	}
	localctx.(*ParametersContext).params = append(localctx.(*ParametersContext).params, localctx.(*ParametersContext).GetP1().GetP())
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GramaticaParserT__1 {
		{
			p.SetState(142)

			var _x = p.Param()

			localctx.(*ParametersContext).p2 = _x
		}
		localctx.(*ParametersContext).params = append(localctx.(*ParametersContext).params, localctx.(*ParametersContext).GetP2().GetP())

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_fit_type returns the _fit_type rule contexts.
	Get_fit_type() IFit_typeContext

	// Get_unit_type returns the _unit_type rule contexts.
	Get_unit_type() IUnit_typeContext

	// Get_texto returns the _texto rule contexts.
	Get_texto() ITextoContext

	// Get_part_type returns the _part_type rule contexts.
	Get_part_type() IPart_typeContext

	// Get_format_type returns the _format_type rule contexts.
	Get_format_type() IFormat_typeContext

	// Set_fit_type sets the _fit_type rule contexts.
	Set_fit_type(IFit_typeContext)

	// Set_unit_type sets the _unit_type rule contexts.
	Set_unit_type(IUnit_typeContext)

	// Set_texto sets the _texto rule contexts.
	Set_texto(ITextoContext)

	// Set_part_type sets the _part_type rule contexts.
	Set_part_type(IPart_typeContext)

	// Set_format_type sets the _format_type rule contexts.
	Set_format_type(IFormat_typeContext)

	// GetP returns the p attribute.
	GetP() util.Parameter

	// SetP sets the p attribute.
	SetP(util.Parameter)

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            util.Parameter
	_ENTERO      antlr.Token
	_fit_type    IFit_typeContext
	_unit_type   IUnit_typeContext
	_texto       ITextoContext
	_part_type   IPart_typeContext
	_format_type IFormat_typeContext
	_ID          antlr.Token
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *ParamContext) Get_ID() antlr.Token { return s._ID }

func (s *ParamContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *ParamContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ParamContext) Get_fit_type() IFit_typeContext { return s._fit_type }

func (s *ParamContext) Get_unit_type() IUnit_typeContext { return s._unit_type }

func (s *ParamContext) Get_texto() ITextoContext { return s._texto }

func (s *ParamContext) Get_part_type() IPart_typeContext { return s._part_type }

func (s *ParamContext) Get_format_type() IFormat_typeContext { return s._format_type }

func (s *ParamContext) Set_fit_type(v IFit_typeContext) { s._fit_type = v }

func (s *ParamContext) Set_unit_type(v IUnit_typeContext) { s._unit_type = v }

func (s *ParamContext) Set_texto(v ITextoContext) { s._texto = v }

func (s *ParamContext) Set_part_type(v IPart_typeContext) { s._part_type = v }

func (s *ParamContext) Set_format_type(v IFormat_typeContext) { s._format_type = v }

func (s *ParamContext) GetP() util.Parameter { return s.p }

func (s *ParamContext) SetP(v util.Parameter) { s.p = v }

func (s *ParamContext) SIZE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserSIZE, 0)
}

func (s *ParamContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GramaticaParserENTERO, 0)
}

func (s *ParamContext) FIT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFIT, 0)
}

func (s *ParamContext) Fit_type() IFit_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFit_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFit_typeContext)
}

func (s *ParamContext) UNIT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserUNIT, 0)
}

func (s *ParamContext) Unit_type() IUnit_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnit_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnit_typeContext)
}

func (s *ParamContext) PATH() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPATH, 0)
}

func (s *ParamContext) Texto() ITextoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextoContext)
}

func (s *ParamContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserTYPE, 0)
}

func (s *ParamContext) Part_type() IPart_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPart_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPart_typeContext)
}

func (s *ParamContext) Format_type() IFormat_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormat_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormat_typeContext)
}

func (s *ParamContext) NAME() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNAME, 0)
}

func (s *ParamContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserIDENT, 0)
}

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *ParamContext) USUARIO() antlr.TerminalNode {
	return s.GetToken(GramaticaParserUSUARIO, 0)
}

func (s *ParamContext) PASSWORD() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPASSWORD, 0)
}

func (s *ParamContext) PWD() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPWD, 0)
}

func (s *ParamContext) GRP() antlr.TerminalNode {
	return s.GetToken(GramaticaParserGRP, 0)
}

func (s *ParamContext) RP() antlr.TerminalNode {
	return s.GetToken(GramaticaParserRP, 0)
}

func (s *ParamContext) PP() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPP, 0)
}

func (s *ParamContext) CONT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCONT, 0)
}

func (s *ParamContext) RUTA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserRUTA, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *GramaticaParser) Param() (localctx IParamContext) {
	this := p
	_ = this

	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramaticaParserRULE_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(151)
			p.Match(GramaticaParserSIZE)
		}
		{
			p.SetState(152)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(153)

			var _m = p.Match(GramaticaParserENTERO)

			localctx.(*ParamContext)._ENTERO = _m
		}
		localctx.(*ParamContext).p = util.NewParameter("size", (func() int {
			if localctx.(*ParamContext).Get_ENTERO() == nil {
				return 0
			} else {
				i, _ := strconv.Atoi(localctx.(*ParamContext).Get_ENTERO().GetText())
				return i
			}
		}()), true)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(156)
			p.Match(GramaticaParserFIT)
		}
		{
			p.SetState(157)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(158)

			var _x = p.Fit_type()

			localctx.(*ParamContext)._fit_type = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("fit", localctx.(*ParamContext).Get_fit_type().GetFit(), false)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(162)
			p.Match(GramaticaParserUNIT)
		}
		{
			p.SetState(163)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(164)

			var _x = p.Unit_type()

			localctx.(*ParamContext)._unit_type = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("unit", localctx.(*ParamContext).Get_unit_type().GetU(), false)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(167)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(168)
			p.Match(GramaticaParserPATH)
		}
		{
			p.SetState(169)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(170)

			var _x = p.Texto()

			localctx.(*ParamContext)._texto = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("path", localctx.(*ParamContext).Get_texto().GetVal(), true)

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(173)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(174)
			p.Match(GramaticaParserTYPE)
		}
		{
			p.SetState(175)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(176)

			var _x = p.Part_type()

			localctx.(*ParamContext)._part_type = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("type", localctx.(*ParamContext).Get_part_type().GetT(), false)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(179)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(180)
			p.Match(GramaticaParserTYPE)
		}
		{
			p.SetState(181)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(182)

			var _x = p.Format_type()

			localctx.(*ParamContext)._format_type = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("type", localctx.(*ParamContext).Get_format_type().GetT(), false)

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(185)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(186)
			p.Match(GramaticaParserNAME)
		}
		{
			p.SetState(187)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(188)

			var _x = p.Texto()

			localctx.(*ParamContext)._texto = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("name", localctx.(*ParamContext).Get_texto().GetVal(), true)

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(191)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(192)
			p.Match(GramaticaParserIDENT)
		}
		{
			p.SetState(193)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(194)

			var _m = p.Match(GramaticaParserID)

			localctx.(*ParamContext)._ID = _m
		}
		localctx.(*ParamContext).p = util.NewParameter("id", (func() string {
			if localctx.(*ParamContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParamContext).Get_ID().GetText()
			}
		}()), true)

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(196)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(197)
			p.Match(GramaticaParserUSUARIO)
		}
		{
			p.SetState(198)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(199)

			var _x = p.Texto()

			localctx.(*ParamContext)._texto = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("usuario", localctx.(*ParamContext).Get_texto().GetVal(), true)

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(202)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(203)
			p.Match(GramaticaParserPASSWORD)
		}
		{
			p.SetState(204)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(205)

			var _x = p.Texto()

			localctx.(*ParamContext)._texto = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("password", localctx.(*ParamContext).Get_texto().GetVal(), true)

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(208)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(209)
			p.Match(GramaticaParserPWD)
		}
		{
			p.SetState(210)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(211)

			var _x = p.Texto()

			localctx.(*ParamContext)._texto = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("pwd", localctx.(*ParamContext).Get_texto().GetVal(), true)

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(214)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(215)
			p.Match(GramaticaParserGRP)
		}
		{
			p.SetState(216)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(217)

			var _x = p.Texto()

			localctx.(*ParamContext)._texto = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("grp", localctx.(*ParamContext).Get_texto().GetVal(), true)

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(220)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(221)
			p.Match(GramaticaParserRP)
		}
		localctx.(*ParamContext).p = util.NewParameter("r", nil, false)

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(223)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(224)
			p.Match(GramaticaParserPP)
		}
		localctx.(*ParamContext).p = util.NewParameter("p", nil, false)

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(226)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(227)
			p.Match(GramaticaParserCONT)
		}
		{
			p.SetState(228)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(229)

			var _x = p.Texto()

			localctx.(*ParamContext)._texto = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("cont", localctx.(*ParamContext).Get_texto().GetVal(), false)

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(232)
			p.Match(GramaticaParserT__1)
		}
		{
			p.SetState(233)
			p.Match(GramaticaParserRUTA)
		}
		{
			p.SetState(234)
			p.Match(GramaticaParserT__2)
		}
		{
			p.SetState(235)

			var _x = p.Texto()

			localctx.(*ParamContext)._texto = _x
		}
		localctx.(*ParamContext).p = util.NewParameter("ruta", localctx.(*ParamContext).Get_texto().GetVal(), false)

	}

	return localctx
}
