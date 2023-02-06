package repository

import (
	"CitiesService/internal/domain/dto"
)

type Cities interface {
	Initer
	CitiesCreater
	CitiesGeter
	CitiesDeleter
	CitiesUpdater
}

type Districts interface {
	Initer
	DistrictsCreater
	DistrictsGeter
	DistrictsDeleter
	DistrictsUpdater
	Searcher
}

type Regions interface {
	Initer
	RegionsGeter
	RegionsCreater
	RegionsDeleter
	RegionsUpdater
	Searcher
}

type Initer interface {
	Init()
}

type Geter interface {
	Len() int
	GetName(ID int) (string, error)
}

type Searcher interface {
	SearchByName(name string) (int, error)
}

type Deleter interface {
	Delete(ID int) (string, error)
}

type Creater interface {
	Create(ID int, name string) (string, error)
}

type Updater interface {
	UpdateID(ID int, newID int) (int, error)
	UpdateName(ID int, newName string) (string, error)
}

type CitiesGeter interface {
	Geter
	GetCityInfo(ID int) (dto.CityDTO, error)
	GetAll() dto.CitiesDTO
	GetAllSortByField(nameField string) dto.CitiesDTO
}

type CitiesDeleter interface {
	Deleter
}

type CitiesCreater interface {
	Creater
}

type CitiesUpdater interface {
	Updater
	UpdateRegionID(ID int, newRegionID int) (int, error)
	UpdateDistrictID(ID int, newDistrictID int) (int, error)
	UpdatePopulation(ID int, newPopulation int) (int, error)
	UpdateFoundation(ID int, newFoundation int) (int, error)
}

type DistrictsCreater interface {
	Creater
}

type DistrictsGeter interface {
	Geter
	GetDistrictInfo(ID int) (dto.DistrictDTO, error)
	GetAll() dto.DistrictsDTO
}
type DistrictsDeleter interface {
	Deleter
}

type DistrictsUpdater interface {
	Updater
}

type RegionsGeter interface {
	Geter
	GetRegionInfo(ID int) (dto.RegionDTO, error)
	GetAll() dto.RegionsDTO
}

type RegionsCreater interface {
	Creater
}

type RegionsDeleter interface {
	Deleter
}

type RegionsUpdater interface {
	Updater
}
