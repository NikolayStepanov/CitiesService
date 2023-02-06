package csv_file

import (
	"CitiesService/internal/repository/file_storage/json_csv_convert"
	"encoding/json"
	"log"
	"strconv"
)

type CityStructCSV struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	RegionName   string `json:"region_name"`
	DistrictName string `json:"district_name"`
	Population   int    `json:"population"`
	Foundation   int    `json:"foundation"`
}

type CitiesCSVStorage struct {
	fileCSV CSVFile
	convert json_csv_convert.CSVJSONConvert
}

func NewCitiesCSVStorage(fileCSV CSVFile) *CitiesCSVStorage {
	return &CitiesCSVStorage{fileCSV: fileCSV}
}

func ParsingCityStringsToStructCSV(cityRecord []string) CityStructCSV {
	err := error(nil)
	city := CityStructCSV{}
	cityID := 0
	cityRegionName := ""
	cityDistrictName := ""
	cityPopulation := 0
	cityFoundation := 0
	cityName := ""

	if cityID, err = strconv.Atoi(cityRecord[0]); err != nil {
		panic(err)
	}
	if cityName = cityRecord[1]; cityName == "" {
		panic(err)
	}
	if cityRegionName = cityRecord[2]; cityRegionName == "" {
		panic(err)
	}
	if cityDistrictName = cityRecord[3]; cityDistrictName == "" {
		panic(err)
	}
	if cityPopulation, err = strconv.Atoi(cityRecord[4]); err != nil {
		panic(err)
	}
	if cityFoundation, err = strconv.Atoi(cityRecord[5]); err != nil {
		panic(err)
	}

	city.ID = cityID
	city.Name = cityName
	city.RegionName = cityRegionName
	city.DistrictName = cityDistrictName
	city.Population = cityPopulation
	city.Foundation = cityFoundation

	return city
}

func (c *CitiesCSVStorage) ParsingCityStructCSVtoStrings(city CityStructCSV, fieldSequence []string) []string {
	err := error(nil)
	cityRecord := []string{}
	cityJson, err := json.Marshal(city)

	if err != nil {
		log.Println(err)
	} else {
		cityRecord = c.convert.JSONtoCSVСonverteObject(cityJson, fieldSequence)
	}

	return cityRecord
}

func (c *CitiesCSVStorage) ReadAll() ([]CityStructCSV, error) {
	err := error(nil)
	recordsCities := [][]string{}

	recordsCities, err = c.fileCSV.ReadAll()

	cities := []CityStructCSV{}
	for _, cityRecord := range recordsCities {
		city := ParsingCityStringsToStructCSV(cityRecord)
		cities = append(cities, city)
	}
	return cities, err
}

func (c *CitiesCSVStorage) WriteAll(cities []CityStructCSV) error {
	err := error(nil)
	citiesRecords := make([][]string, len(cities))
	fieldSequence := c.convert.GetFields(CityStructCSV{})
	for index, city := range cities {
		citiesRecords[index] = c.ParsingCityStructCSVtoStrings(city, fieldSequence)
	}
	err = c.fileCSV.WriteAll(citiesRecords)
	return err
}
