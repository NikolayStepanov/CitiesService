package domain

type Region struct {
	id   int
	name string
}

func NewRegion(id int, name string) *Region {
	return &Region{id: id, name: name}
}

func (r *Region) ID() int {
	return r.id
}

func (r *Region) SetID(id int) {
	r.id = id
}

func (r *Region) Name() string {
	return r.name
}

func (r *Region) SetName(name string) {
	r.name = name
}
