package model

import (
	"fmt"
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"

	p "github.com/mattmunz/designlanguage/model/designlanguage/parser"
)

func Parse(path, namespace string) (Design, error) {
	fileBytes, err := readFile(path)
	if err != nil {
		return nil, err
	}
	lexer := p.NewDesignLanguageLexer(antlr.NewInputStream(string(fileBytes)))
	parser := p.NewDesignLanguageParser(antlr.NewCommonTokenStream(lexer, 0))

	components := []Component{}
	entities := []Entity{}
	objects := []Object{}
	for _, cc := range parser.Design().AllComponent() {
		objects, components, entities, err = addComponent(cc, objects, components, entities)
		if err != nil {
			return nil, err
		}
	}

	return NewDesign(namespace, components, entities, objects), nil
}

func addComponent(cc p.IComponentContext, objects []Object, components []Component, entities []Entity) ([]Object, []Component, []Entity, error) {
	name := cc.NAME().GetText()
	fcs := cc.AllField()
	attributes := []Attribute{}
	methods := []Method{}
	var err error
	for _, fc := range fcs {
		attr := fc.Attribute()
		if attr != nil {
			attributes, err = addAttribute(attr, attributes)
			if err != nil {
				return nil, nil, nil, err
			}

			continue
		}

		method := fc.Method()
		if method == nil {
			continue
		}

		methods, err = addMethod(method, methods)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	if len(methods) > 0 {
		obj := NewObject(name, attributes, methods)
		objects = append(objects, obj)
		components = append(components, obj)
	} else {
		ent := NewEntity(name, attributes)
		entities = append(entities, ent)
		components = append(components, ent)
	}
	return objects, components, entities, nil
}

func addAttribute(attr p.IAttributeContext, attributes []Attribute) ([]Attribute, error) {
	param, err := GetParam(attr.Param())
	if err != nil {
		return nil, err
	}

	return append(attributes, NewAttribute(param.Name(), param.Type().Name(), param.Type().IsArray())), nil
}

func addMethod(method p.IMethodContext, methods []Method) ([]Method, error) {
	methodName := method.NAME().GetText()
	pcs := method.AllParams()
	if len(pcs) != 2 {
		return methods, fmt.Errorf("Expected 2 params")
	}

	param1, err := GetParams(pcs[0])
	if err != nil {
		return nil, err
	}

	param2, err := GetParams(pcs[1])
	if err != nil {
		return nil, err
	}

	return append(methods, NewMethod(methodName, param1, param2)), nil
}

func GetParams(params p.IParamsContext) ([]Param, error) {
	paramCount := len(params.AllParam())

	if paramCount < 1 {
		return []Param{}, nil
	}

	paramsList := make([]Param, paramCount)
	for i, param := range params.AllParam() {
		newParam, err := GetParam(param)
		if err != nil {
			return nil, err
		}

		paramsList[i] = newParam
	}

	return paramsList, nil
}

func GetParam(param p.IParamContext) (Param, error) {
	if param.NAME() == nil {
		return nil, fmt.Errorf("Expected param name")
	}

	attrName := param.NAME().GetText()

	attrType := param.Type_()
	if attrType == nil {
		return nil, fmt.Errorf("Expected param type")
	}

	typeName := attrType.NAME().GetText()

	typeArray := attrType.ARRAY()
	if typeArray == nil {
		return NewParam(attrName, NewType(typeName, false)), nil
	}

	isArray := len(typeArray.GetText()) > 0
	return NewParam(attrName, NewType(typeName, isArray)), nil
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}
