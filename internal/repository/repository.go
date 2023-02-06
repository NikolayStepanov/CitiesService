package repository

type Repositories struct {
	Cities    Cities
	Regions   Regions
	Districts Districts
}

func NewRepositories(cities Cities, regions Regions, districts Districts) *Repositories {
	return &Repositories{Cities: cities, Regions: regions, Districts: districts}
}

func (s *Repositories) Init() {
	s.Cities.Init()
	s.Regions.Init()
	s.Districts.Init()
}
