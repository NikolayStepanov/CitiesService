package dto

import (
	. "CitiesService/internal/domain/common"
	"reflect"
	"sort"
)

type RegionDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RegionsDTO []RegionDTO

type RegionsByID struct {
	RegionsDTO
}

type RegionsByName struct {
	RegionsDTO
}

func (r RegionsDTO) Len() int {
	return len(r)
}

func (r RegionsDTO) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RegionsByID) Less(i, j int) bool {
	return r.RegionsDTO[i].ID < r.RegionsDTO[j].ID
}

func (r RegionsByName) Less(i, j int) bool {
	return r.RegionsDTO[i].Name < r.RegionsDTO[j].Name
}

func (r RegionsDTO) SortByField(nameField string) {
	switch nameField {
	case ID:
		sort.Sort(RegionsByID{r})
	case NAME:
		sort.Sort(RegionsByName{r})
	default:
		sort.Sort(RegionsByID{r})
	}
}

func (r RegionsDTO) Search(nameField string, val any) int {
	index := r.Len()
	x := reflect.ValueOf(val)
	switch nameField {
	case ID:
		index = sort.Search(r.Len(), func(i int) bool {
			return r[i].ID >= int(x.Int())
		})
	case NAME:
		index = sort.Search(r.Len(), func(i int) bool {
			return r[i].Name >= x.String()
		})
	}
	return index
}

func (r RegionsDTO) SelectByID(regionsID ...int) RegionsDTO {
	resultRegions := RegionsDTO{}
	r.SortByField(ID)
	for _, regionID := range regionsID {
		index := r.Search(ID, regionID)
		if index < r.Len() && r[index].ID == regionID {
			resultRegions = append(resultRegions, r[index])
		}
	}
	return resultRegions
}

func (r RegionsDTO) SelectByName(regionsName ...string) RegionsDTO {
	resultRegions := RegionsDTO{}
	r.SortByField(NAME)
	for _, regionName := range regionsName {
		index := r.Search(NAME, regionName)
		if index < r.Len() && r[index].Name == regionName {
			resultRegions = append(resultRegions, r[index])
		}
	}
	return resultRegions
}
