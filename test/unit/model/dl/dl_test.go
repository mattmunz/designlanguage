package dl

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mattmunz/designlanguage/model/designlanguage"
	dl1 "github.com/mattmunz/designlanguage/model/dl"
)

// TODO MAke test for antlr parsing a document with a method with no arrow.

var plane = dl1.NewComponent("Plane")

var rectangle = designlanguage.NewEntity(
	"Rectangle", []designlanguage.Attribute{designlanguage.NewAttribute("Length", "int", false), designlanguage.NewAttribute("Width", "int", false)},
)

var labelledValue = designlanguage.NewEntity(
	"LabelledValue", []designlanguage.Attribute{designlanguage.NewAttribute("Label", "String", false), designlanguage.NewAttribute("Value", "Int", false)},
)

var personRepo = designlanguage.NewObject(
	"PersonRepository", []designlanguage.Attribute{},
	[]designlanguage.Method{
		designlanguage.NewMethod("Add", []designlanguage.Param{designlanguage.NewParam("Person", designlanguage.NewType("Person", false))}, []designlanguage.Param{}),
		designlanguage.NewMethod("Get",
			[]designlanguage.Param{designlanguage.NewParam("name", designlanguage.NewType("String", false))},
			[]designlanguage.Param{designlanguage.NewParam("person", designlanguage.NewType("Person", false))}),
	})

func TestRenderDesign(t *testing.T) {
	spinner := newSpinner()

	design := designlanguage.NewDesign(
		"Geometry", []designlanguage.Component{plane, rectangle, spinner}, []designlanguage.Entity{rectangle}, []designlanguage.Object{spinner},
	)

	source := dl1.RenderDesignSummary(design)

	require.Equal(t, "Namespace: Geometry, Components: 3, Entities: 1, Objects: 1", source)
}

func newSpinner() designlanguage.Object {
	intType := designlanguage.NewType("int", false)

	method := []designlanguage.Method{designlanguage.NewMethod("Spin",
		[]designlanguage.Param{designlanguage.NewParam("velocity", intType), designlanguage.NewParam("duration", intType)},
		[]designlanguage.Param{}),
	}

	return designlanguage.NewObject(
		"Spinner", []designlanguage.Attribute{designlanguage.NewAttribute("Radius", "int", false)},
		method,
	)
}

func TestRenderEmptyComponent(t *testing.T) {
	expectedSource := "type Plane interface{}"

	src, err := dl1.RenderComponentSource(plane)

	require.NoError(t, err)
	require.Equal(t, expectedSource, src)
}

func TestRenderEntity(t *testing.T) {
	expectedSource := `type Rectangle interface {
	Length() int
	Width() int
}`

	src, err := dl1.RenderEntitySource(rectangle)

	require.NoError(t, err)
	require.Equal(t, expectedSource, src)

}

// TestTypeAlias verifies that type names are rendered to their correlates in Go types when needed.
func TestTypeAlias(t *testing.T) {
	expectedSource := `type LabelledValue interface {
	Label() string
	Value() int
}`

	src, err := dl1.RenderEntitySource(labelledValue)

	require.NoError(t, err)
	require.Equal(t, expectedSource, src)
}

func TestRenderObject(t *testing.T) {
	expectedSource := `type Spinner interface {
	Radius() int
	Spin(velocity int, duration int)
}`

	src, err := dl1.RenderObjectSource(newSpinner())

	require.NoError(t, err)
	require.Equal(t, expectedSource, src)
}

func TestRenderPersonRepository(t *testing.T) {
	expectedSource := `type PersonRepository interface {
	Add(person Person)
	Get(name string) (person Person)
}`

	src, err := dl1.RenderObjectSource(personRepo)

	require.NoError(t, err)
	require.Equal(t, expectedSource, src)
}

func TestRenderDesignSource(t *testing.T) {
	expectedSource := `package geometry

type Plane interface{}

type Rectangle interface {
	Length() int
	Width() int
}

type Spinner interface {
	Radius() int
	Spin(velocity int, duration int)
}
`
	spinner := newSpinner()

	design := designlanguage.NewDesign(
		"geometry", []designlanguage.Component{plane, rectangle, spinner}, []designlanguage.Entity{rectangle}, []designlanguage.Object{spinner},
	)

	src, err := dl1.RenderDesignSource(design)
	require.NoError(t, err)
	require.Equal(t, expectedSource, src)
}
