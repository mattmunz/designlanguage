// Generated from /Users/mattmunz/Google Drive/Projects/Programming/gowork/src/github.com/mattmunz/designlanguage/documentation/design/DesignLanguage.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class DesignLanguageParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, ARRAY=5, ARROW=6, ASTERISK=7, FIELD_START=8, 
		AUTHOR_NAME=9, AUTHOR_START=10, COMMENT=11, COMMENT_START=12, COMMENT_TEXT=13, 
		NAME=14, NEWLINE=15, SPECIAL_CHAR=16;
	public static final int
		RULE_design = 0, RULE_preamble = 1, RULE_author = 2, RULE_component = 3, 
		RULE_simpleComponent = 4, RULE_field = 5, RULE_attribute = 6, RULE_method = 7, 
		RULE_param = 8, RULE_params = 9, RULE_type = 10;
	private static String[] makeRuleNames() {
		return new String[] {
			"design", "preamble", "author", "component", "simpleComponent", "field", 
			"attribute", "method", "param", "params", "type"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "' '", "'()'", "'('", "')'", "'[]'", "'->'", "'*'", null, null, 
			null, null, "'-- '"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "ARRAY", "ARROW", "ASTERISK", "FIELD_START", 
			"AUTHOR_NAME", "AUTHOR_START", "COMMENT", "COMMENT_START", "COMMENT_TEXT", 
			"NAME", "NEWLINE", "SPECIAL_CHAR"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "DesignLanguage.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public DesignLanguageParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DesignContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(DesignLanguageParser.EOF, 0); }
		public PreambleContext preamble() {
			return getRuleContext(PreambleContext.class,0);
		}
		public List<ComponentContext> component() {
			return getRuleContexts(ComponentContext.class);
		}
		public ComponentContext component(int i) {
			return getRuleContext(ComponentContext.class,i);
		}
		public List<TerminalNode> NEWLINE() { return getTokens(DesignLanguageParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(DesignLanguageParser.NEWLINE, i);
		}
		public DesignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_design; }
	}

	public final DesignContext design() throws RecognitionException {
		DesignContext _localctx = new DesignContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_design);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(23);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AUTHOR_START || _la==COMMENT) {
				{
				setState(22);
				preamble();
				}
			}

			setState(30);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NAME) {
				{
				{
				setState(25);
				component();
				setState(26);
				match(NEWLINE);
				}
				}
				setState(32);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(33);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PreambleContext extends ParserRuleContext {
		public List<TerminalNode> NEWLINE() { return getTokens(DesignLanguageParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(DesignLanguageParser.NEWLINE, i);
		}
		public AuthorContext author() {
			return getRuleContext(AuthorContext.class,0);
		}
		public TerminalNode COMMENT() { return getToken(DesignLanguageParser.COMMENT, 0); }
		public PreambleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_preamble; }
	}

	public final PreambleContext preamble() throws RecognitionException {
		PreambleContext _localctx = new PreambleContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_preamble);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case AUTHOR_START:
				{
				setState(35);
				author();
				setState(36);
				match(NEWLINE);
				setState(39);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMENT) {
					{
					setState(37);
					match(COMMENT);
					setState(38);
					match(NEWLINE);
					}
				}

				}
				break;
			case COMMENT:
				{
				{
				setState(41);
				match(COMMENT);
				setState(42);
				match(NEWLINE);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(45);
			match(NEWLINE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AuthorContext extends ParserRuleContext {
		public TerminalNode AUTHOR_START() { return getToken(DesignLanguageParser.AUTHOR_START, 0); }
		public TerminalNode AUTHOR_NAME() { return getToken(DesignLanguageParser.AUTHOR_NAME, 0); }
		public AuthorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_author; }
	}

	public final AuthorContext author() throws RecognitionException {
		AuthorContext _localctx = new AuthorContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_author);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			match(AUTHOR_START);
			setState(48);
			match(AUTHOR_NAME);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ComponentContext extends ParserRuleContext {
		public SimpleComponentContext simpleComponent() {
			return getRuleContext(SimpleComponentContext.class,0);
		}
		public List<FieldContext> field() {
			return getRuleContexts(FieldContext.class);
		}
		public FieldContext field(int i) {
			return getRuleContext(FieldContext.class,i);
		}
		public List<TerminalNode> NEWLINE() { return getTokens(DesignLanguageParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(DesignLanguageParser.NEWLINE, i);
		}
		public ComponentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_component; }
	}

	public final ComponentContext component() throws RecognitionException {
		ComponentContext _localctx = new ComponentContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_component);
		int _la;
		try {
			setState(59);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(50);
				simpleComponent();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(51);
				simpleComponent();
				setState(55); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(52);
					field();
					setState(53);
					match(NEWLINE);
					}
					}
					setState(57); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==FIELD_START );
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SimpleComponentContext extends ParserRuleContext {
		public TerminalNode NAME() { return getToken(DesignLanguageParser.NAME, 0); }
		public List<TerminalNode> NEWLINE() { return getTokens(DesignLanguageParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(DesignLanguageParser.NEWLINE, i);
		}
		public TerminalNode COMMENT() { return getToken(DesignLanguageParser.COMMENT, 0); }
		public SimpleComponentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleComponent; }
	}

	public final SimpleComponentContext simpleComponent() throws RecognitionException {
		SimpleComponentContext _localctx = new SimpleComponentContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_simpleComponent);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(61);
			match(NAME);
			setState(62);
			match(NEWLINE);
			setState(65);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMENT) {
				{
				setState(63);
				match(COMMENT);
				setState(64);
				match(NEWLINE);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FieldContext extends ParserRuleContext {
		public TerminalNode FIELD_START() { return getToken(DesignLanguageParser.FIELD_START, 0); }
		public AttributeContext attribute() {
			return getRuleContext(AttributeContext.class,0);
		}
		public MethodContext method() {
			return getRuleContext(MethodContext.class,0);
		}
		public FieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_field; }
	}

	public final FieldContext field() throws RecognitionException {
		FieldContext _localctx = new FieldContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_field);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(67);
			match(FIELD_START);
			setState(70);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				{
				setState(68);
				attribute();
				}
				break;
			case 2:
				{
				setState(69);
				method();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeContext extends ParserRuleContext {
		public ParamContext param() {
			return getRuleContext(ParamContext.class,0);
		}
		public TerminalNode COMMENT() { return getToken(DesignLanguageParser.COMMENT, 0); }
		public AttributeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attribute; }
	}

	public final AttributeContext attribute() throws RecognitionException {
		AttributeContext _localctx = new AttributeContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_attribute);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			param();
			setState(75);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(73);
				match(T__0);
				setState(74);
				match(COMMENT);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MethodContext extends ParserRuleContext {
		public TerminalNode NAME() { return getToken(DesignLanguageParser.NAME, 0); }
		public List<ParamsContext> params() {
			return getRuleContexts(ParamsContext.class);
		}
		public ParamsContext params(int i) {
			return getRuleContext(ParamsContext.class,i);
		}
		public TerminalNode ARROW() { return getToken(DesignLanguageParser.ARROW, 0); }
		public TerminalNode COMMENT() { return getToken(DesignLanguageParser.COMMENT, 0); }
		public MethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_method; }
	}

	public final MethodContext method() throws RecognitionException {
		MethodContext _localctx = new MethodContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_method);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(77);
			match(NAME);
			setState(78);
			match(T__0);
			setState(79);
			params();
			setState(80);
			match(T__0);
			setState(84);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ARROW) {
				{
				setState(81);
				match(ARROW);
				setState(82);
				match(T__0);
				setState(83);
				params();
				}
			}

			setState(88);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(86);
				match(T__0);
				setState(87);
				match(COMMENT);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParamContext extends ParserRuleContext {
		public TerminalNode NAME() { return getToken(DesignLanguageParser.NAME, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			match(NAME);
			setState(91);
			match(T__0);
			setState(92);
			type();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParamsContext extends ParserRuleContext {
		public List<ParamContext> param() {
			return getRuleContexts(ParamContext.class);
		}
		public ParamContext param(int i) {
			return getRuleContext(ParamContext.class,i);
		}
		public ParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_params; }
	}

	public final ParamsContext params() throws RecognitionException {
		ParamsContext _localctx = new ParamsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_params);
		int _la;
		try {
			setState(107);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
				enterOuterAlt(_localctx, 1);
				{
				setState(94);
				match(T__1);
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 2);
				{
				setState(95);
				match(T__2);
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==NAME) {
					{
					setState(96);
					param();
					}
				}

				setState(103);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0) {
					{
					{
					setState(99);
					match(T__0);
					setState(100);
					param();
					}
					}
					setState(105);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(106);
				match(T__3);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeContext extends ParserRuleContext {
		public TerminalNode NAME() { return getToken(DesignLanguageParser.NAME, 0); }
		public TerminalNode ARRAY() { return getToken(DesignLanguageParser.ARRAY, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(110);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ARRAY) {
				{
				setState(109);
				match(ARRAY);
				}
			}

			setState(112);
			match(NAME);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u0010s\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0001\u0000\u0003\u0000\u0018"+
		"\b\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0005\u0000\u001d\b\u0000"+
		"\n\u0000\f\u0000 \t\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0003\u0001(\b\u0001\u0001\u0001\u0001\u0001"+
		"\u0003\u0001,\b\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0004\u00038\b\u0003\u000b\u0003\f\u00039\u0003\u0003<\b\u0003\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004B\b\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0003\u0005G\b\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0003\u0006L\b\u0006\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007U\b"+
		"\u0007\u0001\u0007\u0001\u0007\u0003\u0007Y\b\u0007\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\t\u0001\t\u0001\t\u0003\tb\b\t\u0001\t\u0001\t\u0005"+
		"\tf\b\t\n\t\f\ti\t\t\u0001\t\u0003\tl\b\t\u0001\n\u0003\no\b\n\u0001\n"+
		"\u0001\n\u0001\n\u0000\u0000\u000b\u0000\u0002\u0004\u0006\b\n\f\u000e"+
		"\u0010\u0012\u0014\u0000\u0000v\u0000\u0017\u0001\u0000\u0000\u0000\u0002"+
		"+\u0001\u0000\u0000\u0000\u0004/\u0001\u0000\u0000\u0000\u0006;\u0001"+
		"\u0000\u0000\u0000\b=\u0001\u0000\u0000\u0000\nC\u0001\u0000\u0000\u0000"+
		"\fH\u0001\u0000\u0000\u0000\u000eM\u0001\u0000\u0000\u0000\u0010Z\u0001"+
		"\u0000\u0000\u0000\u0012k\u0001\u0000\u0000\u0000\u0014n\u0001\u0000\u0000"+
		"\u0000\u0016\u0018\u0003\u0002\u0001\u0000\u0017\u0016\u0001\u0000\u0000"+
		"\u0000\u0017\u0018\u0001\u0000\u0000\u0000\u0018\u001e\u0001\u0000\u0000"+
		"\u0000\u0019\u001a\u0003\u0006\u0003\u0000\u001a\u001b\u0005\u000f\u0000"+
		"\u0000\u001b\u001d\u0001\u0000\u0000\u0000\u001c\u0019\u0001\u0000\u0000"+
		"\u0000\u001d \u0001\u0000\u0000\u0000\u001e\u001c\u0001\u0000\u0000\u0000"+
		"\u001e\u001f\u0001\u0000\u0000\u0000\u001f!\u0001\u0000\u0000\u0000 \u001e"+
		"\u0001\u0000\u0000\u0000!\"\u0005\u0000\u0000\u0001\"\u0001\u0001\u0000"+
		"\u0000\u0000#$\u0003\u0004\u0002\u0000$\'\u0005\u000f\u0000\u0000%&\u0005"+
		"\u000b\u0000\u0000&(\u0005\u000f\u0000\u0000\'%\u0001\u0000\u0000\u0000"+
		"\'(\u0001\u0000\u0000\u0000(,\u0001\u0000\u0000\u0000)*\u0005\u000b\u0000"+
		"\u0000*,\u0005\u000f\u0000\u0000+#\u0001\u0000\u0000\u0000+)\u0001\u0000"+
		"\u0000\u0000,-\u0001\u0000\u0000\u0000-.\u0005\u000f\u0000\u0000.\u0003"+
		"\u0001\u0000\u0000\u0000/0\u0005\n\u0000\u000001\u0005\t\u0000\u00001"+
		"\u0005\u0001\u0000\u0000\u00002<\u0003\b\u0004\u000037\u0003\b\u0004\u0000"+
		"45\u0003\n\u0005\u000056\u0005\u000f\u0000\u000068\u0001\u0000\u0000\u0000"+
		"74\u0001\u0000\u0000\u000089\u0001\u0000\u0000\u000097\u0001\u0000\u0000"+
		"\u00009:\u0001\u0000\u0000\u0000:<\u0001\u0000\u0000\u0000;2\u0001\u0000"+
		"\u0000\u0000;3\u0001\u0000\u0000\u0000<\u0007\u0001\u0000\u0000\u0000"+
		"=>\u0005\u000e\u0000\u0000>A\u0005\u000f\u0000\u0000?@\u0005\u000b\u0000"+
		"\u0000@B\u0005\u000f\u0000\u0000A?\u0001\u0000\u0000\u0000AB\u0001\u0000"+
		"\u0000\u0000B\t\u0001\u0000\u0000\u0000CF\u0005\b\u0000\u0000DG\u0003"+
		"\f\u0006\u0000EG\u0003\u000e\u0007\u0000FD\u0001\u0000\u0000\u0000FE\u0001"+
		"\u0000\u0000\u0000G\u000b\u0001\u0000\u0000\u0000HK\u0003\u0010\b\u0000"+
		"IJ\u0005\u0001\u0000\u0000JL\u0005\u000b\u0000\u0000KI\u0001\u0000\u0000"+
		"\u0000KL\u0001\u0000\u0000\u0000L\r\u0001\u0000\u0000\u0000MN\u0005\u000e"+
		"\u0000\u0000NO\u0005\u0001\u0000\u0000OP\u0003\u0012\t\u0000PT\u0005\u0001"+
		"\u0000\u0000QR\u0005\u0006\u0000\u0000RS\u0005\u0001\u0000\u0000SU\u0003"+
		"\u0012\t\u0000TQ\u0001\u0000\u0000\u0000TU\u0001\u0000\u0000\u0000UX\u0001"+
		"\u0000\u0000\u0000VW\u0005\u0001\u0000\u0000WY\u0005\u000b\u0000\u0000"+
		"XV\u0001\u0000\u0000\u0000XY\u0001\u0000\u0000\u0000Y\u000f\u0001\u0000"+
		"\u0000\u0000Z[\u0005\u000e\u0000\u0000[\\\u0005\u0001\u0000\u0000\\]\u0003"+
		"\u0014\n\u0000]\u0011\u0001\u0000\u0000\u0000^l\u0005\u0002\u0000\u0000"+
		"_a\u0005\u0003\u0000\u0000`b\u0003\u0010\b\u0000a`\u0001\u0000\u0000\u0000"+
		"ab\u0001\u0000\u0000\u0000bg\u0001\u0000\u0000\u0000cd\u0005\u0001\u0000"+
		"\u0000df\u0003\u0010\b\u0000ec\u0001\u0000\u0000\u0000fi\u0001\u0000\u0000"+
		"\u0000ge\u0001\u0000\u0000\u0000gh\u0001\u0000\u0000\u0000hj\u0001\u0000"+
		"\u0000\u0000ig\u0001\u0000\u0000\u0000jl\u0005\u0004\u0000\u0000k^\u0001"+
		"\u0000\u0000\u0000k_\u0001\u0000\u0000\u0000l\u0013\u0001\u0000\u0000"+
		"\u0000mo\u0005\u0005\u0000\u0000nm\u0001\u0000\u0000\u0000no\u0001\u0000"+
		"\u0000\u0000op\u0001\u0000\u0000\u0000pq\u0005\u000e\u0000\u0000q\u0015"+
		"\u0001\u0000\u0000\u0000\u000f\u0017\u001e\'+9;AFKTXagkn";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}