package storages

import (
	"CitiesService/internal/domain"
	"CitiesService/internal/domain/dto"
	"fmt"
	"log"
	"sync"
)

const citiesCap = 200

type CitiesStorage struct {
	sync.Mutex
	citiesMap map[int]*domain.City
}

func (C *CitiesStorage) GetCityInfo(ID int) (dto.CityDTO, error) {
	err := error(nil)
	cityInfo := dto.CityDTO{}
	if city, ok := C.citiesMap[ID]; !ok {
		err = fmt.Errorf("getCityInfo: City with ID=%d is not in storages", ID)
		log.Println(err)
	} else {
		cityInfo.ID = city.ID()
		cityInfo.Name = city.Name()
		cityInfo.RegionID = city.RegionID()
		cityInfo.DistrictID = city.DistrictID()
		cityInfo.Population = city.Population()
		cityInfo.Foundation = int(city.Foundation())
	}
	return cityInfo, err
}

func (C *CitiesStorage) Init() {
	C.citiesMap = make(map[int]*domain.City, citiesCap)
}

func (C *CitiesStorage) Len() int {
	return len(C.citiesMap)
}

func (C *CitiesStorage) GetAll() dto.CitiesDTO {
	citiesDTO := make([]dto.CityDTO, len(C.citiesMap))
	indexCity := 0
	for _, city := range C.citiesMap {
		cityDTO := dto.CityDTO{ID: city.ID(), RegionID: city.RegionID(), DistrictID: city.DistrictID(),
			Population: city.Population(), Foundation: int(city.Foundation()), Name: city.Name()}
		citiesDTO[indexCity] = cityDTO
		indexCity++
	}
	return citiesDTO
}

func (C *CitiesStorage) GetAllSortByField(nameField string) dto.CitiesDTO {
	cities := C.GetAll()
	cities.SortByField(nameField)
	return cities
}

func (C *CitiesStorage) Create(ID int, name string) (string, error) {
	err := error(nil)
	cityName := ""

	if cityVal, ok := C.citiesMap[ID]; ok {
		err = fmt.Errorf("create: City with ID=%d Name=%s is in storages", cityVal.ID(), cityVal.Name())
		log.Println(err)
	} else {
		city := domain.City{}
		city.SetID(ID)
		city.SetName(name)
		C.citiesMap[city.ID()] = &city
		cityName = city.Name()
	}
	return cityName, err
}

func (C *CitiesStorage) GetName(ID int) (string, error) {
	err := error(nil)
	cityName := ""

	if city, ok := C.citiesMap[ID]; !ok {
		err = fmt.Errorf("getName: City with ID=%d is not in storages", ID)
		log.Println(err)
	} else {
		cityName = city.Name()
	}

	return cityName, err
}

func (C *CitiesStorage) Delete(ID int) (string, error) {
	err := error(nil)
	cityName := ""

	if city, ok := C.citiesMap[ID]; !ok {
		err = fmt.Errorf("delete: City with ID=%d is not in storages", ID)
		log.Println(err)
	} else {
		cityName = city.Name()
		delete(C.citiesMap, ID)
	}
	return cityName, err
}

func (C *CitiesStorage) UpdateID(ID int, newID int) (int, error) {
	err := error(nil)
	cityNewID := 0
	if city, ok := C.citiesMap[ID]; !ok {
		err = fmt.Errorf("updata: City with ID=%d is not in storages", ID)
		log.Println(err)
	} else if ID == newID {
		err = fmt.Errorf("updata: City ID=%d equals newID=%d", ID, newID)
		log.Println(err)
	} else {
		city.SetID(newID)
		C.citiesMap[newID] = city
		cityNewID = newID
		delete(C.citiesMap, ID)
	}
	return cityNewID, err
}

func (C *CitiesStorage) UpdateName(ID int, newName string) (string, error) {
	err := error(nil)
	cityName := ""

	if city, ok := C.citiesMap[ID]; !ok {
		err = fmt.Errorf("updataName: City with ID=%d is not in storages", ID)
		log.Println(err)
	} else if city.Name() == newName {
		err = fmt.Errorf("updata: City ID=%d Name=%s equals new name", ID, city.Name())
		log.Println(err)
	} else {
		cityName = city.Name()
		city.SetName(newName)
	}
	return cityName, err
}

func (C *CitiesStorage) UpdateRegionID(ID int, newRegionID int) (int, error) {
	err := error(nil)
	regionOldID := 0

	if city, ok := C.citiesMap[ID]; !ok {
		err = fmt.Errorf("updataName: City with ID=%d is not in storages", ID)
		log.Println(err)
	} else if city.RegionID() == newRegionID {
		err = fmt.Errorf("updata: City ID=%d RegionID=%d equals new regionID", ID, city.RegionID())
		log.Println(err)
	} else {
		regionOldID = city.RegionID()
		city.SetRegionID(newRegionID)
	}
	return regionOldID, err
}

func (C *CitiesStorage) UpdateDistrictID(ID int, newDistrictID int) (int, error) {
	err := error(nil)
	districtOldID := 0

	if city, ok := C.citiesMap[ID]; !ok {
		err = fmt.Errorf("updataName: City with ID=%d is not in storages", ID)
		log.Println(err)
	} else if city.DistrictID() == newDistrictID {
		err = fmt.Errorf("updata: City ID=%d DistrictID=%d equals new districtID", ID, city.DistrictID())
		log.Println(err)
	} else {
		districtOldID = city.DistrictID()
		city.SetDistrictID(newDistrictID)
	}
	return districtOldID, err
}

func (C *CitiesStorage) UpdatePopulation(ID int, newPopulation int) (int, error) {
	err := error(nil)
	populationOld := 0

	if city, ok := C.citiesMap[ID]; !ok {
		err = fmt.Errorf("updataName: City with ID=%d is not in storages", ID)
		log.Println(err)
	} else if newPopulation != 0 && city.Population() == newPopulation {
		err = fmt.Errorf("updata: City ID=%d Population=%d equals new population", ID, city.Population())
		log.Println(err)
	} else {
		populationOld = city.Population()
		city.SetPopulation(newPopulation)
	}
	return populationOld, err
}

func (C *CitiesStorage) UpdateFoundation(ID int, newFoundation int) (int, error) {
	err := error(nil)
	foundationOld := 0

	if city, ok := C.citiesMap[ID]; !ok {
		err = fmt.Errorf("updataName: City with ID=%d is not in storages", ID)
		log.Println(err)
	} else if newFoundation != 0 && int(city.Foundation()) == newFoundation {
		err = fmt.Errorf("updata: City ID=%d Foundation=%d equals new foundation", ID, city.Foundation())
		log.Println(err)
	} else {
		foundationOld = int(city.Foundation())
		city.SetFoundation(uint16(newFoundation))
	}
	return foundationOld, err
}
