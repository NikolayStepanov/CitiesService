package dto

import (
	. "CitiesService/internal/domain/common"
	"reflect"
	"sort"
)

type CityDTO struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	RegionID   int    `json:"region_id"`
	DistrictID int    `json:"district_id"`
	Population int    `json:"population,omitempty"`
	Foundation int    `json:"foundation,omitempty"`
}

type CitiesDTO []CityDTO
type CitiesByID struct{ CitiesDTO }
type CitiesByName struct{ CitiesDTO }
type CitiesByRegionID struct{ CitiesDTO }
type CitiesByDistrictID struct{ CitiesDTO }
type CitiesByPopulation struct{ CitiesDTO }
type CitiesByFoundation struct{ CitiesDTO }

func (c CitiesDTO) Len() int {
	return len(c)
}
func (c CitiesDTO) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CitiesByID) Less(i, j int) bool   { return c.CitiesDTO[i].ID < c.CitiesDTO[j].ID }
func (c CitiesByName) Less(i, j int) bool { return c.CitiesDTO[i].Name < c.CitiesDTO[j].Name }

func (c CitiesByRegionID) Less(i, j int) bool {
	return c.CitiesDTO[i].RegionID < c.CitiesDTO[j].RegionID
}

func (c CitiesByDistrictID) Less(i, j int) bool {
	return c.CitiesDTO[i].DistrictID < c.CitiesDTO[j].DistrictID
}

func (c CitiesByPopulation) Less(i, j int) bool {
	return c.CitiesDTO[i].Population < c.CitiesDTO[j].Population
}

func (c CitiesByFoundation) Less(i, j int) bool {
	return c.CitiesDTO[i].Foundation < c.CitiesDTO[j].Foundation
}

func (c CitiesDTO) SortByField(nameField string) {
	switch nameField {
	case ID:
		sort.Sort(CitiesByID{c})
	case NAME:
		sort.Sort(CitiesByName{c})
	case REGION_ID:
		sort.Sort(CitiesByRegionID{c})
	case DISTRICT_ID:
		sort.Sort(CitiesByDistrictID{c})
	case POPULATION:
		sort.Sort(CitiesByPopulation{c})
	case FOUNDATION:
		sort.Sort(CitiesByFoundation{c})
	default:
		sort.Sort(CitiesByID{c})
	}
}

func (c CitiesDTO) Search(nameField string, val any) int {
	index := c.Len()
	x := reflect.ValueOf(val)
	switch nameField {
	case ID:
		index = sort.Search(c.Len(), func(i int) bool {
			return c[i].ID >= int(x.Int())
		})
	case NAME:
		index = sort.Search(c.Len(), func(i int) bool {
			return c[i].Name >= x.String()
		})
	case REGION_ID:
		index = sort.Search(c.Len(), func(i int) bool {
			return c[i].RegionID >= int(x.Int())
		})
	case DISTRICT_ID:
		index = sort.Search(c.Len(), func(i int) bool {
			return c[i].DistrictID >= int(x.Int())
		})
	case POPULATION:
		index = sort.Search(c.Len(), func(i int) bool {
			return c[i].Population >= int(x.Int())
		})
	case FOUNDATION:
		index = sort.Search(c.Len(), func(i int) bool {
			return c[i].Foundation >= int(x.Int())
		})
	default:
		index = sort.Search(c.Len(), func(i int) bool {
			return c[i].ID >= int(x.Int())
		})
	}
	return index
}

func (c CitiesDTO) SelectByID(citiesID ...int) CitiesDTO {
	resultCities := CitiesDTO{}
	c.SortByField(ID)
	for _, cityID := range citiesID {
		index := c.Search(ID, cityID)
		if index < c.Len() && c[index].ID == cityID {
			resultCities = append(resultCities, c[index])
		}
	}
	return resultCities
}

func (c CitiesDTO) SelectByName(citiesName ...string) CitiesDTO {
	resultCities := CitiesDTO{}
	c.SortByField(NAME)
	for _, cityName := range citiesName {
		index := c.Search(NAME, cityName)
		if index < c.Len() && c[index].Name == cityName {
			resultCities = append(resultCities, c[index])
		}
	}
	return resultCities
}

func (c CitiesDTO) SelectByRegionID(regionsID ...int) CitiesDTO {
	resultCities := CitiesDTO{}
	c.SortByField(REGION_ID)
	for _, regionID := range regionsID {
		indexStart := c.Search(REGION_ID, regionID)
		if indexStart < c.Len() && c[indexStart].RegionID == regionID {
			indexEnd := c.Len()
			for i := indexStart; i < c.Len(); i++ {
				if c[i].RegionID != regionID {
					indexEnd = i
					break
				}
			}
			resultCities = append(resultCities, c[indexStart:indexEnd]...)
		}
	}
	return resultCities
}

func (c CitiesDTO) SelectByDistrictID(districtsID ...int) CitiesDTO {
	resultCities := CitiesDTO{}
	c.SortByField(DISTRICT_ID)
	for _, districtID := range districtsID {
		indexStart := c.Search(DISTRICT_ID, districtID)
		if indexStart < c.Len() && c[indexStart].DistrictID == districtID {
			indexEnd := c.Len()
			for i := indexStart; i < c.Len(); i++ {
				if c[i].DistrictID != districtID {
					indexEnd = i
					break
				}
			}
			resultCities = append(resultCities, c[indexStart:indexEnd]...)
		}
	}
	return resultCities
}

func (c CitiesDTO) SelectRangePopulation(minPopulation, maxPopulation int) CitiesDTO {
	resultCities := CitiesDTO{}
	c.SortByField(POPULATION)
	indexStart := c.Search(POPULATION, minPopulation)
	if indexStart < c.Len() && c[indexStart].Population >= minPopulation {
		indexEnd := c.Len()
		for i := indexStart; i < c.Len(); i++ {
			if c[i].Population > maxPopulation {
				indexEnd = i
				break
			}
		}
		resultCities = append(resultCities, c[indexStart:indexEnd]...)
	}
	return resultCities
}

func (c CitiesDTO) SelectRangeFoundation(minFoundation, maxFoundation int) CitiesDTO {
	resultCities := CitiesDTO{}
	c.SortByField(FOUNDATION)
	indexStart := c.Search(FOUNDATION, minFoundation)
	if indexStart < c.Len() && c[indexStart].Foundation >= minFoundation {
		indexEnd := c.Len()
		for i := indexStart; i < c.Len(); i++ {
			if c[i].Foundation > maxFoundation {
				indexEnd = i
				break
			}
		}
		resultCities = append(resultCities, c[indexStart:indexEnd]...)
	}
	return resultCities
}
