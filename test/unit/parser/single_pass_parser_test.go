package parser

import (
	"testing"

	"github.com/mattmunz/designlanguage/model"
	"github.com/mattmunz/designlanguage/parser"
	"github.com/stretchr/testify/require"
)

func TestParseMethod1(t *testing.T) {
	method, err := model.NewMethod("Foo", "", []model.Param{}, []model.Param{})
	require.NoError(t, err)
	requireParsedMethodEqual(t, "Foo ()", method)
}

func TestParseMethod2(t *testing.T) {
	params := []model.Param{model.NewParam("Bar", model.NewType("Baz", false))}
	method, err := model.NewMethod("Foo", "", params, []model.Param{})
	require.NoError(t, err)
	requireParsedMethodEqual(t, "Foo (Bar Baz)", method)
}

func TestParseMethod3(t *testing.T) {
	method, err := model.NewMethod("Foo", "", []model.Param{}, []model.Param{})
	require.NoError(t, err)
	requireParsedMethodEqual(t, "Foo () -> ()", method)
}

func TestParseMethod4(t *testing.T) {
	params := []model.Param{model.NewParam("Bar", model.NewType("Baz", false))}
	returnVals := []model.Param{model.NewParam("This", model.NewType("That", false))}
	method, err := model.NewMethod("Foo", "", params, returnVals)
	require.NoError(t, err)
	requireParsedMethodEqual(t, "Foo (Bar Baz) -> (This That)", method)
}

func TestParseMethod5(t *testing.T) {
	params := []model.Param{model.NewParam("Bar", model.NewType("Baz", false))}
	returnVals := []model.Param{model.NewParam("This", model.NewType("That", false))}
	method, err := model.NewMethod("Foo", "A comment!", params, returnVals)
	require.NoError(t, err)
	requireParsedMethodEqual(t, "Foo (Bar Baz) -> (This That) -- A comment!", method)
}

func TestParseMethod6(t *testing.T) {
	params := []model.Param{
		model.NewParam("Bar", model.NewType("Baz", false)),
		model.NewParam("This", model.NewType("That", false)),
	}
	returnVals := []model.Param{model.NewParam("Which", model.NewType("What", false))}
	method, err := model.NewMethod("Foo", "", params, returnVals)
	require.NoError(t, err)
	requireParsedMethodEqual(t, "Foo (Bar Baz, This That) -> (Which What)", method)
}

func requireParsedMethodEqual(t *testing.T, text string, expectedMethod model.Method) {
	method, err := parser.ParseMethod(text)
	require.NoError(t, err)

	require.Equal(t, expectedMethod.Name(), method.Name())
	require.Equal(t, expectedMethod.Params(), method.Params())
	require.Equal(t, expectedMethod.ReturnVals(), method.ReturnVals())
}

func TestIsComponentLine1(t *testing.T) {
	require.False(t, parser.IsComponentLine1(""), "1")
	require.False(t, parser.IsComponentLine1(" "), "2")
	require.False(t, parser.IsComponentLine1(" Foo"), "3")
	require.True(t, parser.IsComponentLine1("Foo"), "4")
	require.False(t, parser.IsComponentLine1("foo"), "5")
	require.True(t, parser.IsComponentLine1("Foo"), "6")
	require.False(t, parser.IsComponentLine1("Foo::"), "7")
	require.False(t, parser.IsComponentLine1("Foo ::"), "8")
	require.True(t, parser.IsComponentLine1("Foo :: BarBaz"), "9")
	require.False(t, parser.IsComponentLine1("Foo :: barBaz"), "10")
}
