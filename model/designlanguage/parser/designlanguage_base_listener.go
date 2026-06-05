// Code generated from documentation/design/DesignLanguage.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // DesignLanguage
import "github.com/antlr4-go/antlr/v4"

// BaseDesignLanguageListener is a complete listener for a parse tree produced by DesignLanguageParser.
type BaseDesignLanguageListener struct{}

var _ DesignLanguageListener = &BaseDesignLanguageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDesignLanguageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDesignLanguageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDesignLanguageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDesignLanguageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDesign is called when production design is entered.
func (s *BaseDesignLanguageListener) EnterDesign(ctx *DesignContext) {}

// ExitDesign is called when production design is exited.
func (s *BaseDesignLanguageListener) ExitDesign(ctx *DesignContext) {}

// EnterComponent is called when production component is entered.
func (s *BaseDesignLanguageListener) EnterComponent(ctx *ComponentContext) {}

// ExitComponent is called when production component is exited.
func (s *BaseDesignLanguageListener) ExitComponent(ctx *ComponentContext) {}

// EnterField is called when production field is entered.
func (s *BaseDesignLanguageListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseDesignLanguageListener) ExitField(ctx *FieldContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseDesignLanguageListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseDesignLanguageListener) ExitAttribute(ctx *AttributeContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseDesignLanguageListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseDesignLanguageListener) ExitMethod(ctx *MethodContext) {}

// EnterParam is called when production param is entered.
func (s *BaseDesignLanguageListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseDesignLanguageListener) ExitParam(ctx *ParamContext) {}

// EnterParams is called when production params is entered.
func (s *BaseDesignLanguageListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseDesignLanguageListener) ExitParams(ctx *ParamsContext) {}

// EnterType is called when production type is entered.
func (s *BaseDesignLanguageListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseDesignLanguageListener) ExitType(ctx *TypeContext) {}
