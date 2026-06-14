package unit

import (
	"testing"

	"github.com/mattmunz/designlanguage/model"
	"github.com/stretchr/testify/require"
)

var intType = model.NewType("Int", false)

func NewPlane(t *testing.T) model.Component {
	p, err := model.NewComponent("Plane", "", "")
	require.NoError(t, err)
	return p
}

func NewRectangle(t *testing.T) model.Entity {
	rectangle, err := model.NewEntity(
		"Rectangle", "", "", []model.Attribute{model.NewAttribute("Length", "", "int", false), model.NewAttribute("Width", "", "int", false)},
	)
	require.NoError(t, err)
	return rectangle
}

func NewSquare(t *testing.T) model.Entity {
	square, err := model.NewEntity(
		"Square", "", "Rectangle", []model.Attribute{model.NewAttribute("Area", "", "int", false)},
	)
	require.NoError(t, err)
	return square
}

func NewLabelledValue(t *testing.T) model.Entity {
	labelledValue, err := model.NewEntity(
		"LabelledValue", "", "", []model.Attribute{model.NewAttribute("Label", "", "String", false), model.NewAttribute("Value", "", "Int", false)},
	)
	require.NoError(t, err)
	return labelledValue
}

func NewPersonRepo(t *testing.T) model.Object {
	addMethod, err := model.NewMethod("Add", "", []model.Param{model.NewParam("Person", model.NewType("Person", false))}, []model.Param{})
	require.NoError(t, err)

	getMethod, err := model.NewMethod("Get", "",
		[]model.Param{model.NewParam("name", model.NewType("String", false))},
		[]model.Param{model.NewParam("person", model.NewType("Person", false))})
	require.NoError(t, err)

	personRepo, err := model.NewObject(
		"PersonRepository", "", "", []model.Attribute{},
		[]model.Method{
			addMethod,
			getMethod,
		})
	require.NoError(t, err)
	return personRepo
}

func NewSpinner(t *testing.T) model.Object {
	method, err := model.NewMethod("Spin", "",
		[]model.Param{model.NewParam("velocity", intType), model.NewParam("duration", intType)},
		[]model.Param{})
	require.NoError(t, err)

	o, err := model.NewObject(
		"Spinner", "", "", []model.Attribute{model.NewAttribute("Radius", "", "int", false)},
		[]model.Method{method},
	)
	require.NoError(t, err)

	return o
}

func NewSpinningSquare(t *testing.T) model.Object {
	method, err := model.NewMethod("Spin", "",
		[]model.Param{model.NewParam("Velocity", intType), model.NewParam("Duration", intType)},
		[]model.Param{})
	require.NoError(t, err)

	o, err := model.NewObject(
		"SpinningSquare", "", "Square", []model.Attribute{},
		[]model.Method{method},
	)
	require.NoError(t, err)

	return o
}
