package geometry

type Person interface {
	Name() string
}

type Client interface {
	User() Person
	PersonRepository() PersonRepository
}

type Rectangle interface {
	Length() int
	Width() int
}

type PersonRepository interface {
	Add(person Person)
	Get(name string) (person Person)
}
