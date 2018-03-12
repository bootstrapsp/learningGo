package main

// GenericThing struct
type GenericThing struct {
	id             int64
	name           string
	classification string
}

// LiveThing struct with embedded GenericThing type
type LiveThing struct {
	GenericThing
	isHomosapein bool
}

// OtherThing struct with embedded GenericThing
type OtherThing struct {
	GenericThing
	isOrganic bool
}

func (o OtherThing) EntityCheck() bool {

	var confirmCheck bool

	return confirmCheck
}
