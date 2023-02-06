package service

import (
	"CitiesService/internal/domain/dto"
	"CitiesService/internal/repository"
)

type RegionsService struct {
	repositoryRegions repository.Regions
}

func (r RegionsService) GetInformationRegions() dto.RegionsDTO {
	return r.repositoryRegions.GetAll()
}

func (r RegionsService) CreateRegion(region dto.RegionDTO) error {
	_, errCreate := r.repositoryRegions.Create(region.ID, region.Name)
	return errCreate
}

func (r RegionsService) DeleteRegion(ID int) (string, error) {
	name, err := r.repositoryRegions.Delete(ID)
	return name, err
}

func (r RegionsService) UpdateIDRegion(ID int, newID int) error {
	_, err := r.repositoryRegions.UpdateID(ID, newID)
	return err
}

func (r RegionsService) UpdateNameRegion(ID int, name string) error {
	_, err := r.repositoryRegions.UpdateName(ID, name)
	return err
}

func (r RegionsService) GetInformationRegion(ID int) (dto.RegionDTO, error) {
	cityInfo, err := r.repositoryRegions.GetRegionInfo(ID)
	return cityInfo, err
}

func NewRegionsService(repositoryRegions repository.Regions) *RegionsService {
	return &RegionsService{repositoryRegions: repositoryRegions}
}
