package service

import (
	"CitiesService/internal/domain/dto"
	"CitiesService/internal/repository"
)

type DistrictsService struct {
	repositoryDistricts repository.Districts
}

func (d DistrictsService) GetInformationDistricts() dto.DistrictsDTO {
	return d.repositoryDistricts.GetAll()
}

func (d DistrictsService) CreateDistrict(district dto.DistrictDTO) error {
	_, errCreate := d.repositoryDistricts.Create(district.ID, district.Name)
	return errCreate
}

func (d DistrictsService) DeleteDistrict(ID int) (string, error) {
	name, err := d.repositoryDistricts.Delete(ID)
	return name, err
}

func (d DistrictsService) UpdateIDDistrict(ID int, newID int) error {
	_, err := d.repositoryDistricts.UpdateID(ID, newID)
	return err
}

func (d DistrictsService) UpdateNameDistrict(ID int, name string) error {
	_, err := d.repositoryDistricts.UpdateName(ID, name)
	return err
}

func (d DistrictsService) GetInformationDistrict(ID int) (dto.DistrictDTO, error) {
	cityInfo, err := d.repositoryDistricts.GetDistrictInfo(ID)
	return cityInfo, err
}

func NewDistrictsService(repositoryDistrict repository.Districts) *DistrictsService {
	return &DistrictsService{repositoryDistricts: repositoryDistrict}
}
