// Code generated from documentation/design/DesignLanguage.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // DesignLanguage
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DesignLanguageParser struct {
	*antlr.BaseParser
}

var DesignLanguageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func designlanguageParserInit() {
	staticData := &DesignLanguageParserStaticData
	staticData.LiteralNames = []string{
		"", "' '", "'()'", "'('", "')'", "'[]'", "'->'", "'*'", "", "", "",
		"", "'-- '",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "ARRAY", "ARROW", "ASTERISK", "FIELD_START", "AUTHOR_NAME",
		"AUTHOR_START", "COMMENT", "COMMENT_START", "COMMENT_TEXT", "NAME",
		"NEWLINE", "SPECIAL_CHAR",
	}
	staticData.RuleNames = []string{
		"design", "preamble", "author", "component", "simpleComponent", "field",
		"attribute", "method", "param", "params", "type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 115, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 3, 0, 24, 8, 0, 1, 0, 1, 0, 1, 0, 5, 0, 29, 8, 0, 10, 0, 12,
		0, 32, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 40, 8, 1, 1, 1,
		1, 1, 3, 1, 44, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 4, 3, 56, 8, 3, 11, 3, 12, 3, 57, 3, 3, 60, 8, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 66, 8, 4, 1, 5, 1, 5, 1, 5, 3, 5, 71, 8, 5, 1, 6, 1,
		6, 1, 6, 3, 6, 76, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		85, 8, 7, 1, 7, 1, 7, 3, 7, 89, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 3, 9, 98, 8, 9, 1, 9, 1, 9, 5, 9, 102, 8, 9, 10, 9, 12, 9, 105,
		9, 9, 1, 9, 3, 9, 108, 8, 9, 1, 10, 3, 10, 111, 8, 10, 1, 10, 1, 10, 1,
		10, 0, 0, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 0, 118, 0, 23,
		1, 0, 0, 0, 2, 43, 1, 0, 0, 0, 4, 47, 1, 0, 0, 0, 6, 59, 1, 0, 0, 0, 8,
		61, 1, 0, 0, 0, 10, 67, 1, 0, 0, 0, 12, 72, 1, 0, 0, 0, 14, 77, 1, 0, 0,
		0, 16, 90, 1, 0, 0, 0, 18, 107, 1, 0, 0, 0, 20, 110, 1, 0, 0, 0, 22, 24,
		3, 2, 1, 0, 23, 22, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 30, 1, 0, 0, 0,
		25, 26, 3, 6, 3, 0, 26, 27, 5, 15, 0, 0, 27, 29, 1, 0, 0, 0, 28, 25, 1,
		0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31,
		33, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 34, 5, 0, 0, 1, 34, 1, 1, 0, 0,
		0, 35, 36, 3, 4, 2, 0, 36, 39, 5, 15, 0, 0, 37, 38, 5, 11, 0, 0, 38, 40,
		5, 15, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 44, 1, 0, 0, 0,
		41, 42, 5, 11, 0, 0, 42, 44, 5, 15, 0, 0, 43, 35, 1, 0, 0, 0, 43, 41, 1,
		0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 5, 15, 0, 0, 46, 3, 1, 0, 0, 0, 47,
		48, 5, 10, 0, 0, 48, 49, 5, 9, 0, 0, 49, 5, 1, 0, 0, 0, 50, 60, 3, 8, 4,
		0, 51, 55, 3, 8, 4, 0, 52, 53, 3, 10, 5, 0, 53, 54, 5, 15, 0, 0, 54, 56,
		1, 0, 0, 0, 55, 52, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0,
		57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 50, 1, 0, 0, 0, 59, 51, 1,
		0, 0, 0, 60, 7, 1, 0, 0, 0, 61, 62, 5, 14, 0, 0, 62, 65, 5, 15, 0, 0, 63,
		64, 5, 11, 0, 0, 64, 66, 5, 15, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0,
		0, 0, 66, 9, 1, 0, 0, 0, 67, 70, 5, 8, 0, 0, 68, 71, 3, 12, 6, 0, 69, 71,
		3, 14, 7, 0, 70, 68, 1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 11, 1, 0, 0, 0,
		72, 75, 3, 16, 8, 0, 73, 74, 5, 1, 0, 0, 74, 76, 5, 11, 0, 0, 75, 73, 1,
		0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 13, 1, 0, 0, 0, 77, 78, 5, 14, 0, 0, 78,
		79, 5, 1, 0, 0, 79, 80, 3, 18, 9, 0, 80, 84, 5, 1, 0, 0, 81, 82, 5, 6,
		0, 0, 82, 83, 5, 1, 0, 0, 83, 85, 3, 18, 9, 0, 84, 81, 1, 0, 0, 0, 84,
		85, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 87, 5, 1, 0, 0, 87, 89, 5, 11,
		0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 15, 1, 0, 0, 0, 90, 91,
		5, 14, 0, 0, 91, 92, 5, 1, 0, 0, 92, 93, 3, 20, 10, 0, 93, 17, 1, 0, 0,
		0, 94, 108, 5, 2, 0, 0, 95, 97, 5, 3, 0, 0, 96, 98, 3, 16, 8, 0, 97, 96,
		1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 103, 1, 0, 0, 0, 99, 100, 5, 1, 0,
		0, 100, 102, 3, 16, 8, 0, 101, 99, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105, 103,
		1, 0, 0, 0, 106, 108, 5, 4, 0, 0, 107, 94, 1, 0, 0, 0, 107, 95, 1, 0, 0,
		0, 108, 19, 1, 0, 0, 0, 109, 111, 5, 5, 0, 0, 110, 109, 1, 0, 0, 0, 110,
		111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 5, 14, 0, 0, 113, 21,
		1, 0, 0, 0, 15, 23, 30, 39, 43, 57, 59, 65, 70, 75, 84, 88, 97, 103, 107,
		110,
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

// DesignLanguageParserInit initializes any static state used to implement DesignLanguageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDesignLanguageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DesignLanguageParserInit() {
	staticData := &DesignLanguageParserStaticData
	staticData.once.Do(designlanguageParserInit)
}

// NewDesignLanguageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDesignLanguageParser(input antlr.TokenStream) *DesignLanguageParser {
	DesignLanguageParserInit()
	this := new(DesignLanguageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DesignLanguageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "DesignLanguage.g4"

	return this
}

// DesignLanguageParser tokens.
const (
	DesignLanguageParserEOF           = antlr.TokenEOF
	DesignLanguageParserT__0          = 1
	DesignLanguageParserT__1          = 2
	DesignLanguageParserT__2          = 3
	DesignLanguageParserT__3          = 4
	DesignLanguageParserARRAY         = 5
	DesignLanguageParserARROW         = 6
	DesignLanguageParserASTERISK      = 7
	DesignLanguageParserFIELD_START   = 8
	DesignLanguageParserAUTHOR_NAME   = 9
	DesignLanguageParserAUTHOR_START  = 10
	DesignLanguageParserCOMMENT       = 11
	DesignLanguageParserCOMMENT_START = 12
	DesignLanguageParserCOMMENT_TEXT  = 13
	DesignLanguageParserNAME          = 14
	DesignLanguageParserNEWLINE       = 15
	DesignLanguageParserSPECIAL_CHAR  = 16
)

// DesignLanguageParser rules.
const (
	DesignLanguageParserRULE_design          = 0
	DesignLanguageParserRULE_preamble        = 1
	DesignLanguageParserRULE_author          = 2
	DesignLanguageParserRULE_component       = 3
	DesignLanguageParserRULE_simpleComponent = 4
	DesignLanguageParserRULE_field           = 5
	DesignLanguageParserRULE_attribute       = 6
	DesignLanguageParserRULE_method          = 7
	DesignLanguageParserRULE_param           = 8
	DesignLanguageParserRULE_params          = 9
	DesignLanguageParserRULE_type            = 10
)

// IDesignContext is an interface to support dynamic dispatch.
type IDesignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	Preamble() IPreambleContext
	AllComponent() []IComponentContext
	Component(i int) IComponentContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsDesignContext differentiates from other interfaces.
	IsDesignContext()
}

type DesignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignContext() *DesignContext {
	var p = new(DesignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_design
	return p
}

func InitEmptyDesignContext(p *DesignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_design
}

func (*DesignContext) IsDesignContext() {}

func NewDesignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignContext {
	var p = new(DesignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_design

	return p
}

func (s *DesignContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignContext) EOF() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserEOF, 0)
}

func (s *DesignContext) Preamble() IPreambleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPreambleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPreambleContext)
}

func (s *DesignContext) AllComponent() []IComponentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComponentContext); ok {
			len++
		}
	}

	tst := make([]IComponentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComponentContext); ok {
			tst[i] = t.(IComponentContext)
			i++
		}
	}

	return tst
}

func (s *DesignContext) Component(i int) IComponentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentContext); ok {
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

	return t.(IComponentContext)
}

func (s *DesignContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(DesignLanguageParserNEWLINE)
}

func (s *DesignContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserNEWLINE, i)
}

func (s *DesignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterDesign(s)
	}
}

func (s *DesignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitDesign(s)
	}
}

func (p *DesignLanguageParser) Design() (localctx IDesignContext) {
	localctx = NewDesignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DesignLanguageParserRULE_design)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DesignLanguageParserAUTHOR_START || _la == DesignLanguageParserCOMMENT {
		{
			p.SetState(22)
			p.Preamble()
		}

	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DesignLanguageParserNAME {
		{
			p.SetState(25)
			p.Component()
		}
		{
			p.SetState(26)
			p.Match(DesignLanguageParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Match(DesignLanguageParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPreambleContext is an interface to support dynamic dispatch.
type IPreambleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Author() IAuthorContext
	COMMENT() antlr.TerminalNode

	// IsPreambleContext differentiates from other interfaces.
	IsPreambleContext()
}

type PreambleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreambleContext() *PreambleContext {
	var p = new(PreambleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_preamble
	return p
}

func InitEmptyPreambleContext(p *PreambleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_preamble
}

func (*PreambleContext) IsPreambleContext() {}

func NewPreambleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreambleContext {
	var p = new(PreambleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_preamble

	return p
}

func (s *PreambleContext) GetParser() antlr.Parser { return s.parser }

func (s *PreambleContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(DesignLanguageParserNEWLINE)
}

func (s *PreambleContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserNEWLINE, i)
}

func (s *PreambleContext) Author() IAuthorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorContext)
}

func (s *PreambleContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserCOMMENT, 0)
}

func (s *PreambleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreambleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreambleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterPreamble(s)
	}
}

func (s *PreambleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitPreamble(s)
	}
}

func (p *DesignLanguageParser) Preamble() (localctx IPreambleContext) {
	localctx = NewPreambleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DesignLanguageParserRULE_preamble)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DesignLanguageParserAUTHOR_START:
		{
			p.SetState(35)
			p.Author()
		}
		{
			p.SetState(36)
			p.Match(DesignLanguageParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == DesignLanguageParserCOMMENT {
			{
				p.SetState(37)
				p.Match(DesignLanguageParserCOMMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(38)
				p.Match(DesignLanguageParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case DesignLanguageParserCOMMENT:
		{
			p.SetState(41)
			p.Match(DesignLanguageParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Match(DesignLanguageParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(45)
		p.Match(DesignLanguageParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAuthorContext is an interface to support dynamic dispatch.
type IAuthorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AUTHOR_START() antlr.TerminalNode
	AUTHOR_NAME() antlr.TerminalNode

	// IsAuthorContext differentiates from other interfaces.
	IsAuthorContext()
}

type AuthorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuthorContext() *AuthorContext {
	var p = new(AuthorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_author
	return p
}

func InitEmptyAuthorContext(p *AuthorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_author
}

func (*AuthorContext) IsAuthorContext() {}

func NewAuthorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuthorContext {
	var p = new(AuthorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_author

	return p
}

func (s *AuthorContext) GetParser() antlr.Parser { return s.parser }

func (s *AuthorContext) AUTHOR_START() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserAUTHOR_START, 0)
}

func (s *AuthorContext) AUTHOR_NAME() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserAUTHOR_NAME, 0)
}

func (s *AuthorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuthorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuthorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterAuthor(s)
	}
}

func (s *AuthorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitAuthor(s)
	}
}

func (p *DesignLanguageParser) Author() (localctx IAuthorContext) {
	localctx = NewAuthorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DesignLanguageParserRULE_author)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(DesignLanguageParserAUTHOR_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(DesignLanguageParserAUTHOR_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComponentContext is an interface to support dynamic dispatch.
type IComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleComponent() ISimpleComponentContext
	AllField() []IFieldContext
	Field(i int) IFieldContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsComponentContext differentiates from other interfaces.
	IsComponentContext()
}

type ComponentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentContext() *ComponentContext {
	var p = new(ComponentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_component
	return p
}

func InitEmptyComponentContext(p *ComponentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_component
}

func (*ComponentContext) IsComponentContext() {}

func NewComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentContext {
	var p = new(ComponentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_component

	return p
}

func (s *ComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentContext) SimpleComponent() ISimpleComponentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleComponentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleComponentContext)
}

func (s *ComponentContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *ComponentContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *ComponentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(DesignLanguageParserNEWLINE)
}

func (s *ComponentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserNEWLINE, i)
}

func (s *ComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterComponent(s)
	}
}

func (s *ComponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitComponent(s)
	}
}

func (p *DesignLanguageParser) Component() (localctx IComponentContext) {
	localctx = NewComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DesignLanguageParserRULE_component)
	var _la int

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.SimpleComponent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.SimpleComponent()
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DesignLanguageParserFIELD_START {
			{
				p.SetState(52)
				p.Field()
			}
			{
				p.SetState(53)
				p.Match(DesignLanguageParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleComponentContext is an interface to support dynamic dispatch.
type ISimpleComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	COMMENT() antlr.TerminalNode

	// IsSimpleComponentContext differentiates from other interfaces.
	IsSimpleComponentContext()
}

type SimpleComponentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleComponentContext() *SimpleComponentContext {
	var p = new(SimpleComponentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_simpleComponent
	return p
}

func InitEmptySimpleComponentContext(p *SimpleComponentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_simpleComponent
}

func (*SimpleComponentContext) IsSimpleComponentContext() {}

func NewSimpleComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleComponentContext {
	var p = new(SimpleComponentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_simpleComponent

	return p
}

func (s *SimpleComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleComponentContext) NAME() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserNAME, 0)
}

func (s *SimpleComponentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(DesignLanguageParserNEWLINE)
}

func (s *SimpleComponentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserNEWLINE, i)
}

func (s *SimpleComponentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserCOMMENT, 0)
}

func (s *SimpleComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleComponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterSimpleComponent(s)
	}
}

func (s *SimpleComponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitSimpleComponent(s)
	}
}

func (p *DesignLanguageParser) SimpleComponent() (localctx ISimpleComponentContext) {
	localctx = NewSimpleComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DesignLanguageParserRULE_simpleComponent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(DesignLanguageParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Match(DesignLanguageParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DesignLanguageParserCOMMENT {
		{
			p.SetState(63)
			p.Match(DesignLanguageParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(DesignLanguageParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FIELD_START() antlr.TerminalNode
	Attribute() IAttributeContext
	Method() IMethodContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FIELD_START() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserFIELD_START, 0)
}

func (s *FieldContext) Attribute() IAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *FieldContext) Method() IMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *DesignLanguageParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DesignLanguageParserRULE_field)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(DesignLanguageParserFIELD_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(68)
			p.Attribute()
		}

	case 2:
		{
			p.SetState(69)
			p.Method()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Param() IParamContext
	COMMENT() antlr.TerminalNode

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_attribute
	return p
}

func InitEmptyAttributeContext(p *AttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_attribute
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) Param() IParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *AttributeContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserCOMMENT, 0)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *DesignLanguageParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DesignLanguageParserRULE_attribute)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Param()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DesignLanguageParserT__0 {
		{
			p.SetState(73)
			p.Match(DesignLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(DesignLanguageParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	AllParams() []IParamsContext
	Params(i int) IParamsContext
	ARROW() antlr.TerminalNode
	COMMENT() antlr.TerminalNode

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_method
	return p
}

func InitEmptyMethodContext(p *MethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_method
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) NAME() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserNAME, 0)
}

func (s *MethodContext) AllParams() []IParamsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamsContext); ok {
			len++
		}
	}

	tst := make([]IParamsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamsContext); ok {
			tst[i] = t.(IParamsContext)
			i++
		}
	}

	return tst
}

func (s *MethodContext) Params(i int) IParamsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
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

	return t.(IParamsContext)
}

func (s *MethodContext) ARROW() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserARROW, 0)
}

func (s *MethodContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserCOMMENT, 0)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *DesignLanguageParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DesignLanguageParserRULE_method)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(DesignLanguageParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Match(DesignLanguageParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Params()
	}
	{
		p.SetState(80)
		p.Match(DesignLanguageParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DesignLanguageParserARROW {
		{
			p.SetState(81)
			p.Match(DesignLanguageParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(DesignLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Params()
		}

	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DesignLanguageParserT__0 {
		{
			p.SetState(86)
			p.Match(DesignLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(DesignLanguageParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	Type_() ITypeContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) NAME() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserNAME, 0)
}

func (s *ParamContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *DesignLanguageParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DesignLanguageParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(DesignLanguageParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Match(DesignLanguageParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllParam() []IParamContext {
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

func (s *ParamsContext) Param(i int) IParamContext {
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

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitParams(s)
	}
}

func (p *DesignLanguageParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DesignLanguageParserRULE_params)
	var _la int

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DesignLanguageParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Match(DesignLanguageParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DesignLanguageParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Match(DesignLanguageParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == DesignLanguageParserNAME {
			{
				p.SetState(96)
				p.Param()
			}

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DesignLanguageParserT__0 {
			{
				p.SetState(99)
				p.Match(DesignLanguageParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(100)
				p.Param()
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(106)
			p.Match(DesignLanguageParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	ARRAY() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DesignLanguageParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignLanguageParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) NAME() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserNAME, 0)
}

func (s *TypeContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(DesignLanguageParserARRAY, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignLanguageListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *DesignLanguageParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DesignLanguageParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DesignLanguageParserARRAY {
		{
			p.SetState(109)
			p.Match(DesignLanguageParserARRAY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(112)
		p.Match(DesignLanguageParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
