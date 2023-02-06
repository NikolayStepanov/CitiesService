package json_file

import (
	"CitiesService/internal/domain/dto"
	"encoding/json"
)

type DistrictsJSONStorage struct {
	districtsFile JSONFile
}

func NewDistrictsJSONStorage(districtsFile JSONFile) *DistrictsJSONStorage {
	return &DistrictsJSONStorage{districtsFile: districtsFile}
}

func (c *DistrictsJSONStorage) ReadAll() ([]dto.DistrictDTO, error) {
	var (
		districtsDTO []dto.DistrictDTO
		data         []byte
		err          error
	)

	data, err = c.districtsFile.ReadAll()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &districtsDTO)
	if err != nil {
		panic(err)
	}

	return districtsDTO, err
}

func (c *DistrictsJSONStorage) WriteAll(districts []dto.DistrictDTO) error {
	var (
		data []byte
		err  error
	)
	data, err = json.MarshalIndent(districts, "", "  ")
	if err != nil {
		panic(err)
	}
	err = c.districtsFile.WriteAll(data)
	if err != nil {
		panic(err)
	}
	return err
}
