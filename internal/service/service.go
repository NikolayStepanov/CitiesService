package service

import (
	"CitiesService/internal/domain/dto"
	"CitiesService/internal/repository"
)

type ServicesDeps struct {
	Repos repository.Repositories
}

func NewServicesDeps(repos repository.Repositories) *ServicesDeps {
	return &ServicesDeps{Repos: repos}
}

type Cities interface {
	GetInformationCities() dto.CitiesDTO
	GetInformationCity(ID int) (dto.CityDTO, error)
	GetCitiesByRegion(regionID int) (dto.CitiesDTO, error)
	GetCitiesByDistricts(districtID int) (dto.CitiesDTO, error)
	GetCitiesByPopulation(lowerPopulation int, upperPopulation int) (dto.CitiesDTO, error)
	GetCitiesByFoundation(lowerFoundation int, upperFoundation int) (dto.CitiesDTO, error)
	CreateCity(city dto.CityDTO) error
	DeleteCity(ID int) (string, error)
	UpdateIDCity(ID int, newID int) error
	UpdateNameCity(ID int, name string) error
	UpdatePopulationCity(ID int, population int) error
	UpdateRegionIDCity(ID int, regionID int) error
	UpdateDistrictIDCity(ID int, districtId int) error
	UpdateFoundationCity(ID int, foundation int) error
}

type Regions interface {
	GetInformationRegion(ID int) (dto.RegionDTO, error)
	GetInformationRegions() dto.RegionsDTO
	CreateRegion(region dto.RegionDTO) error
	DeleteRegion(ID int) (string, error)
	UpdateIDRegion(ID int, newID int) error
	UpdateNameRegion(ID int, name string) error
}

type Districts interface {
	GetInformationDistrict(ID int) (dto.DistrictDTO, error)
	GetInformationDistricts() dto.DistrictsDTO
	CreateDistrict(district dto.DistrictDTO) error
	DeleteDistrict(ID int) (string, error)
	UpdateIDDistrict(ID int, newID int) error
	UpdateNameDistrict(ID int, name string) error
}

type Services struct {
	Cities    Cities
	Regions   Regions
	Districts Districts
}

func NewServices(deps ServicesDeps) *Services {
	citiesService := NewCitiesService(deps.Repos.Cities)
	regionsService := NewRegionsService(deps.Repos.Regions)
	districtsService := NewDistrictsService(deps.Repos.Districts)

	return &Services{Cities: citiesService, Regions: regionsService, Districts: districtsService}
}
