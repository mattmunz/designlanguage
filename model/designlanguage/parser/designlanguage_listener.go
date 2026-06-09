// Code generated from documentation/design/DesignLanguage.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // DesignLanguage
import "github.com/antlr4-go/antlr/v4"

// DesignLanguageListener is a complete listener for a parse tree produced by DesignLanguageParser.
type DesignLanguageListener interface {
	antlr.ParseTreeListener

	// EnterDesign is called when entering the design production.
	EnterDesign(c *DesignContext)

	// EnterPreamble is called when entering the preamble production.
	EnterPreamble(c *PreambleContext)

	// EnterAuthor is called when entering the author production.
	EnterAuthor(c *AuthorContext)

	// EnterComponent is called when entering the component production.
	EnterComponent(c *ComponentContext)

	// EnterSimpleComponent is called when entering the simpleComponent production.
	EnterSimpleComponent(c *SimpleComponentContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// ExitDesign is called when exiting the design production.
	ExitDesign(c *DesignContext)

	// ExitPreamble is called when exiting the preamble production.
	ExitPreamble(c *PreambleContext)

	// ExitAuthor is called when exiting the author production.
	ExitAuthor(c *AuthorContext)

	// ExitComponent is called when exiting the component production.
	ExitComponent(c *ComponentContext)

	// ExitSimpleComponent is called when exiting the simpleComponent production.
	ExitSimpleComponent(c *SimpleComponentContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)
}
