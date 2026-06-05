package gen

import (
	"testing"

	"github.com/stretchr/testify/require"

	g "github.com/mattmunz/designlanguage/gen"
)

func TestParseDLMFilePath1(t *testing.T) {
	ns, fn, err := g.ParseDLMFilePath("/foo/bar/projects/baz/documentation/design", "/foo/bar/projects/baz/documentation/design/Thing1.Design.md")

	require.NoError(t, err)
	require.Equal(t, "", ns)
	require.Equal(t, "Thing1.Design.md", fn)
}

func TestParseDLMFilePath2(t *testing.T) {
	ns, fn, err := g.ParseDLMFilePath("/foo/bar/projects/baz/documentation/design", "/foo/bar/projects/baz/documentation/design/universe/Thing1.Design.md")

	require.NoError(t, err)
	require.Equal(t, "universe", ns)
	require.Equal(t, "Thing1.Design.md", fn)
}

func TestParseDLMFilePath3(t *testing.T) {
	ns, fn, err := g.ParseDLMFilePath("/foo/bar/projects/baz/documentation/design", "/foo/bar/projects/baz/documentation/design/universe/sol/Earth.Design.md")

	require.NoError(t, err)
	require.Equal(t, "universe/sol", ns)
	require.Equal(t, "Earth.Design.md", fn)
}
