package service

import (
	"CitiesService/internal/domain/dto"
	"CitiesService/internal/repository"
	"errors"
	"fmt"
)

type CitiesService struct {
	repositoryCities repository.Cities
}

func (c CitiesService) GetInformationCities() dto.CitiesDTO {
	return c.repositoryCities.GetAll()
}

func (c CitiesService) UpdateIDCity(ID int, newID int) error {
	_, err := c.repositoryCities.UpdateID(ID, newID)
	return err
}

func (c CitiesService) UpdateNameCity(ID int, name string) error {
	_, err := c.repositoryCities.UpdateName(ID, name)
	return err
}

func (c CitiesService) UpdateDistrictIDCity(ID int, districtId int) error {
	_, err := c.repositoryCities.UpdateDistrictID(ID, districtId)
	return err
}

func (c CitiesService) UpdateRegionIDCity(ID int, regionID int) error {
	_, err := c.repositoryCities.UpdateRegionID(ID, regionID)
	return err
}

func (c CitiesService) UpdateFoundationCity(ID int, foundation int) error {
	_, err := c.repositoryCities.UpdateFoundation(ID, foundation)
	return err
}

func (c CitiesService) GetInformationCity(ID int) (dto.CityDTO, error) {
	cityInfo, err := c.repositoryCities.GetCityInfo(ID)
	return cityInfo, err
}

func (c CitiesService) GetCitiesByRegion(regionID int) (dto.CitiesDTO, error) {
	err := error(nil)
	citiesInfo := dto.CitiesDTO{}
	citiesInfo = citiesInfo.SelectByRegionID(regionID)
	if citiesInfo == nil {
		err = fmt.Errorf("cities with Region ID=%d not found", regionID)
	}
	return citiesInfo, err
}

func (c CitiesService) GetCitiesByDistricts(districtID int) (dto.CitiesDTO, error) {
	err := error(nil)
	citiesInfo := dto.CitiesDTO{}
	citiesInfo = citiesInfo.SelectByDistrictID(districtID)
	if citiesInfo == nil {
		err = fmt.Errorf("cities with District ID=%d not found", districtID)
	}
	return citiesInfo, err
}

func (c CitiesService) GetCitiesByPopulation(lowerPopulation int, upperPopulation int) (dto.CitiesDTO, error) {
	err := error(nil)
	citiesInfo := dto.CitiesDTO{}
	citiesInfo = citiesInfo.SelectRangePopulation(lowerPopulation, upperPopulation)
	if citiesInfo == nil {
		err = fmt.Errorf("cities included in range lower population %d and upper population %d "+
			"not found", lowerPopulation, upperPopulation)
	}
	return citiesInfo, err
}

func (c CitiesService) GetCitiesByFoundation(lowerFoundation int, upperFoundation int) (dto.CitiesDTO, error) {
	err := error(nil)
	citiesInfo := dto.CitiesDTO{}
	citiesInfo.SelectRangeFoundation(lowerFoundation, upperFoundation)
	if citiesInfo == nil {
		err = fmt.Errorf("cities included in range lower foundation %d and upper foundation %d "+
			"not found", lowerFoundation, upperFoundation)
	}
	return citiesInfo, err
}

func (c CitiesService) CreateCity(city dto.CityDTO) error {
	_, errCreate := c.repositoryCities.Create(city.ID, city.Name)
	errText := ""
	if errCreate == nil {
		if _, errUpdateRegionID := c.repositoryCities.UpdateRegionID(city.ID, city.RegionID); errUpdateRegionID != nil {
			errText = errText + errUpdateRegionID.Error()
			errCreate = errors.New(errText)
		}
		if _, errUpdateDistrictID := c.repositoryCities.UpdateDistrictID(city.ID, city.DistrictID); errUpdateDistrictID != nil {
			errText = errText + errUpdateDistrictID.Error()
			errCreate = errors.New(errText)
		}
		if _, errUpdateFoundation := c.repositoryCities.UpdateFoundation(city.ID, city.Foundation); errUpdateFoundation != nil {
			errText = errText + errUpdateFoundation.Error()
			errCreate = errors.New(errText)
		}
		if _, errUpdatePopulation := c.repositoryCities.UpdatePopulation(city.ID, city.Population); errUpdatePopulation != nil {
			errText = errText + errUpdatePopulation.Error()
			errCreate = errors.New(errText)
		}
	}
	return errCreate
}

func (c CitiesService) DeleteCity(ID int) (string, error) {
	nameCity, err := c.repositoryCities.Delete(ID)
	return nameCity, err
}

func (c CitiesService) UpdatePopulationCity(ID int, population int) error {
	_, err := c.repositoryCities.UpdatePopulation(ID, population)
	return err
}

func NewCitiesService(repository repository.Cities) *CitiesService {
	return &CitiesService{repositoryCities: repository}
}
