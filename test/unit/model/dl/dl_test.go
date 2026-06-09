package dl

import (
	"testing"

	"github.com/mattmunz/designlanguage/gen"
	"github.com/mattmunz/designlanguage/model"
	"github.com/stretchr/testify/require"
)

var plane = model.NewComponent("Plane", "")

var rectangle = model.NewEntity(
	"Rectangle", "", []model.Attribute{model.NewAttribute("Length", "", "int", false), model.NewAttribute("Width", "", "int", false)},
)

var labelledValue = model.NewEntity(
	"LabelledValue", "", []model.Attribute{model.NewAttribute("Label", "", "String", false), model.NewAttribute("Value", "", "Int", false)},
)

var personRepo = model.NewObject(
	"PersonRepository", "", []model.Attribute{},
	[]model.Method{
		model.NewMethod("Add", "", []model.Param{model.NewParam("Person", model.NewType("Person", false))}, []model.Param{}),
		model.NewMethod("Get", "",
			[]model.Param{model.NewParam("name", model.NewType("String", false))},
			[]model.Param{model.NewParam("person", model.NewType("Person", false))}),
	})

func TestRenderDesignSummary(t *testing.T) {
	spinner := newSpinner()

	design := model.NewDesign("", "",
		"Geometry", []model.Component{plane, rectangle, spinner}, []model.Entity{rectangle}, []model.Object{spinner},
	)

	source := gen.RenderDesignSummary(design)

	require.Equal(t, "Namespace: Geometry, Components: 3, Entities: 1, Objects: 1", source)
}

func newSpinner() model.Object {
	intType := model.NewType("int", false)

	method := []model.Method{model.NewMethod("Spin", "",
		[]model.Param{model.NewParam("velocity", intType), model.NewParam("duration", intType)},
		[]model.Param{}),
	}

	return model.NewObject(
		"Spinner", "", []model.Attribute{model.NewAttribute("Radius", "", "int", false)},
		method,
	)
}

func TestRenderEmptyComponent(t *testing.T) {
	expectedSource := "type Plane interface{}"

	src, err := gen.RenderComponentSource(plane)

	require.NoError(t, err)
	require.Equal(t, expectedSource, src)
}

func TestRenderEntity(t *testing.T) {
	expectedSource := `type Rectangle interface {
	Length() int
	Width() int
}`

	src, err := gen.RenderEntitySource(rectangle)

	require.NoError(t, err)
	require.Equal(t, expectedSource, src)

}

// TestTypeAlias verifies that type names are rendered to their correlates in Go types when needed.
func TestTypeAlias(t *testing.T) {
	expectedSource := `type LabelledValue interface {
	Label() string
	Value() int
}`

	src, err := gen.RenderEntitySource(labelledValue)

	require.NoError(t, err)
	require.Equal(t, expectedSource, src)
}

func TestRenderObject(t *testing.T) {
	expectedSource := `type Spinner interface {
	Radius() int
	Spin(velocity int, duration int)
}`

	src, err := gen.RenderObjectSource(newSpinner())

	require.NoError(t, err)
	require.Equal(t, expectedSource, src)
}

func TestRenderPersonRepository(t *testing.T) {
	expectedSource := `type PersonRepository interface {
	Add(person Person)
	Get(name string) (person Person)
}`

	src, err := gen.RenderObjectSource(personRepo)

	require.NoError(t, err)
	require.Equal(t, expectedSource, src)
}

func TestRenderDesignSourceGeometry(t *testing.T) {
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

	design := model.NewDesign(
		"geometry", "", "", []model.Component{plane, rectangle, spinner}, []model.Entity{rectangle}, []model.Object{spinner},
	)

	src, err := gen.RenderDesignSource(design)
	require.NoError(t, err)
	require.Equal(t, expectedSource, src)
}

// TODO Make sure the following ends up in the source code for the appkit example
// See the above. Break out unit tests as needed
// TODO Test author
// TODO Test package documentation
// TODO Test inline documentation
// TODO Test method with no arrow
// TODO Test Logger mapper
// TODO Test commandImpl mapper
func TestRenderDesignSourceAppkit(t *testing.T) {
	require.Fail(t, "NYI")
}
