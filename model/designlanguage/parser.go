package designlanguage

import (
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"

	p "github.com/mattmunz/designlanguage/model/designlanguage/parser"
)

func Parse(path, namespace string) Design {
	// TODO Can get namespace from path, or from aa new namespace decl. in the design doc.
	// Maybe add a mapping file that maps paths to namespace strings. Maybe can put other namespace metadata in this file. Like a package.java file
	lexer := p.NewDesignLanguageLexer(antlr.NewInputStream(string(readFile(path))))
	parser := p.NewDesignLanguageParser(antlr.NewCommonTokenStream(lexer, 0))

	components := []Component{}
	entities := []Entity{}
	objects := []Object{}
	for _, cc := range parser.Design().AllComponent() {
		objects, components, entities = addComponent(cc, objects, components, entities)
	}

	return NewDesign(namespace, components, entities, objects)
}

func addComponent(cc p.IComponentContext, objects []Object, components []Component, entities []Entity) ([]Object, []Component, []Entity) {
	name := cc.NAME().GetText()
	fcs := cc.AllField()
	attributes := []Attribute{}
	methods := []Method{}
	for _, fc := range fcs {
		attr := fc.Attribute()
		if attr != nil {
			attributes = addAttribute(attr, attributes)
			continue
		}

		method := fc.Method()
		if method == nil {
			continue
		}

		methods = addMethod(method, methods)
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
	return objects, components, entities
}

func addAttribute(attr p.IAttributeContext, attributes []Attribute) []Attribute {
	if attr == nil {
		panic("TODO Nil attr")
	}

	attrName, typeName, isArray := GetParam(attr.Param())
	return append(attributes, NewAttribute(attrName, typeName, isArray))
}

func addMethod(method p.IMethodContext, methods []Method) []Method {
	methodName := method.NAME().GetText()
	pcs := method.AllParams()
	if len(pcs) != 2 {
		panic("Expected 2 params")
	}

	methodObj := NewMethod(methodName, GetParams(pcs[0]), GetParams(pcs[1]))
	return append(methods, methodObj)
}

func GetParams(params p.IParamsContext) []Param {
	paramCount := len(params.AllParam())

	if paramCount < 1 {
		return []Param{}
	}

	paramsList := make([]Param, paramCount)
	for i, param := range params.AllParam() {
		name, typeI, isArray := GetParam(param)
		// TODO Add newParamForParam method
		paramsList[i] = NewParam(name, NewType(typeI, isArray))
	}

	return paramsList
}

// TODO Modify to return Param Obj. Param Obj shouldl accept empty list or something?
func GetParam(param p.IParamContext) (string, string, bool) {
	if param == nil {
		panic("TODO Nil param")
	}

	if param.NAME() == nil {
		panic("TODO Nil param name NYI")
	}

	attrName := param.NAME().GetText()
	attrType := param.Type_()

	if attrType == nil {
		panic("TODO Nil attrType")
	}

	typeName := attrType.NAME().GetText()

	typeArray := attrType.ARRAY()
	if typeArray == nil {
		return attrName, typeName, false
	}
	arrayText := typeArray.GetText()

	isArray := len(arrayText) > 0
	return attrName, typeName, isArray
}

func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		// TODO
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		// TODO
		panic(err)
	}
	return content
}
