package designlanguage

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	dl "github.com/mattmunz/designlanguage/model/designlanguage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewObject(t *testing.T) {
	name := "Shape"
	component := dl.NewObject(name, nil, nil)
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
			assertAttributeEqual(t, entity.Attributes()[0], "Version", "String", false)
			assertAttributeEqual(t, entity.Attributes()[1], "CommandHandlers", "CommandHandler", true)
			continue
		default:
			require.Fail(t, "Unknown entity: %s", entity.Name())
		}
	}

	for _, object := range design.Objects() {
		switch object.Name() {
		case "CommandHandler":
			assert.Equal(t, 1, len(object.Attributes()))
			assertAttributeEqual(t, object.Attributes()[0], "Description", "String", false)

			assert.Equal(t, 2, len(object.Methods()))

			for _, method := range object.Methods() {
				switch method.Name() {
				case "Interpet":
					require.Equal(t, 1, len(method.Params()))
					assertAttributeEqual(t, method.Params()[0], "Command", "String", false)
					require.Equal(t, 1, len(method.ReturnVals()))
					assertAttributeEqual(t, method.ReturnVals()[0], "Response", "String", false)
				case "GetConfigs":
					require.Equal(t, 1, len(method.Params()))
					assertAttributeEqual(t, method.Params()[0], "Time", "DateTime", false)
					require.Equal(t, 1, len(method.ReturnVals()))
					assertAttributeEqual(t, method.ReturnVals()[0], "Configs", "String", true)
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
			assertAttributeEqual(t, method.Params()[0], "ID", "String", false)
			require.Equal(t, 1, len(method.ReturnVals()))
			assertAttributeEqual(t, method.ReturnVals()[0], "Element", "Element", false)
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

	param := method.Params()[0]
	assertAttributeEqual(t, param, "Args", "String", false)
}

func getDesign(t *testing.T, testFileSubpath string, namespace string, expectedComponentsLength int, expectedEntitiesLength int, expectedObjectsLength int) dl.Design {
	absPath, err := filepath.Abs(
		filepath.Join(getTestDataPath(t), testFileSubpath),
	)
	require.NoError(t, err)

	design := dl.Parse(absPath, namespace)

	require.Equal(t, expectedComponentsLength, len(design.Components()))
	require.Equal(t, expectedEntitiesLength, len(design.Entities()))
	require.Equal(t, expectedObjectsLength, len(design.Objects()))
	return design
}

func assertAttributeEqual(t *testing.T, attr dl.Attribute, expectedName string, expectedType string, expectedIsArray bool) {
	require.Equal(t, expectedName, attr.Name())
	require.Equal(t, expectedType, attr.Type().Name())
	require.Equal(t, expectedIsArray, attr.Type().Array())
}

func getTestDataPath(t *testing.T) string {
	value := os.Getenv("TEST_DATA_PATH")
	if value == "" {
		require.Fail(t, "TEST_DATA_PATH is not set")
	}
	return value
}
