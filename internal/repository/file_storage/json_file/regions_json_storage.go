package json_file

import (
	"CitiesService/internal/domain/dto"
	"encoding/json"
)

type RegionsJSONStorage struct {
	regionsFile JSONFile
}

func NewRegionsJSONStorage(regionsFile JSONFile) *RegionsJSONStorage {
	return &RegionsJSONStorage{regionsFile: regionsFile}
}

func (c *RegionsJSONStorage) ReadAll() ([]dto.RegionDTO, error) {
	var (
		regions []dto.RegionDTO
		data    []byte
		err     error
	)

	data, err = c.regionsFile.ReadAll()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &regions)
	if err != nil {
		panic(err)
	}
	return regions, err
}

func (c *RegionsJSONStorage) WriteAll(regions []dto.RegionDTO) error {
	var (
		data []byte
		err  error
	)

	data, err = json.MarshalIndent(regions, "", "  ")
	if err != nil {
		panic(err)
	}

	err = c.regionsFile.WriteAll(data)
	if err != nil {
		panic(err)
	}
	return err
}
