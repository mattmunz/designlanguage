package model

type Design interface {
	Commented

	Author() string
	Namespace() string
	Components() []Component
	Entities() []Entity
	Objects() []Object
}

type Named interface {
	Name() string
}

type named struct {
	name string
}

func (n *named) Name() string {
	return n.name
}

func newNamed(name string) Named {
	return &named{name}
}

type Param interface {
	Named
	Type() Type
}

type param struct {
	Named
	typ Type
}

func (p *param) Type() Type {
	return p.typ
}

func NewParam(name string, type1 Type) Param {
	return &param{
		newNamed(name),
		type1,
	}
}

type Commented interface {
	Comment() string
}

type Type interface {
	Named
	IsArray() bool
}

// Component is a computational object that may be composed of other components.
// It exposes its functionality through a single interface.
type Component interface {
	Named
	Commented
}

type component struct {
	Named
	Commented
}

func NewComponent(name, comment string) Component {
	return &component{
		newNamed(name),
		newCommented(comment),
	}
}

// Attribute is the named part of an interface, designating a sub-component.
type Attribute interface {
	Commented
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

func NewObject(name, comment string, attributes []Attribute, methods []Method) Object {
	return &object{
		NewEntity(name, comment, attributes),
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

func NewEntity(name, comment string, attributes []Attribute) Entity {
	return &entity{NewComponent(name, comment), attributes}
}

type entity struct {
	Component
	attributes []Attribute
}

func (e *entity) Attributes() []Attribute {
	return e.attributes
}

type commented struct {
	comment string
}

func (c *commented) Comment() string {
	return c.comment
}

func newCommented(commentText string) Commented {
	return &commented{commentText}
}

// design is a concrete implementation of the Design interface
type design struct {
	Commented
	author     string
	namespace  string
	components []Component
	entities   []Entity
	objects    []Object
}

func (d *design) Author() string {
	return d.author
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

func NewDesign(author, comment, namespace string, components []Component, entities []Entity, objects []Object) Design {
	return &design{
		newCommented(comment),
		author,
		namespace,
		components,
		entities,
		objects,
	}
}

// attribute is a concrete implementation of the Attribute interface
type attribute struct {
	Named
	Commented
	typ Type
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

func NewAttribute(name, comment, typeName string, isArray bool) Attribute {
	return &attribute{
		newNamed(name),
		newCommented(comment),
		NewType(typeName, isArray),
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
	Component
	params     []Param
	returnVals []Param
}

func (m *method) Params() []Param {
	return m.params
}

func (m *method) ReturnVals() []Param {
	return m.returnVals
}

func NewMethod(name, comment string, params, returnVals []Param) Method {
	return &method{
		NewComponent(name, comment),
		params,
		returnVals,
	}
}
