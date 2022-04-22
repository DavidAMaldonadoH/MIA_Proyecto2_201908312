// Code generated from Gramatica.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type BaseGramaticaListener struct{}

var _ GramaticaListener = &BaseGramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseGramaticaListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseGramaticaListener) ExitStart(ctx *StartContext) {}

// EnterCommands is called when production commands is entered.
func (s *BaseGramaticaListener) EnterCommands(ctx *CommandsContext) {}

// ExitCommands is called when production commands is exited.
func (s *BaseGramaticaListener) ExitCommands(ctx *CommandsContext) {}

// EnterComando is called when production comando is entered.
func (s *BaseGramaticaListener) EnterComando(ctx *ComandoContext) {}

// ExitComando is called when production comando is exited.
func (s *BaseGramaticaListener) ExitComando(ctx *ComandoContext) {}

// EnterFit_type is called when production fit_type is entered.
func (s *BaseGramaticaListener) EnterFit_type(ctx *Fit_typeContext) {}

// ExitFit_type is called when production fit_type is exited.
func (s *BaseGramaticaListener) ExitFit_type(ctx *Fit_typeContext) {}

// EnterUnit_type is called when production unit_type is entered.
func (s *BaseGramaticaListener) EnterUnit_type(ctx *Unit_typeContext) {}

// ExitUnit_type is called when production unit_type is exited.
func (s *BaseGramaticaListener) ExitUnit_type(ctx *Unit_typeContext) {}

// EnterPart_type is called when production part_type is entered.
func (s *BaseGramaticaListener) EnterPart_type(ctx *Part_typeContext) {}

// ExitPart_type is called when production part_type is exited.
func (s *BaseGramaticaListener) ExitPart_type(ctx *Part_typeContext) {}

// EnterFormat_type is called when production format_type is entered.
func (s *BaseGramaticaListener) EnterFormat_type(ctx *Format_typeContext) {}

// ExitFormat_type is called when production format_type is exited.
func (s *BaseGramaticaListener) ExitFormat_type(ctx *Format_typeContext) {}

// EnterTexto is called when production texto is entered.
func (s *BaseGramaticaListener) EnterTexto(ctx *TextoContext) {}

// ExitTexto is called when production texto is exited.
func (s *BaseGramaticaListener) ExitTexto(ctx *TextoContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseGramaticaListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseGramaticaListener) ExitParameters(ctx *ParametersContext) {}

// EnterParam is called when production param is entered.
func (s *BaseGramaticaListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseGramaticaListener) ExitParam(ctx *ParamContext) {}
