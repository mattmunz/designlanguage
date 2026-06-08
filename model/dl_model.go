package model

type Design interface {
	Namespace() string
	Components() []Component
	Entities() []Entity
	Objects() []Object
}

type Named interface {
	Name() string
}

type Param interface {
	Named
	Type() Type
}

type param struct {
	name string
	typ  Type
}

func (p *param) Name() string {
	return p.name
}

func (p *param) Type() Type {
	return p.typ
}

func NewParam(name string, type1 Type) Param {
	return &param{
		name: name,
		typ:  type1,
	}
}

type Type interface {
	Named
	IsArray() bool
}

// Component is a computational object that may be composed of other components.
// It exposes its functionality through a single interface.
type Component interface {
	Named
}

// Attribute is the named part of an interface, designating a sub-component.
type Attribute interface {
	Param
}

// Method is a means by which a component computes.
// It's a function from some entity to some entity.
type Method interface {
	Named
	Params() []Param
	ReturnVals() []Param
}

// Entity is a special type of Component with no methods, and whose
// attributes are all entities.
// It embeds the Component interface to inherit its base functionality.
type Entity interface {
	Component
	Attributes() []Attribute
}

type Object interface {
	Entity
	Methods() []Method
}

// Representation is an encoding of an entity.
type Representation interface {
	// Encode converts the entity into an encoded string representation.
	Encode() string
	// Decode converts an encoded string back into an entity.
	Decode(encoded string) (Entity, error)
}

func NewObject(name string, attributes []Attribute, methods []Method) Object {
	return &object{
		NewEntity(name, attributes),
		methods,
	}
}

type object struct {
	Entity
	methods []Method
}

func (o *object) Methods() []Method {
	return o.methods
}

func NewEntity(name string, attributes []Attribute) Entity {
	return &entity{name, attributes}
}

type entity struct {
	name       string
	attributes []Attribute
}

func (e *entity) Name() string {
	return e.name
}

func (e *entity) Attributes() []Attribute {
	return e.attributes
}

// design is a concrete implementation of the Design interface
type design struct {
	namespace  string
	components []Component
	entities   []Entity
	objects    []Object
}

func (d *design) Namespace() string {
	return d.namespace
}

func (d *design) Components() []Component {
	return d.components
}

func (d *design) Entities() []Entity {
	return d.entities
}

func (d *design) Objects() []Object {
	return d.objects
}

func NewDesign(namespace string, components []Component, entities []Entity, objects []Object) Design {
	return &design{
		namespace,
		components,
		entities,
		objects,
	}
}

// attribute is a concrete implementation of the Attribute interface
type attribute struct {
	name string
	typ  Type
}

func (a *attribute) Name() string {
	return a.name
}

func (a *attribute) Type() Type {
	return a.typ
}

// typ is a concrete implementation of the Type interface
type typ struct {
	name  string
	array bool
}

func (t *typ) Name() string {
	return t.name
}

func (t *typ) IsArray() bool {
	return t.array
}

func NewAttribute(name, typeName string, isArray bool) Attribute {
	return &attribute{
		name: name,
		typ:  NewType(typeName, isArray),
	}
}

func NewType(name string, isArray bool) Type {
	return &typ{
		name:  name,
		array: isArray,
	}
}

// method is a concrete implementation of the Method interface
type method struct {
	name       string
	params     []Param
	returnVals []Param
}

func (m *method) Name() string {
	return m.name
}

func (m *method) Params() []Param {
	return m.params
}

func (m *method) ReturnVals() []Param {
	return m.returnVals
}

func NewMethod(name string, params, returnVals []Param) Method {
	return &method{
		name:       name,
		params:     params,
		returnVals: returnVals,
	}
}
