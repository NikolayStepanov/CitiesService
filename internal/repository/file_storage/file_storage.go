package file_storage

import (
	"CitiesService/internal/domain/dto"
)

type FileStorage struct {
	CitiesFileStorage    CitiesFileStorager
	DistrictsFileStorage DistrictsFileStorager
	RegionsFileStorage   RegionsFileStorager
}

func NewFileStorage(citiesFileStorage CitiesFileStorager, districtsFileStorage DistrictsFileStorager, regionsFileStorage RegionsFileStorager) *FileStorage {
	return &FileStorage{CitiesFileStorage: citiesFileStorage, DistrictsFileStorage: districtsFileStorage, RegionsFileStorage: regionsFileStorage}
}

type CitiesFileStorager interface {
	CitiesReader
	CitiesWriter
}

type DistrictsFileStorager interface {
	DistrictsReader
	DistrictsWriter
}

type RegionsFileStorager interface {
	RegionsReader
	RegionsWriter
}

type CitiesReader interface {
	ReadAll() ([]dto.CityDTO, error)
}

type CitiesWriter interface {
	WriteAll(cities []dto.CityDTO) error
}

type DistrictsReader interface {
	ReadAll() ([]dto.DistrictDTO, error)
}

type DistrictsWriter interface {
	WriteAll(districts []dto.DistrictDTO) error
}

type RegionsReader interface {
	ReadAll() ([]dto.RegionDTO, error)
}

type RegionsWriter interface {
	WriteAll(regions []dto.RegionDTO) error
}
