package domain

type District struct {
	id   int
	name string
}

func NewDistrict(id int, name string) *District {
	return &District{id: id, name: name}
}

func (d *District) ID() int {
	return d.id
}

func (d *District) SetID(id int) {
	d.id = id
}

func (d *District) Name() string {
	return d.name
}

func (d *District) SetName(name string) {
	d.name = name
}
