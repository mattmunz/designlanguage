// Code generated from documentation/design/DesignLanguage.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type DesignLanguageLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var DesignLanguageLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func designlanguagelexerLexerInit() {
	staticData := &DesignLanguageLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "' '", "'()'", "'('", "')'", "'[]'", "", "'*'", "'->'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "ARRAY", "NAME", "ASTERISK", "ARROW", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "ARRAY", "NAME", "ASTERISK", "ARROW",
		"NEWLINE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 48, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5,
		5, 34, 8, 5, 10, 5, 12, 5, 37, 9, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8,
		3, 8, 45, 8, 8, 1, 8, 1, 8, 0, 0, 9, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 1, 0, 2, 1, 0, 65, 90, 2, 0, 65, 90, 97, 122, 49,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0, 3, 21, 1, 0, 0, 0, 5, 24, 1, 0,
		0, 0, 7, 26, 1, 0, 0, 0, 9, 28, 1, 0, 0, 0, 11, 31, 1, 0, 0, 0, 13, 38,
		1, 0, 0, 0, 15, 40, 1, 0, 0, 0, 17, 44, 1, 0, 0, 0, 19, 20, 5, 32, 0, 0,
		20, 2, 1, 0, 0, 0, 21, 22, 5, 40, 0, 0, 22, 23, 5, 41, 0, 0, 23, 4, 1,
		0, 0, 0, 24, 25, 5, 40, 0, 0, 25, 6, 1, 0, 0, 0, 26, 27, 5, 41, 0, 0, 27,
		8, 1, 0, 0, 0, 28, 29, 5, 91, 0, 0, 29, 30, 5, 93, 0, 0, 30, 10, 1, 0,
		0, 0, 31, 35, 7, 0, 0, 0, 32, 34, 7, 1, 0, 0, 33, 32, 1, 0, 0, 0, 34, 37,
		1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 12, 1, 0, 0, 0,
		37, 35, 1, 0, 0, 0, 38, 39, 5, 42, 0, 0, 39, 14, 1, 0, 0, 0, 40, 41, 5,
		45, 0, 0, 41, 42, 5, 62, 0, 0, 42, 16, 1, 0, 0, 0, 43, 45, 5, 13, 0, 0,
		44, 43, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 5,
		10, 0, 0, 47, 18, 1, 0, 0, 0, 3, 0, 35, 44, 0,
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

// DesignLanguageLexerInit initializes any static state used to implement DesignLanguageLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewDesignLanguageLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func DesignLanguageLexerInit() {
	staticData := &DesignLanguageLexerLexerStaticData
	staticData.once.Do(designlanguagelexerLexerInit)
}

// NewDesignLanguageLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewDesignLanguageLexer(input antlr.CharStream) *DesignLanguageLexer {
	DesignLanguageLexerInit()
	l := new(DesignLanguageLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &DesignLanguageLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "DesignLanguage.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DesignLanguageLexer tokens.
const (
	DesignLanguageLexerT__0     = 1
	DesignLanguageLexerT__1     = 2
	DesignLanguageLexerT__2     = 3
	DesignLanguageLexerT__3     = 4
	DesignLanguageLexerARRAY    = 5
	DesignLanguageLexerNAME     = 6
	DesignLanguageLexerASTERISK = 7
	DesignLanguageLexerARROW    = 8
	DesignLanguageLexerNEWLINE  = 9
)
