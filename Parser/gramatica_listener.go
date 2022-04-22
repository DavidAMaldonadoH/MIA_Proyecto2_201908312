// Code generated from Gramatica.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterCommands is called when entering the commands production.
	EnterCommands(c *CommandsContext)

	// EnterComando is called when entering the comando production.
	EnterComando(c *ComandoContext)

	// EnterFit_type is called when entering the fit_type production.
	EnterFit_type(c *Fit_typeContext)

	// EnterUnit_type is called when entering the unit_type production.
	EnterUnit_type(c *Unit_typeContext)

	// EnterPart_type is called when entering the part_type production.
	EnterPart_type(c *Part_typeContext)

	// EnterFormat_type is called when entering the format_type production.
	EnterFormat_type(c *Format_typeContext)

	// EnterTexto is called when entering the texto production.
	EnterTexto(c *TextoContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitCommands is called when exiting the commands production.
	ExitCommands(c *CommandsContext)

	// ExitComando is called when exiting the comando production.
	ExitComando(c *ComandoContext)

	// ExitFit_type is called when exiting the fit_type production.
	ExitFit_type(c *Fit_typeContext)

	// ExitUnit_type is called when exiting the unit_type production.
	ExitUnit_type(c *Unit_typeContext)

	// ExitPart_type is called when exiting the part_type production.
	ExitPart_type(c *Part_typeContext)

	// ExitFormat_type is called when exiting the format_type production.
	ExitFormat_type(c *Format_typeContext)

	// ExitTexto is called when exiting the texto production.
	ExitTexto(c *TextoContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)
}
