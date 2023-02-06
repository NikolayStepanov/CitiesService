package domain

type City struct {
	id         int
	regionID   int
	districtID int
	population int
	foundation uint16
	name       string
}

func NewCity(name string) *City {
	return &City{name: name}
}

func (c *City) ID() int {
	return c.id
}

func (c *City) SetID(id int) {
	c.id = id
}

func (c *City) Name() string {
	return c.name
}

func (c *City) SetName(name string) {
	c.name = name
}

func (c *City) RegionID() int {
	return c.regionID
}

func (c *City) SetRegionID(regionID int) {
	c.regionID = regionID
}

func (c *City) DistrictID() int {
	return c.districtID
}

func (c *City) SetDistrictID(districtID int) {
	c.districtID = districtID
}

func (c *City) Population() int {
	return c.population
}

func (c *City) SetPopulation(population int) {
	c.population = population
}

func (c *City) Foundation() uint16 {
	return c.foundation
}

func (c *City) SetFoundation(foundation uint16) {
	c.foundation = foundation
}
