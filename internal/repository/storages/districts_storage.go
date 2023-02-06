package storages

import (
	"CitiesService/internal/domain"
	"CitiesService/internal/domain/dto"
	"fmt"
	"log"
	"strings"
	"sync"
)

const districtsCap = 10

type DistrictsStorage struct {
	sync.Mutex
	districtsMap map[int]*domain.District
}

func (D *DistrictsStorage) GetDistrictInfo(ID int) (dto.DistrictDTO, error) {
	err := error(nil)
	districtInfo := dto.DistrictDTO{}
	if district, ok := D.districtsMap[ID]; !ok {
		err = fmt.Errorf("getDistrictInfo: District with ID=%d is not in storages", ID)
		log.Println(err)
	} else {
		districtInfo.ID = district.ID()
		districtInfo.Name = district.Name()
	}
	return districtInfo, err
}

func (D *DistrictsStorage) Len() int {
	return len(D.districtsMap)
}

func (D *DistrictsStorage) GetAll() dto.DistrictsDTO {
	districtsDTO := make([]dto.DistrictDTO, len(D.districtsMap))
	indexDistrict := 0
	for _, district := range D.districtsMap {
		districtDTO := dto.DistrictDTO{ID: district.ID(), Name: district.Name()}
		districtsDTO[indexDistrict] = districtDTO
		indexDistrict++
	}
	return districtsDTO
}

func (D *DistrictsStorage) SearchByName(name string) (int, error) {
	err := error(nil)
	districtID := 0
	for keyID, district := range D.districtsMap {
		if strings.Contains(district.Name(), name) {
			districtID = keyID
			break
		}
	}
	if districtID == 0 {
		err = fmt.Errorf("district with name=%s was not found", name)
	}
	return districtID, err
}

func (D *DistrictsStorage) Init() {
	D.districtsMap = make(map[int]*domain.District, districtsCap)
}

func (D *DistrictsStorage) Create(ID int, name string) (string, error) {
	err := error(nil)
	districtName := ""

	if districtVal, ok := D.districtsMap[ID]; ok {
		err = fmt.Errorf("create: District with ID=%d Name=%s is in storages", districtVal.ID(), districtVal.Name())
		log.Println(err)
	} else {
		district := domain.District{}
		district.SetID(ID)
		district.SetName(name)
		D.districtsMap[district.ID()] = &district
		districtName = district.Name()
	}
	return districtName, err
}

func (D *DistrictsStorage) GetName(ID int) (string, error) {
	err := error(nil)
	districtName := ""

	if district, ok := D.districtsMap[ID]; !ok {
		err = fmt.Errorf("getName: District with ID=%d is not in storages", ID)
		log.Println(err)
	} else {
		districtName = district.Name()
	}

	return districtName, err
}

func (D *DistrictsStorage) Delete(ID int) (string, error) {
	err := error(nil)
	districtName := ""

	if district, ok := D.districtsMap[ID]; !ok {
		err = fmt.Errorf("delete: District with ID=%d is not in storages", ID)
		log.Println(err)
	} else {
		districtName = district.Name()
		delete(D.districtsMap, ID)
	}
	return districtName, err
}

func (D *DistrictsStorage) UpdateID(ID int, newID int) (int, error) {
	err := error(nil)
	districtNewID := 0
	if district, ok := D.districtsMap[ID]; !ok {
		err = fmt.Errorf("updata: District with ID=%d is not in storages", ID)
		log.Println(err)
	} else if ID == newID {
		err = fmt.Errorf("updata: District ID=%d equals newID=%d", ID, newID)
		log.Println(err)
	} else {
		district.SetID(newID)
		D.districtsMap[newID] = district
		districtNewID = newID
		delete(D.districtsMap, ID)
	}
	return districtNewID, err
}

func (D *DistrictsStorage) UpdateName(ID int, newName string) (string, error) {
	err := error(nil)
	districtName := ""

	if district, ok := D.districtsMap[ID]; !ok {
		err = fmt.Errorf("updataName: District with ID=%d is not in storages", ID)
		log.Println(err)
	} else if district.Name() == newName {
		err = fmt.Errorf("updata: District ID=%d Name=%s equals new name", ID, district.Name())
		log.Println(err)
	} else {
		districtName = district.Name()
		district.SetName(newName)
	}
	return districtName, err
}
