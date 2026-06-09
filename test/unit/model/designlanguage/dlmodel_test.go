package designlanguage

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/mattmunz/designlanguage/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewObject(t *testing.T) {
	name := "Shape"
	component := model.NewObject(name, "", nil, nil)
	require.NotNil(t, component)
	require.Equal(t, name, component.Name())
}

func TestLoad1(t *testing.T) {
	design := getDesign(t, "Test.1.nzsd.txt", "ns1", 4, 2, 2)

	for _, entity := range design.Entities() {
		switch entity.Name() {
		case "User":
			require.Equal(t, 0, len(entity.Attributes()))
			continue
		case "Console":
			require.Equal(t, 2, len(entity.Attributes()))
			assertAttributeEqual(t, entity.Attributes()[0], "Version", "", "String", false)
			assertAttributeEqual(t, entity.Attributes()[1], "CommandHandlers", "", "CommandHandler", true)
			continue
		default:
			require.Fail(t, "Unknown entity: %s", entity.Name())
		}
	}

	for _, object := range design.Objects() {
		switch object.Name() {
		case "CommandHandler":
			assert.Equal(t, 1, len(object.Attributes()))
			assertAttributeEqual(t, object.Attributes()[0], "Description", "", "String", false)

			assert.Equal(t, 2, len(object.Methods()))

			for _, method := range object.Methods() {
				switch method.Name() {
				case "Interpet":
					require.Equal(t, 1, len(method.Params()))
					assertParamEqual(t, method.Params()[0], "Command", "String", false)
					require.Equal(t, 1, len(method.ReturnVals()))
					assertParamEqual(t, method.ReturnVals()[0], "Response", "String", false)
				case "GetConfigs":
					require.Equal(t, 1, len(method.Params()))
					assertParamEqual(t, method.Params()[0], "Time", "DateTime", false)
					require.Equal(t, 1, len(method.ReturnVals()))
					assertParamEqual(t, method.ReturnVals()[0], "Configs", "String", true)
				default:
					require.Fail(t, "Unknown method: %s", method.Name())
				}
			}

		case "Browser":
			require.Equal(t, 0, len(object.Attributes()))
			assert.Equal(t, 1, len(object.Methods()))
			method := object.Methods()[0]

			require.Equal(t, "GetElement", method.Name())
			require.Equal(t, 1, len(method.Params()))
			assertParamEqual(t, method.Params()[0], "ID", "String", false)
			require.Equal(t, 1, len(method.ReturnVals()))
			assertParamEqual(t, method.ReturnVals()[0], "Element", "Element", false)
		default:
			require.Fail(t, fmt.Sprintf("Unknown object: %s", object.Name()))
		}
	}
}

func TestLoad3(t *testing.T) {
	design := getDesign(t, "Test.3.nzsd.txt", "ns3", 1, 0, 1)

	obj := design.Objects()[0]

	require.Equal(t, "Runner", obj.Name())
	assert.Equal(t, 1, len(obj.Methods()))

	method := obj.Methods()[0]

	require.Equal(t, "Run", method.Name())
	require.Equal(t, 1, len(method.Params()))
	require.Equal(t, 0, len(method.ReturnVals()))

	assertParamEqual(t, method.Params()[0], "Args", "String", false)
}

func TestLoad4(t *testing.T) {
	design := getDesign(t, "documentation/documentation/design/Appkit.nzsd.txt", "appkit", 3, 1, 2)

	require.Equal(t, "Nonzero Sum", design.Author())
	require.Equal(t, "Appkit is an application development kit, part of the Nonzero Sum Stack.", design.Comment())

	require.Len(t, design.Entities(), 1)
	entity := design.Entities()[0]
	require.Equal(t, "A software application.", entity.Comment())

	require.Len(t, entity.Attributes(), 3)

	attr2 := entity.Attributes()[1]
	require.Equal(t, "Version", attr2.Name())
	require.Equal(t, "Semver preferred.", attr2.Comment())

	obj1 := design.Objects()[1]
	require.Equal(t, "Command", obj1.Name())
	require.Len(t, obj1.Methods(), 1)
	method1 := obj1.Methods()[0]
	// TODO Extract RequireNameEqual()
	require.Equal(t, "Command", obj1.Name())
	require.Equal(t, "Execute", method1.Name())
	require.Len(t, method1.Params(), 0)
	require.Len(t, method1.ReturnVals(), 0)

	obj2 := design.Objects()[1]
	require.Equal(t, "CLI", obj2.Name())
	require.Len(t, obj2.Methods(), 2)

	method21 := obj2.Methods()[0]

	require.Equal(t, "NewRootCommand", method21.Name())
	require.Equal(t, 2, len(method21.Params()))

	param1 := method21.Params()[0]
	require.Equal(t, "Cmd", param1.Name())
	require.Equal(t, "CommandImpl", param1.Type().Name())

	method22 := obj2.Methods()[1]

	require.Equal(t, "SetLogger", method22.Name())
	require.Equal(t, 1, len(method22.Params()))
	param21 := method22.Params()[0]
	require.Equal(t, "Logger", param21.Name())
	require.Equal(t, "Logger", param21.Type().Name())
}

func getDesign(t *testing.T, testFileSubpath string, namespace string, expectedComponentsLength int, expectedEntitiesLength int, expectedObjectsLength int) model.Design {
	absPath, err := filepath.Abs(
		filepath.Join(getTestDataPath(t), testFileSubpath),
	)
	require.NoError(t, err)

	design, err := model.Parse(absPath, namespace)
	require.NoError(t, err)

	require.Equal(t, expectedComponentsLength, len(design.Components()))
	require.Equal(t, expectedEntitiesLength, len(design.Entities()))
	require.Equal(t, expectedObjectsLength, len(design.Objects()))
	return design
}

func assertAttributeEqual(t *testing.T, attr model.Attribute, expectedName, expectedComment, expectedType string, expectedIsArray bool) {
	require.Equal(t, expectedName, attr.Name())
	require.Equal(t, expectedComment, attr.Comment())
	require.Equal(t, expectedType, attr.Type().Name())
	require.Equal(t, expectedIsArray, attr.Type().IsArray())
}

// TODO Attr and Param are similar. Refactoring opportunity?
func assertParamEqual(t *testing.T, param model.Param, expectedName string, expectedType string, expectedIsArray bool) {
	require.Equal(t, expectedName, param.Name())
	require.Equal(t, expectedType, param.Type().Name())
	require.Equal(t, expectedIsArray, param.Type().IsArray())
}

func getTestDataPath(t *testing.T) string {
	value := os.Getenv("TEST_DATA_PATH")
	if value == "" {
		require.Fail(t, "TEST_DATA_PATH is not set")
	}
	return value
}
