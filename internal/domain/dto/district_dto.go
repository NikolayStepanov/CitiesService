package dto

import (
	. "CitiesService/internal/domain/common"
	"reflect"
	"sort"
)

type DistrictDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DistrictsDTO []DistrictDTO

type DistrictsByID struct {
	DistrictsDTO
}

type DistrictsByName struct {
	DistrictsDTO
}

func (d DistrictsDTO) Len() int {
	return len(d)
}

func (d DistrictsDTO) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d DistrictsByID) Less(i, j int) bool {
	return d.DistrictsDTO[i].ID < d.DistrictsDTO[j].ID
}

func (d DistrictsByName) Less(i, j int) bool {
	return d.DistrictsDTO[i].Name < d.DistrictsDTO[j].Name
}

func (d DistrictsDTO) SortByField(nameField string) {
	switch nameField {
	case ID:
		sort.Sort(DistrictsByID{d})
	case NAME:
		sort.Sort(DistrictsByName{d})
	default:
		sort.Sort(DistrictsByID{d})
	}
}

func (d DistrictsDTO) Search(nameField string, val any) int {
	index := d.Len()
	x := reflect.ValueOf(val)
	switch nameField {
	case ID:
		index = sort.Search(d.Len(), func(i int) bool {
			return d[i].ID >= int(x.Int())
		})
	case NAME:
		index = sort.Search(d.Len(), func(i int) bool {
			return d[i].Name >= x.String()
		})
	}
	return index
}

func (d DistrictsDTO) SelectByID(districtsID ...int) DistrictsDTO {
	resultDistricts := DistrictsDTO{}
	d.SortByField(ID)
	for _, regionID := range districtsID {
		index := d.Search(ID, regionID)
		if index < d.Len() && d[index].ID == regionID {
			resultDistricts = append(resultDistricts, d[index])
		}
	}
	return resultDistricts
}

func (d DistrictsDTO) SelectByName(districtsName ...string) DistrictsDTO {
	resultDistricts := DistrictsDTO{}
	d.SortByField(NAME)
	for _, districtName := range districtsName {
		index := d.Search(NAME, districtName)
		if index < d.Len() && d[index].Name == districtName {
			resultDistricts = append(resultDistricts, d[index])
		}
	}
	return resultDistricts
}
