// Generated from /Users/mattmunz/Google Drive/Projects/Programming/gowork/src/github.com/mattmunz/sundries/documentation/design/DesignLanguage.g4 by ANTLR 4.13.1
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
		T__0=1, T__1=2, T__2=3, T__3=4, ARRAY=5, NAME=6, ASTERISK=7, ARROW=8, 
		NEWLINE=9;
	public static final int
		RULE_design = 0, RULE_component = 1, RULE_field = 2, RULE_attribute = 3, 
		RULE_method = 4, RULE_param = 5, RULE_params = 6, RULE_type = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"design", "component", "field", "attribute", "method", "param", "params", 
			"type"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "' '", "'()'", "'('", "')'", "'[]'", null, "'*'", "'->'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "ARRAY", "NAME", "ASTERISK", "ARROW", "NEWLINE"
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
			setState(22);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NAME) {
				{
				{
				setState(16);
				component();
				setState(18);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==NEWLINE) {
					{
					setState(17);
					match(NEWLINE);
					}
				}

				}
				}
				setState(24);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(25);
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
	public static class ComponentContext extends ParserRuleContext {
		public TerminalNode NAME() { return getToken(DesignLanguageParser.NAME, 0); }
		public TerminalNode NEWLINE() { return getToken(DesignLanguageParser.NEWLINE, 0); }
		public List<FieldContext> field() {
			return getRuleContexts(FieldContext.class);
		}
		public FieldContext field(int i) {
			return getRuleContext(FieldContext.class,i);
		}
		public ComponentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_component; }
	}

	public final ComponentContext component() throws RecognitionException {
		ComponentContext _localctx = new ComponentContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_component);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(27);
			match(NAME);
			setState(28);
			match(NEWLINE);
			setState(32);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ASTERISK) {
				{
				{
				setState(29);
				field();
				}
				}
				setState(34);
				_errHandler.sync(this);
				_la = _input.LA(1);
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
		enterRule(_localctx, 4, RULE_field);
		try {
			setState(37);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(35);
				attribute();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(36);
				method();
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
	public static class AttributeContext extends ParserRuleContext {
		public TerminalNode ASTERISK() { return getToken(DesignLanguageParser.ASTERISK, 0); }
		public ParamContext param() {
			return getRuleContext(ParamContext.class,0);
		}
		public TerminalNode NEWLINE() { return getToken(DesignLanguageParser.NEWLINE, 0); }
		public AttributeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attribute; }
	}

	public final AttributeContext attribute() throws RecognitionException {
		AttributeContext _localctx = new AttributeContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_attribute);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			match(ASTERISK);
			setState(40);
			match(T__0);
			setState(41);
			param();
			setState(42);
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
	public static class MethodContext extends ParserRuleContext {
		public TerminalNode ASTERISK() { return getToken(DesignLanguageParser.ASTERISK, 0); }
		public TerminalNode NAME() { return getToken(DesignLanguageParser.NAME, 0); }
		public List<ParamsContext> params() {
			return getRuleContexts(ParamsContext.class);
		}
		public ParamsContext params(int i) {
			return getRuleContext(ParamsContext.class,i);
		}
		public TerminalNode ARROW() { return getToken(DesignLanguageParser.ARROW, 0); }
		public TerminalNode NEWLINE() { return getToken(DesignLanguageParser.NEWLINE, 0); }
		public MethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_method; }
	}

	public final MethodContext method() throws RecognitionException {
		MethodContext _localctx = new MethodContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_method);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			match(ASTERISK);
			setState(45);
			match(T__0);
			setState(46);
			match(NAME);
			setState(47);
			match(T__0);
			setState(48);
			params();
			setState(49);
			match(T__0);
			setState(50);
			match(ARROW);
			setState(51);
			match(T__0);
			setState(52);
			params();
			setState(53);
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
		enterRule(_localctx, 10, RULE_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(55);
			match(NAME);
			setState(56);
			match(T__0);
			setState(57);
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
		enterRule(_localctx, 12, RULE_params);
		int _la;
		try {
			setState(72);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
				enterOuterAlt(_localctx, 1);
				{
				setState(59);
				match(T__1);
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				match(T__2);
				setState(62);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==NAME) {
					{
					setState(61);
					param();
					}
				}

				setState(68);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0) {
					{
					{
					setState(64);
					match(T__0);
					setState(65);
					param();
					}
					}
					setState(70);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(71);
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
		enterRule(_localctx, 14, RULE_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(75);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ARRAY) {
				{
				setState(74);
				match(ARRAY);
				}
			}

			setState(77);
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
		"\u0004\u0001\tP\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0001"+
		"\u0000\u0001\u0000\u0003\u0000\u0013\b\u0000\u0005\u0000\u0015\b\u0000"+
		"\n\u0000\f\u0000\u0018\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0005\u0001\u001f\b\u0001\n\u0001\f\u0001\"\t\u0001"+
		"\u0001\u0002\u0001\u0002\u0003\u0002&\b\u0002\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006?\b\u0006\u0001\u0006"+
		"\u0001\u0006\u0005\u0006C\b\u0006\n\u0006\f\u0006F\t\u0006\u0001\u0006"+
		"\u0003\u0006I\b\u0006\u0001\u0007\u0003\u0007L\b\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0000\u0000\b\u0000\u0002\u0004\u0006\b\n\f\u000e\u0000"+
		"\u0000O\u0000\u0016\u0001\u0000\u0000\u0000\u0002\u001b\u0001\u0000\u0000"+
		"\u0000\u0004%\u0001\u0000\u0000\u0000\u0006\'\u0001\u0000\u0000\u0000"+
		"\b,\u0001\u0000\u0000\u0000\n7\u0001\u0000\u0000\u0000\fH\u0001\u0000"+
		"\u0000\u0000\u000eK\u0001\u0000\u0000\u0000\u0010\u0012\u0003\u0002\u0001"+
		"\u0000\u0011\u0013\u0005\t\u0000\u0000\u0012\u0011\u0001\u0000\u0000\u0000"+
		"\u0012\u0013\u0001\u0000\u0000\u0000\u0013\u0015\u0001\u0000\u0000\u0000"+
		"\u0014\u0010\u0001\u0000\u0000\u0000\u0015\u0018\u0001\u0000\u0000\u0000"+
		"\u0016\u0014\u0001\u0000\u0000\u0000\u0016\u0017\u0001\u0000\u0000\u0000"+
		"\u0017\u0019\u0001\u0000\u0000\u0000\u0018\u0016\u0001\u0000\u0000\u0000"+
		"\u0019\u001a\u0005\u0000\u0000\u0001\u001a\u0001\u0001\u0000\u0000\u0000"+
		"\u001b\u001c\u0005\u0006\u0000\u0000\u001c \u0005\t\u0000\u0000\u001d"+
		"\u001f\u0003\u0004\u0002\u0000\u001e\u001d\u0001\u0000\u0000\u0000\u001f"+
		"\"\u0001\u0000\u0000\u0000 \u001e\u0001\u0000\u0000\u0000 !\u0001\u0000"+
		"\u0000\u0000!\u0003\u0001\u0000\u0000\u0000\" \u0001\u0000\u0000\u0000"+
		"#&\u0003\u0006\u0003\u0000$&\u0003\b\u0004\u0000%#\u0001\u0000\u0000\u0000"+
		"%$\u0001\u0000\u0000\u0000&\u0005\u0001\u0000\u0000\u0000\'(\u0005\u0007"+
		"\u0000\u0000()\u0005\u0001\u0000\u0000)*\u0003\n\u0005\u0000*+\u0005\t"+
		"\u0000\u0000+\u0007\u0001\u0000\u0000\u0000,-\u0005\u0007\u0000\u0000"+
		"-.\u0005\u0001\u0000\u0000./\u0005\u0006\u0000\u0000/0\u0005\u0001\u0000"+
		"\u000001\u0003\f\u0006\u000012\u0005\u0001\u0000\u000023\u0005\b\u0000"+
		"\u000034\u0005\u0001\u0000\u000045\u0003\f\u0006\u000056\u0005\t\u0000"+
		"\u00006\t\u0001\u0000\u0000\u000078\u0005\u0006\u0000\u000089\u0005\u0001"+
		"\u0000\u00009:\u0003\u000e\u0007\u0000:\u000b\u0001\u0000\u0000\u0000"+
		";I\u0005\u0002\u0000\u0000<>\u0005\u0003\u0000\u0000=?\u0003\n\u0005\u0000"+
		">=\u0001\u0000\u0000\u0000>?\u0001\u0000\u0000\u0000?D\u0001\u0000\u0000"+
		"\u0000@A\u0005\u0001\u0000\u0000AC\u0003\n\u0005\u0000B@\u0001\u0000\u0000"+
		"\u0000CF\u0001\u0000\u0000\u0000DB\u0001\u0000\u0000\u0000DE\u0001\u0000"+
		"\u0000\u0000EG\u0001\u0000\u0000\u0000FD\u0001\u0000\u0000\u0000GI\u0005"+
		"\u0004\u0000\u0000H;\u0001\u0000\u0000\u0000H<\u0001\u0000\u0000\u0000"+
		"I\r\u0001\u0000\u0000\u0000JL\u0005\u0005\u0000\u0000KJ\u0001\u0000\u0000"+
		"\u0000KL\u0001\u0000\u0000\u0000LM\u0001\u0000\u0000\u0000MN\u0005\u0006"+
		"\u0000\u0000N\u000f\u0001\u0000\u0000\u0000\b\u0012\u0016 %>DHK";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}