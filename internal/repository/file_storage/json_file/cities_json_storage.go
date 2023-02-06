package json_file

import (
	"CitiesService/internal/domain/dto"
	"encoding/json"
)

type CitiesJSONStorage struct {
	citiesFile JSONFile
}

func NewCitiesJSONStorage(citiesFile JSONFile) *CitiesJSONStorage {
	return &CitiesJSONStorage{citiesFile: citiesFile}
}

func (c *CitiesJSONStorage) ReadAll() ([]dto.CityDTO, error) {
	var (
		cities []dto.CityDTO
		data   []byte
		err    error
	)

	data, err = c.citiesFile.ReadAll()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &cities)
	if err != nil {
		panic(err)
	}
	return cities, err
}

func (c *CitiesJSONStorage) WriteAll(cities []dto.CityDTO) error {
	var (
		data []byte
		err  error
	)
	data, err = json.MarshalIndent(cities, "", "  ")
	if err != nil {
		panic(err)
	}
	err = c.citiesFile.WriteAll(data)
	if err != nil {
		panic(err)
	}
	return err
}
